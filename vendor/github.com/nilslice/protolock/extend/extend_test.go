package extend

import (
	"testing"

	"github.com/nilslice/protolock"
)

const fakeData = "test error message"

type fakePlugin struct{}

func (p *fakePlugin) Init(fn PluginFunc) {
	data := &Data{
		Current:            protolock.Protolock{},
		Updated:            protolock.Protolock{},
		PluginWarnings:     nil,
		PluginErrorMessage: fakeData,
	}

	_ = fn(data)
}

func TestPluginInit(t *testing.T) {
	p := &fakePlugin{}
	p.Init(func(data *Data) *Data {
		if data.PluginErrorMessage != "test error message" {
			t.Logf("incorrect error message: %s", data.PluginErrorMessage)
			t.Fail()
		}

		if len(data.Current.Definitions) != 0 {
			t.Fail()
		}

		if len(data.Updated.Definitions) != 0 {
			t.Fail()
		}

		if data.PluginWarnings != nil {
			t.Fail()
		}

		return data
	})
}
