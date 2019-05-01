package atai

import (
	"os"
	"strconv"
	"testing"
)

func TestValueFromEnv(t *testing.T) {
	type testarg struct {
		envName  string
		defValue string
		envValue string
		want     string
	}

	tests := []testarg{
		{"E", "default", "arg", "arg"},
		{"F", "default", "", ""},
		{"not registered", "default", "-", ""},
	}

	for _, test := range tests {
		if test.envName != "not registered" {
			os.Setenv(test.envName, test.envValue)
		}
	}

	for _, test := range tests {
		t.Run(test.envName, func(t *testing.T) {
			got := ValueFromEnv(test.envName)()
			if test.want != got {
				t.Errorf("want: %s, got: %s", test.want, got)
			}

			got = ValueFromEnvWithDefault(test.envName, test.defValue)()
			if test.envName == "not registered" {
				if test.defValue != got {
					t.Errorf("want: %s, got: %s", test.defValue, got)
				}
			} else {
				if test.want != got {
					t.Errorf("want: %s, got: %s", test.want, got)
				}
			}
		})
	}
}

func TestEnvValue(t *testing.T) {
	type testarg struct {
		key      string
		value    string
		defValue string
	}

	tests := []testarg{
		testarg{
			key:      "key1",
			value:    "val1",
			defValue: "def1",
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			os.Setenv(test.key, test.value)
			ev := NewEnvValue(test.key)

			var _ KeyHolder = ev
			var _ ValueHolder = ev
			var _ DefaultValueHolder = ev

			if ev.Key() != test.key {
				t.Errorf("want: %s, got: %s", test.key, ev.Key())
			}
			if ev.Value() != test.value {
				t.Errorf("want: %s, got: %s", test.value, ev.Value())
			}
			if ev.DefValue() != "" {
				t.Errorf("want: %s, got: %s", "", ev.DefValue())
			}

			ev = NewEnvValueWithDefault(test.key, test.defValue)
			if ev.Key() != test.key {
				t.Errorf("want: %s, got: %s", test.key, ev.Key())
			}
			if ev.Value() != test.value {
				t.Errorf("want: %s, got: %s", test.value, ev.Value())
			}

			os.Unsetenv(test.key)
			if ev.DefValue() != test.defValue {
				t.Errorf("want: %s, got: %s", test.defValue, ev.DefValue())
			}
		})
	}
}
