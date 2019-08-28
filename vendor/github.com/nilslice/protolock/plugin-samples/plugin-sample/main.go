package main

import (
	"fmt"
	"os"

	"github.com/nilslice/protolock"
	"github.com/nilslice/protolock/extend"
)

func main() {
	plugin := extend.NewPlugin("sample") // "sample" is arbitrary name used to correlate error messages
	plugin.Init(func(data *extend.Data) *extend.Data {
		// list all existing rules violated from warnings passed into plugin
		// from protolock & write to output file
		out, err := os.Create("violations.txt")
		if err != nil {
			return data
		}
		for _, w := range data.ProtolockWarnings {
			fmt.Fprintln(
				out, "Encountered changes in violation of:", w.RuleName,
			)
		}

		warnings := AddWarningsForExample(data.Current, data.Updated)
		data.PluginWarnings = append(data.PluginWarnings, warnings...)
		return data
	})
}

func AddWarningsForExample(cur, upd protolock.Protolock) []protolock.Warning {
	return []protolock.Warning{
		{
			Filepath: protolock.OSPath(upd.Definitions[0].Filepath),
			Message:  "A sample warning!",
		},
		{
			Filepath: protolock.OSPath(upd.Definitions[0].Filepath),
			Message:  "Another sample warning.. ah!",
		},
	}
}
