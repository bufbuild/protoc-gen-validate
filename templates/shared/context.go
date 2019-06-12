package shared

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/envoyproxy/protoc-gen-validate/gogoproto"
	"github.com/envoyproxy/protoc-gen-validate/validate"
	"github.com/golang/protobuf/proto"
	"github.com/lyft/protoc-gen-star"
)

type RuleContext struct {
	Field pgs.Field
	Rules proto.Message
	MessageRules proto.Message
	Gogo  Gogo

	Typ        string
	WrapperTyp string

	OnKey            bool
	Index            string
	AccessorOverride string
}

type Gogo struct {
	Nullable    bool
	StdDuration bool
	StdTime     bool
}

func rulesContext(f pgs.Field) (out RuleContext, err error) {
	out.Field = f

	out.Gogo.Nullable = true
	f.Extension(gogoproto.E_Nullable, &out.Gogo.Nullable)
	f.Extension(gogoproto.E_Stdduration, &out.Gogo.StdDuration)
	f.Extension(gogoproto.E_Stdtime, &out.Gogo.StdTime)

	var rules validate.FieldRules
	if _, err = f.Extension(validate.E_Rules, &rules); err != nil {
		return
	}

	var wrapped bool
	if out.Typ, out.Rules, out.MessageRules, wrapped = resolveRules(f.Type(), &rules); wrapped {
		out.WrapperTyp = out.Typ
		out.Typ = "wrapper"
	}

	if out.Typ == "error" {
		err = fmt.Errorf("unknown rule type (%T)", rules.Type)
	}

	return
}

func (ctx RuleContext) Key(name, idx string) (out RuleContext, err error) {
	rules, ok := ctx.Rules.(*validate.MapRules)
	if !ok {
		err = fmt.Errorf("cannot get Key RuleContext from %T", ctx.Field)
		return
	}

	out.Field = ctx.Field
	out.AccessorOverride = name
	out.Index = idx
	out.Gogo = ctx.Gogo

	out.Typ, out.Rules, out.MessageRules, _ = resolveRules(ctx.Field.Type().Key(), rules.GetKeys())

	if out.Typ == "error" {
		err = fmt.Errorf("unknown rule type (%T)", rules)
	}

	return
}

func (ctx RuleContext) Elem(name, idx string) (out RuleContext, err error) {
	out.Field = ctx.Field
	out.AccessorOverride = name
	out.Index = idx
	out.Gogo = ctx.Gogo

	var rules *validate.FieldRules
	switch r := ctx.Rules.(type) {
	case *validate.MapRules:
		rules = r.GetValues()
	case *validate.RepeatedRules:
		rules = r.GetItems()
	default:
		err = fmt.Errorf("cannot get Elem RuleContext from %T", ctx.Field)
		return
	}

	var wrapped bool
	if out.Typ, out.Rules, out.MessageRules, wrapped = resolveRules(ctx.Field.Type().Element(), rules); wrapped {
		out.WrapperTyp = out.Typ
		out.Typ = "wrapper"
	}

	if out.Typ == "error" {
		err = fmt.Errorf("unknown rule type (%T)", rules)
	}

	return
}

func (ctx RuleContext) Unwrap(name string) (out RuleContext, err error) {
	if ctx.Typ != "wrapper" {
		err = fmt.Errorf("cannot unwrap non-wrapper type %q", ctx.Typ)
		return
	}

	return RuleContext{
		Field:            ctx.Field,
		Rules:            ctx.Rules,
		MessageRules:	  ctx.MessageRules,
		Typ:              ctx.WrapperTyp,
		AccessorOverride: name,
		Gogo:             ctx.Gogo,
	}, nil
}

func Render(tpl *template.Template) func(ctx RuleContext) (string, error) {
	return func(ctx RuleContext) (string, error) {
		var b bytes.Buffer
		err := tpl.ExecuteTemplate(&b, ctx.Typ, ctx)
		return b.String(), err
	}
}

func resolveRules(typ interface{ IsEmbed() bool }, rules *validate.FieldRules) (ruleType string, rule proto.Message, messageRule proto.Message, wrapped bool) {
	switch r := rules.GetType().(type) {
	case *validate.FieldRules_Float:
		return "float", r.Float, rules.Message, typ.IsEmbed()
	case *validate.FieldRules_Double:
		return "double", r.Double, rules.Message, typ.IsEmbed()
	case *validate.FieldRules_Int32:
		return "int32", r.Int32, rules.Message, typ.IsEmbed()
	case *validate.FieldRules_Int64:
		return "int64", r.Int64, rules.Message, typ.IsEmbed()
	case *validate.FieldRules_Uint32:
		return "uint32", r.Uint32, rules.Message, typ.IsEmbed()
	case *validate.FieldRules_Uint64:
		return "uint64", r.Uint64, rules.Message, typ.IsEmbed()
	case *validate.FieldRules_Sint32:
		return "sint32", r.Sint32, rules.Message, false
	case *validate.FieldRules_Sint64:
		return "sint64", r.Sint64, rules.Message, false
	case *validate.FieldRules_Fixed32:
		return "fixed32", r.Fixed32, rules.Message, false
	case *validate.FieldRules_Fixed64:
		return "fixed64", r.Fixed64, rules.Message, false
	case *validate.FieldRules_Sfixed32:
		return "sfixed32", r.Sfixed32, rules.Message, false
	case *validate.FieldRules_Sfixed64:
		return "sfixed64", r.Sfixed64, rules.Message, false
	case *validate.FieldRules_Bool:
		return "bool", r.Bool, rules.Message, typ.IsEmbed()
	case *validate.FieldRules_String_:
		return "string", r.String_, rules.Message, typ.IsEmbed()
	case *validate.FieldRules_Bytes:
		return "bytes", r.Bytes, rules.Message, typ.IsEmbed()
	case *validate.FieldRules_Enum:
		return "enum", r.Enum, rules.Message, false
	case *validate.FieldRules_Repeated:
		return "repeated", r.Repeated, rules.Message, false
	case *validate.FieldRules_Map:

		return "map", r.Map, rules.Message, false
	case *validate.FieldRules_Any:
		return "any", r.Any, rules.Message, false
	case *validate.FieldRules_Duration:
		return "duration", r.Duration, rules.Message, false
	case *validate.FieldRules_Timestamp:
		return "timestamp", r.Timestamp, rules.Message, false
	case nil:
		if ft, ok := typ.(pgs.FieldType); ok && ft.IsRepeated() {
			return "repeated", &validate.RepeatedRules{}, rules.Message, false
		} else if typ.IsEmbed() {
			return "message", rules.GetMessage(), rules.GetMessage(), false
		}
		return "none", nil, nil, false
	default:
		return "error", nil, rules.Message, false
	}
}
