package licences

import (
	"reflect"
	"testing"
)

func TestLicenses(t *testing.T) {
	testCases := []struct {
		desc string
		l    func() string
	}{
		{
			desc: "MIT",
			l:    MITTemplate,
		},
		{
			desc: "Apache",
			l:    ApacheTemplate,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if reflect.TypeOf(tC.l()) != reflect.TypeOf("string") {
				t.Errorf("TestLicenses: %v | Return type is not string", tC.desc)
			}
		})
	}
}
