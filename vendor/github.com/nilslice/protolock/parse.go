package protolock

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/emicklei/proto"
)

const LockFileName = "proto.lock"

type Protolock struct {
	Definitions []Definition `json:"definitions,omitempty"`
}

type Definition struct {
	Filepath Protopath `json:"protopath,omitempty"`
	Def      Entry     `json:"def,omitempty"`
}

type Entry struct {
	Enums    []Enum    `json:"enums,omitempty"`
	Messages []Message `json:"messages,omitempty"`
	Services []Service `json:"services,omitempty"`
	Imports  []Import  `json:"imports,omitempty"`
	Package  Package   `json:"package,omitempty"`
	Options  []Option  `json:"options,omitempty"`
}

type Import struct {
	Path string `json:"path,omitempty"`
}

type Package struct {
	Name string `json:"name,omitempty"`
}

type Option struct {
	Name       string   `json:"name,omitempty"`
	Value      string   `json:"value,omitempty"`
	Aggregated []Option `json:"aggregated,omitempty"`
}

type Message struct {
	Name          string    `json:"name,omitempty"`
	Fields        []Field   `json:"fields,omitempty"`
	Maps          []Map     `json:"maps,omitempty"`
	ReservedIDs   []int     `json:"reserved_ids,omitempty"`
	ReservedNames []string  `json:"reserved_names,omitempty"`
	Filepath      Protopath `json:"filepath,omitempty"`
	Messages      []Message `json:"messages,omitempty"`
	Options       []Option  `json:"options,omitempty"`
}

type EnumField struct {
	Name    string   `json:"name,omitempty"`
	Integer int      `json:"integer,omitempty"`
	Options []Option `json:"options,omitempty"`
}

type Enum struct {
	Name          string      `json:"name,omitempty"`
	EnumFields    []EnumField `json:"enum_fields,omitempty"`
	ReservedIDs   []int       `json:"reserved_ids,omitempty"`
	ReservedNames []string    `json:"reserved_names,omitempty"`
	AllowAlias    bool        `json:"allow_alias,omitempty"`
}

type Map struct {
	KeyType string `json:"key_type,omitempty"`
	Field   Field  `json:"field,omitempty"`
}

type Field struct {
	ID         int      `json:"id,omitempty"`
	Name       string   `json:"name,omitempty"`
	Type       string   `json:"type,omitempty"`
	IsRepeated bool     `json:"is_repeated,omitempty"`
	Options    []Option `json:"options,omitempty"`
}

type Service struct {
	Name     string    `json:"name,omitempty"`
	RPCs     []RPC     `json:"rpcs,omitempty"`
	Filepath Protopath `json:"filepath,omitempty"`
}

type RPC struct {
	Name        string   `json:"name,omitempty"`
	InType      string   `json:"in_type,omitempty"`
	OutType     string   `json:"out_type,omitempty"`
	InStreamed  bool     `json:"in_streamed,omitempty"`
	OutStreamed bool     `json:"out_streamed,omitempty"`
	Options     []Option `json:"options,omitempty"`
}

type Report struct {
	Current  Protolock `json:"current,omitempty"`
	Updated  Protolock `json:"updated,omitempty"`
	Warnings []Warning `json:"warnings,omitempty"`
}

type Warning struct {
	Filepath Protopath `json:"filepath,omitempty"`
	Message  string    `json:"message,omitempty"`
	RuleName string    `json:"rulename,omitempty"`
}

type ProtoFile struct {
	ProtoPath Protopath
	Entry     Entry
}

var (
	enums []Enum
	msgs  []Message
	svcs  []Service
	imps  []Import
	pkg   Package
	opts  []Option

	ErrWarningsFound = errors.New("comparison found one or more warnings")
)

func Parse(filename string, r io.Reader) (Entry, error) {
	parser := proto.NewParser(r)
	parser.Filename(filename)
	def, err := parser.Parse()
	if err != nil {
		return Entry{}, err
	}

	enums = []Enum{}
	msgs = []Message{}
	svcs = []Service{}
	imps = []Import{}
	opts = []Option{}

	proto.Walk(
		def,
		proto.WithEnum(withEnum),
		proto.WithService(withService),
		proto.WithMessage(withMessage),
		protoWithImport(withImport),
		protoWithPackage(withPackage),
		proto.WithOption(withOption),
	)

	return Entry{
		Enums:    enums,
		Messages: msgs,
		Services: svcs,
		Imports:  imps,
		Package:  pkg,
		Options:  opts,
	}, nil
}

func withEnum(e *proto.Enum) {
	errs := checkComments(e)
	if errs != nil {
		for _, err := range errs {
			switch err {
			case ErrSkipEntry:
				return
			}
		}
	}

	// handle nested enum within message, prepend message name to enum name
	if p, ok := e.Parent.(*proto.Message); ok {
		if p != nil {
			e.Name = fmt.Sprintf("%s.%s", p.Name, e.Name)
		}
	}

	enums = append(enums, parseEnum(e))
}

