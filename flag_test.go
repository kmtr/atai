package atai

import (
	"flag"
	"testing"
)

func TestValueFromFlag(t *testing.T) {
	type testarg struct {
		flagName     string
		defaultValue string
		argValue     string
		want         string
	}

	tests := []testarg{
		{"-a", "default", "arg", "arg"},
		{"-b", "default", "", "default"},
		{"not registered", "-", "-", ""},
	}

	for _, test := range tests {
		if test.flagName != "not registered" {
			flag.String(test.flagName, test.defaultValue, "usage")
			if test.argValue != "" {
				flag.Set(test.flagName, test.argValue)
			}
		}
	}

	flag.Parse()

	for _, test := range tests {
		t.Run(test.flagName, func(t *testing.T) {
			got := ValueFromFlag(test.flagName)()
			if test.want != got {
				t.Errorf("want: %s, got: %s", test.want, got)
			}
		})
	}
}
