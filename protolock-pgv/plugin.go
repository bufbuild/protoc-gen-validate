package main

import (
	"fmt"

	"github.com/nilslice/protolock"
	"github.com/nilslice/protolock/extend"
)

func main() {
	plugin := extend.NewPlugin("pgv")
	plugin.Init(func(data *extend.Data) *extend.Data {
		data.PluginWarnings = append(data.PluginWarnings, compareDefinitions(data.Current, data.Updated)...)
		return data
	})
}

// Only compare definitions that exist in both current and updated. A new definition in updated can't be wrong
// because it is new. A deleted definition in updated is already flagged by base protolock as a breaking change.

func compareDefinitions(current protolock.Protolock, updated protolock.Protolock) []protolock.Warning {
	warnings := []protolock.Warning{}

	// Assume definitions are sorted by Definition.Filepath
	for c, u := 0, 0; c < len(current.Definitions) && u < len(updated.Definitions); {
		cur, upd := current.Definitions[c], updated.Definitions[u]

		if cur.Filepath == upd.Filepath {
			warnings = append(warnings, compareMessages(cur.Def.Messages, upd.Def.Messages, upd.Filepath)...)
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

func compareMessages(current []protolock.Message, updated []protolock.Message, filepath protolock.Protopath) []protolock.Warning {
	warnings := []protolock.Warning{}

	// Assume messages are sorted by Message.Name
	for c, u := 0, 0; c < len(current) && u < len(updated); {
		cur, upd := current[c], updated[u]

		if cur.Name == upd.Name {
			warnings = append(warnings, compareFields(fmt.Sprintf("%q", cur.Name), cur.Fields, upd.Fields, filepath)...)
			// warnings = append(warnings, compareFields("\""+cur.Name+"\"", cur.Fields, upd.Fields, filepath)...)
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

func compareFields(where string, current []protolock.Field, updated []protolock.Field, filepath protolock.Protopath) []protolock.Warning {
	warnings := []protolock.Warning{}

	// Assume messages are sorted by Message.Name
	for c, u := 0, 0; c < len(current) && u < len(updated); {
		cur, upd := current[c], updated[u]

		if cur.Name == upd.Name {
			warnings = append(warnings, compareOptions(fmt.Sprintf("%s field: %q", where, cur.Name), flatOptions(cur.Options), flatOptions(upd.Options), filepath)...)
			// warnings = append(warnings, compareOptions(where+" field: \""+cur.Name+"\"", flatOptions(cur.Options), flatOptions(upd.Options), filepath)...)
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

// Flatten a []protolock.Options, decomposing aggregated options into multiple separate Options
func flatOptions(optionsWithAggregates []protolock.Option) []protolock.Option {
	options := []protolock.Option{}

	for _, opt := range optionsWithAggregates {
		if opt.Aggregated == nil {
			options = append(options, opt)
		} else {
			for _, agg := range opt.Aggregated {
				options = append(options, protolock.Option{
					Name:  opt.Name + "." + agg.Name,
					Value: agg.Value,
				})
			}
		}
	}

	return options
}

func compareOptions(where string, current []protolock.Option, updated []protolock.Option, filepath protolock.Protopath) []protolock.Warning {
	warnings := []protolock.Warning{}

	nullTerm := protolock.Option{
		Name: "NULL-TERMINATOR-NULL-TERMINATOR",
	}

	// Add "null terminators" to ensure empty options lists don't prematurely crash out of the loop
	current = append(current, nullTerm)
	updated = append(updated, nullTerm)

	// Assume messages are sorted by Message.Name
	for c, u := 0, 0; c < len(current) && u < len(updated); {
		cur, upd := current[c], updated[u]

		if cur.Name == upd.Name {
			warnings = append(warnings, compareOption(where, cur, upd, filepath)...)
			c++
			u++
		} else if cur.Name < upd.Name {
			warnings = append(warnings, protolock.Warning{
				Filepath: filepath,
				Message:  fmt.Sprintf("%s constraint: %q has been removed", where, cur.Name),
				// Message:  where + " constraint: \"" + cur.Name + "\" has been removed",
				RuleName: "PGV",
			})
			c++
		} else {
			u++
		}
	}

	return warnings
}

func compareOption(where string, current protolock.Option, updated protolock.Option, filepath protolock.Protopath) []protolock.Warning {
	warnings := []protolock.Warning{}

	if current.Value != updated.Value {
		warnings = append(warnings, protolock.Warning{
			Filepath: filepath,
			Message:  fmt.Sprintf("%s constraint: %q has changed from %s to %s", where, current.Name, current.Value, updated.Value),
			// Message:  where + " constraint: \"" + current.Name + "\" has changed from " + current.Value + " to " + updated.Value,
			RuleName: "PGV",
		})
	}

	return warnings
}