func parseEnum(e *proto.Enum) Enum {
	enum := Enum{
		Name: e.Name,
	}

	for _, v := range e.Elements {
		if ef, ok := v.(*proto.EnumField); ok {
			field := EnumField{
				Name:    ef.Name,
				Integer: ef.Integer,
			}
			for _, ee := range ef.Elements {
				if o, ok := ee.(*proto.Option); ok {
					field.Options = append(field.Options, Option{
						Name:  o.Name,
						Value: o.Constant.Source,
					})
				}
			}
			enum.EnumFields = append(enum.EnumFields, field)
		}

		if r, ok := v.(*proto.Reserved); ok {
			// collect all reserved field IDs from the ranges
			for _, rng := range r.Ranges {
				// if range is only a single value, skip loop and
				// append single value to message's reserved slice
				if rng.From == rng.To {
					enum.ReservedIDs = append(enum.ReservedIDs, rng.From)
					continue
				}
				// add each item from the range inclusively
				for id := rng.From; id <= rng.To; id++ {
					enum.ReservedIDs = append(enum.ReservedIDs, id)
				}
			}

			// add all reserved field names
			enum.ReservedNames = append(enum.ReservedNames, r.FieldNames...)
		}
	}

	return enum
}

func withService(s *proto.Service) {
	errs := checkComments(s)
	if errs != nil {
		for _, err := range errs {
			switch err {
			case ErrSkipEntry:
				return
			}
		}
	}

	svc := Service{
		Name: s.Name,
	}

	for _, v := range s.Elements {
		if r, ok := v.(*proto.RPC); ok {
			svc.RPCs = append(svc.RPCs, RPC{
				Name:        r.Name,
				InType:      r.RequestType,
				OutType:     r.ReturnsType,
				InStreamed:  r.StreamsRequest,
				OutStreamed: r.StreamsReturns,
				Options:     parseOptions(r.Options),
			})
		}
	}

	svcs = append(svcs, svc)
}

func withMessage(m *proto.Message) {
	errs := checkComments(m)
	if errs != nil {
		for _, err := range errs {
			switch err {
			case ErrSkipEntry:
				return
			}
		}
	}

	if _, ok := m.Parent.(*proto.Proto); !ok {
		return
	}

	msgs = append(msgs, parseMessage(m))
}

func parseMessage(m *proto.Message) Message {
	msg := Message{
		Name: m.Name,
	}

	for _, v := range m.Elements {

		if f, ok := v.(*proto.NormalField); ok {
			msg.Fields = append(msg.Fields, Field{
				ID:         f.Sequence,
				Name:       f.Name,
				Type:       f.Type,
				IsRepeated: f.Repeated,
				Options:    parseOptions(f.Options),
			})
		}

		if mp, ok := v.(*proto.MapField); ok {
			f := mp.Field
			msg.Maps = append(msg.Maps, Map{
				KeyType: mp.KeyType,
				Field: Field{
					ID:         f.Sequence,
					Name:       f.Name,
					Type:       f.Type,
					IsRepeated: false,
					Options:    parseOptions(f.Options),
				},
			})
		}

		if oo, ok := v.(*proto.Oneof); ok {
			var fields []Field
			for _, el := range oo.Elements {
				if f, ok := el.(*proto.OneOfField); ok {
					fields = append(fields, Field{
						ID:         f.Sequence,
						Name:       f.Name,
						Type:       f.Type,
						IsRepeated: false,
						Options:    parseOptions(f.Options),
					})
				}
			}
			msg.Fields = append(msg.Fields, fields...)
		}

		if r, ok := v.(*proto.Reserved); ok {
			// collect all reserved field IDs from the ranges
			for _, rng := range r.Ranges {
				// if range is only a single value, skip loop and
				// append single value to message's reserved slice
				if rng.From == rng.To {
					msg.ReservedIDs = append(msg.ReservedIDs, rng.From)
					continue
				}
				// add each item from the range inclusively
				for id := rng.From; id <= rng.To; id++ {
					msg.ReservedIDs = append(msg.ReservedIDs, id)
				}
			}

			// add all reserved field names
			msg.ReservedNames = append(msg.ReservedNames, r.FieldNames...)
		}

		if o, ok := v.(*proto.Option); ok {
			msg.Options = append(msg.Options, parseOption(o))
		}

		if m, ok := v.(*proto.Message); ok {
			msg.Messages = append(msg.Messages, parseMessage(m))
		}
	}

	return msg
}

func withOption(o *proto.Option) {
	if _, ok := o.Parent.(*proto.Proto); !ok {
		return
	}
	opts = append(opts, parseOption(o))
}

func parseOptions(opts []*proto.Option) []Option {
	var msgOpts []Option
	for _, o := range opts {
		msgOpts = append(msgOpts, parseOption(o))
	}
	return msgOpts
}

func parseOption(o *proto.Option) Option {
	option := Option{
		Name: o.Name,
	}
	if isAggregatedOption(o) {
		option.Aggregated = parseAggregatedValues(o)
	} else {
		option.Value = o.Constant.Source
	}
	return option
}

func parseAggregatedValues(o *proto.Option) []Option {
	var aggOpts []Option
	for _, nl := range o.Constant.OrderedMap {
		aggOpts = append(aggOpts, Option{
			Name:  nl.Name,
			Value: nl.Source,
		})
	}
	return aggOpts
}

