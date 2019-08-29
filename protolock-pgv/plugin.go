package main

import (
	"encoding/json"
	"os"

	"github.com/nilslice/protolock"
	"github.com/nilslice/protolock/extend"
)

func main() {
	plugin := extend.NewPlugin("pgv")
	plugin.Init(func(data *extend.Data) *extend.Data {
		f, _ := os.Create("debug.txt")
		defer f.Close()

		data.PluginWarnings = append(data.PluginWarnings, compareDefinitions(data.Current, data.Updated, f)...)

		f.WriteString("======================\n")

		b, _ := json.MarshalIndent(data, "", "  ")
		f.Write(b)

		return data
	})
}

// Only compare definitions that exist in both current and updated. A new definition in updated can't be wrong
// because it is new. A deleted definition in updated is already flagged by base protolock as a breaking change.

func compareDefinitions(current protolock.Protolock, updated protolock.Protolock, f *os.File) []protolock.Warning {
	warnings := []protolock.Warning{}

	// Assume definitions are sorted by Definition.Filepath
	for c, u := 0, 0; c < len(current.Definitions) && u < len(updated.Definitions); {
		cur, upd := current.Definitions[c], updated.Definitions[u]

		if cur.Filepath == upd.Filepath {
			f.WriteString("Comparing " + cur.Filepath.String() + "\n")
			warnings = append(warnings, compareMessages(cur.Def.Messages, upd.Def.Messages, upd.Filepath, f)...)
			c++
			u++
		} else if cur.Filepath < upd.Filepath {
			c++
		} else {
			u++
		}
	}

	return warnings
}

func compareMessages(current []protolock.Message, updated []protolock.Message, filepath protolock.Protopath, f *os.File) []protolock.Warning {
	warnings := []protolock.Warning{}

	// Assume messages are sorted by Message.Name
	for c, u := 0, 0; c < len(current) && u < len(updated); {
		cur, upd := current[c], updated[u]

		if cur.Name == upd.Name {
			f.WriteString("  Comparing " + cur.Name + "\n")
			warnings = append(warnings, compareFields(cur.Fields, upd.Fields, filepath, f)...)
			c++
			u++
		} else if cur.Name < upd.Name {
			c++
		} else {
			u++
		}
	}

	return warnings
}

func compareFields(current []protolock.Field, updated []protolock.Field, filepath protolock.Protopath, f *os.File) []protolock.Warning {
	warnings := []protolock.Warning{}

	// Assume messages are sorted by Message.Name
	for c, u := 0, 0; c < len(current) && u < len(updated); {
		cur, upd := current[c], updated[u]

		if cur.Name == upd.Name {
			f.WriteString("    Comparing " + cur.Name + "\n")
			warnings = append(warnings, compareOptions(cur.Options, upd.Options, filepath, f)...)
			c++
			u++
		} else if cur.Name < upd.Name {
			c++
		} else {
			u++
		}
	}

	return warnings
}

func compareOptions(current []protolock.Option, updated []protolock.Option, filepath protolock.Protopath, f *os.File) []protolock.Warning {
	warnings := []protolock.Warning{}

	// Assume messages are sorted by Message.Name
	for c, u := 0, 0; c < len(current) && u < len(updated); {
		cur, upd := current[c], updated[u]

		if cur.Name == upd.Name {
			f.WriteString("      Comparing " + cur.Name + "\n")
			warnings = append(warnings, compareOption(cur, upd, filepath, f)...)
			c++
			u++
		} else if cur.Name < upd.Name {
			c++
		} else {
			u++
		}
	}

	return warnings
}

func compareOption(current protolock.Option, updated protolock.Option, filepath protolock.Protopath, f *os.File) []protolock.Warning {
	warnings := []protolock.Warning{}

	f.WriteString("        " + current.Value + " : " + updated.Value + "\n")
	if current.Value != updated.Value {
		warnings = append(warnings, protolock.Warning{
			Filepath: filepath,
			Message:  current.Value + " != " + updated.Value,
			RuleName: "PGV",
		})
	}

	return warnings
}
