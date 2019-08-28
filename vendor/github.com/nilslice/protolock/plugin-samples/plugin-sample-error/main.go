package main

import (
	"github.com/nilslice/protolock"
	"github.com/nilslice/protolock/extend"
)

func main() {
	plugin := extend.NewPlugin("sample-error") // "sample-error" is arbitrary name used to correlate error messages
	plugin.Init(func(data *extend.Data) *extend.Data {
		warnings := AddWarningsForExample(data.Current, data.Updated)
		data.PluginWarnings = append(data.PluginWarnings, warnings...)
		data.PluginErrorMessage = "some error"
		return data
	})
}

func AddWarningsForExample(cur, upd protolock.Protolock) []protolock.Warning {
	return []protolock.Warning{
		{Filepath: protolock.OSPath(upd.Definitions[0].Filepath), Message: "A sample warning!"},
		{Filepath: protolock.OSPath(upd.Definitions[0].Filepath), Message: "Another sample warning.. ah!"},
	}
}
