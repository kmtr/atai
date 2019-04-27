package atai

import (
	"os"
	"testing"
)

func TestValueFromEnv(t *testing.T) {
	type testarg struct {
		envName      string
		defaultValue string
		envValue     string
		want         string
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

			got = ValueFromEnvWithDefault(test.envName, test.defaultValue)()
			if test.envName == "not registered" {
				if test.defaultValue != got {
					t.Errorf("want: %s, got: %s", test.defaultValue, got)
				}
			} else {
				if test.want != got {
					t.Errorf("want: %s, got: %s", test.want, got)
				}
			}
		})
	}
}
