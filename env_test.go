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
		explain  string
		defValue string
	}

	tests := []testarg{
		testarg{
			key:      "key1",
			value:    "val1",
			explain:  "test value",
			defValue: "def1",
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			os.Setenv(test.key, test.value)
			ev := NewEnvValue(test.key, test.explain)

			var _ KeyHolder = ev
			var _ ValueProviderHolder = ev
			var _ Explainer = ev
			var _ DefaultValueHolder = ev

			var want, got string
			want, got = test.key, ev.Key()
			if got != want {
				t.Errorf("want: %s, got: %s", want, got)
			}

			want, got = test.value, ev.ValueProvider()()
			if got != want {
				t.Errorf("want: %s, got: %s", want, got)
			}

			want, got = "", ev.DefValue()
			if got != want {
				t.Errorf("want: %s, got: %s", want, got)
			}

			want, got = test.explain, ev.Explain()
			if got != want {
				t.Errorf("want: %s, got: %s", want, got)
			}

			ev = NewEnvValueWithDefault(test.key, test.explain, test.defValue)
			want, got = test.key, ev.Key()
			if got != want {
				t.Errorf("want: %s, got: %s", want, got)
			}

			want, got = test.value, ev.ValueProvider()()
			if got != want {
				t.Errorf("want: %s, got: %s", want, got)
			}

			want, got = test.explain, ev.Explain()
			if got != want {
				t.Errorf("want: %s, got: %s", want, got)
			}

			want, got = test.defValue, ev.DefValue()
			if got != want {
				t.Errorf("want: %s, got: %s", want, got)
			}

			os.Unsetenv(test.key)

			// if the env is unset, Value() returns defValue
			want, got = test.defValue, ev.ValueProvider()()
			if got != want {
				t.Errorf("want: %s, got: %s", want, got)
			}
		})
	}
}