func isAggregatedOption(o *proto.Option) bool {
	return o.Constant.Source == "" && o.Constant.OrderedMap != nil
}

func protoWithImport(apply func(p *proto.Import)) proto.Handler {
	return func(v proto.Visitee) {
		if s, ok := v.(*proto.Import); ok {
			apply(s)
		}
	}
}

func withImport(im *proto.Import) {
	imp := Import{
		Path: im.Filename,
	}
	imps = append(imps, imp)
}

func protoWithPackage(apply func(p *proto.Package)) proto.Handler {
	return func(v proto.Visitee) {
		if s, ok := v.(*proto.Package); ok {
			apply(s)
		}
	}
}

func withPackage(im *proto.Package) {
	pkg = Package{
		Name: im.Name,
	}
}

// openLockFile opens and returns the lock file on disk for reading.
func openLockFile(cfg Config) (io.ReadCloser, error) {
	f, err := os.Open(cfg.LockFilePath())
	if err != nil {
		return nil, err
	}

	return f, nil
}

// FromReader unmarshals a proto.lock file into a Protolock struct.
func FromReader(r io.Reader) (Protolock, error) {
	buf := bytes.Buffer{}
	_, err := io.Copy(&buf, r)
	if err != nil {
		return Protolock{}, err
	}

	var lock Protolock
	err = json.Unmarshal(buf.Bytes(), &lock)
	if err != nil {
		return Protolock{}, err
	}

	return lock, nil
}

// Compare returns a Report struct and an error which indicates that there is
// one or more warnings to report to the caller. If no error is returned, the
// Report can be ignored.
func Compare(current, update Protolock) (*Report, error) {
	var warnings []Warning
	var wg sync.WaitGroup
	report := &Report{
		Current: current,
		Updated: update,
	}
	for _, rule := range Rules {
		wg.Add(1)
		go func() {
			if debug {
				beginRuleDebug(rule.Name)
			}
			_warnings, _ := rule.Func(current, update)
			for i := range _warnings {
				_warnings[i].RuleName = rule.Name
			}
			if debug {
				concludeRuleDebug(rule.Name, _warnings)
			}

			warnings = append(warnings, _warnings...)
			wg.Done()
		}()
		wg.Wait()
	}
	report.Warnings = warnings

	if len(report.Warnings) != 0 {
		return report, ErrWarningsFound
	}

	return report, nil
}

// getProtoFiles finds recursively all .proto files to be processed.
func getProtoFiles(root string, ignores string) ([]string, error) {
	protoFiles := []string{}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// if not a .proto file, do not attempt to parse.
		if !strings.HasSuffix(info.Name(), protoSuffix) {
			return nil
		}

		// skip to next if is a directory
		if info.IsDir() {
			return nil
		}

		// skip if path is within an ignored path
		if ignores != "" {
			for _, ignore := range strings.Split(ignores, ",") {
				rel, err := filepath.Rel(filepath.Join(root, ignore), path)
				if err != nil {
					return nil
				}

				if !strings.HasPrefix(rel, ".."+string(os.PathSeparator)) {
					return nil
				}
			}
		}

		protoFiles = append(protoFiles, path)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return protoFiles, nil
}

// getUpdatedLock finds all .proto files recursively in tree, parse each file
// and accumulate all definitions into an updated Protolock.
func getUpdatedLock(cfg Config) (*Protolock, error) {
	// files is a slice of struct `ProtoFile` to be joined into the proto.lock file.
	var files []ProtoFile

	root, err := filepath.Abs(cfg.ProtoRoot)
	if err != nil {
		return nil, err
	}

	protoFiles, err := getProtoFiles(root, cfg.Ignore)
	if err != nil {
		return nil, err
	}

	for _, path := range protoFiles {
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}

		// Have the parser report the file path
		friendlyPath := path
		cwd, err := os.Getwd()
		if err == nil {
			relpath, err := filepath.Rel(cwd, path)
			if err == nil {
				friendlyPath = relpath
			}
		}
		entry, err := Parse(friendlyPath, f)
		if err != nil {
			printIfErr(f.Close())
			return nil, err
		}

		localPath := strings.TrimPrefix(path, root)
		localPath = strings.TrimPrefix(localPath, string(filepath.Separator))
		protoFile := ProtoFile{
			ProtoPath: ProtoPath(Protopath(localPath)),
			Entry:     entry,
		}
		files = append(files, protoFile)

		// manually close the file to prevent `too many open files` error
		printIfErr(f.Close())
	}

	// add all the definitions from the updated set of protos to a Protolock
	// used for analysis and comparison against the current Protolock, saved
	// as the proto.lock file in the current directory
	var updated Protolock
	for _, file := range files {
		updated.Definitions = append(updated.Definitions, Definition{
			Filepath: file.ProtoPath,
			Def:      file.Entry,
		})
	}

	return &updated, nil
}

func printIfErr(err error) {
	if err != nil {
		fmt.Printf("protolock: %v\n", err)
	}
}
