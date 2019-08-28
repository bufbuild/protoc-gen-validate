package extend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/nilslice/protolock"
)

// Plugin is an interface that defines the protolock plugin specification.
type Plugin interface {
	Init(PluginFunc)
}

// Data contains the current and updated Protolock structs created by the
// `protolock` internal parser and deserializer, and a slice of Warning structs
// for the plugin to append its own custom warnings.
type Data struct {
	Current            protolock.Protolock `json:"current,omitempty"`
	Updated            protolock.Protolock `json:"updated,omitempty"`
	ProtolockWarnings  []protolock.Warning `json:"protolock_warnings,omitempty"`
	PluginWarnings     []protolock.Warning `json:"plugin_warnings,omitempty"`
	PluginErrorMessage string              `json:"plugin_error_message,omitempty"`
}

// PluginFunc is a function which defines plugin behavior, and is provided a
// pointer to Data.
type PluginFunc func(d *Data) *Data

type plugin struct {
	name string
}

// NewPlugin returns a plugin instance for a plugin to be initialized.
func NewPlugin(name string) *plugin {
	return &plugin{
		name: name,
	}
}

// Init is called by plugin code and is provided a PluginFunc from the caller
// to handle the input Data (read from stdin).
func (p *plugin) Init(fn PluginFunc) {
	// read from stdin to get serialized bytes
	input := &bytes.Buffer{}
	_, err := io.Copy(input, os.Stdin)
	if err != nil {
		p.wrapErrAndLog(err)
		return
	}

	// deserialize bytes into *Data
	inputData := &Data{}
	err = json.Unmarshal(input.Bytes(), inputData)
	if err != nil {
		p.wrapErrAndLog(err)
		return
	}

	// execute "fn" and pass it the *Data, where the plugin would read and
	// compare the current and updated Protolock values and append custom
	// Warnings for their own defined rules
	outputData := fn(inputData)
	outputData.Current = inputData.Current
	outputData.Updated = inputData.Updated

	// serialize *Data back and write to stdout
	p.wrapErrAndLog(json.NewEncoder(os.Stdout).Encode(outputData))
}

func (p *plugin) wrapErrAndLog(err error) {
	if err != nil {
		fmt.Fprintf(os.Stdout, "[protolock:plugin] %s: %v", p.name, err)
	}
}

var _ Plugin = &plugin{}
