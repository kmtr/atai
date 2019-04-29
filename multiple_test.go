package atai

import "testing"

func TestMultipleValue(t *testing.T) {
	type testarg struct {
		providers []ValueProvider
		want      string
	}

	tests := []testarg{
		{
			providers: []ValueProvider{Value("0"), Value("1"), Value("2")},
			want:      "0",
		},
		{
			providers: []ValueProvider{Value(""), Value("1"), Value("2")},
			want:      "1",
		},
		{
			providers: []ValueProvider{},
			want:      "",
		},
	}

	for _, test := range tests {
		t.Run(test.want, func(t *testing.T) {
			got := MultipleValueProvider(test.providers...)()
			if test.want != got {
				t.Errorf("want: %s, got: %s", test.want, got)
			}
		})
	}
}
