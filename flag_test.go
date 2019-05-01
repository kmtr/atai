package atai

import (
	"flag"
	"strconv"
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

func TestFlagValue(t *testing.T) {
	type testarg struct {
		key      string
		value    string
		usage    string
		defValue string
	}

	tests := []testarg{
		testarg{
			key:      "key1",
			value:    "val1",
			usage:    "flag usage",
			defValue: "def1",
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			flag.String(test.key, test.defValue, test.usage)
			flag.Set(test.key, test.value)
			fv := NewFlagValue(test.key)

			var _ KeyHolder = fv
			var _ ValueHolder = fv
			var _ Explainer = fv
			var _ DefaultValueHolder = fv

			var want, got string
			want, got = test.key, fv.Key()
			if got != want {
				t.Errorf("want: %s, got: %s", want, got)
			}

			want, got = test.value, fv.Value()
			if got != want {
				t.Errorf("want: %s, got: %s", want, got)
			}

			want, got = test.defValue, fv.DefValue()
			if got != want {
				t.Errorf("want: %s, got: %s", want, got)
			}

			want, got = test.usage, fv.Explain()
			if got != want {
				t.Errorf("want: %s, got: %s", want, got)
			}
		})
	}
}
