package main

import (
	"fmt"

	"github.com/nilslice/protolock/extend"
)

func main() {
	fmt.Println("Hello protolock!")
	plugin := extend.NewPlugin("pgv") // "pgv" is arbitrary name used to correlate error messages
	plugin.Init(func(data *extend.Data) *extend.Data {
		return data
	})
}
