package atai

import "testing"

func TestValue(t *testing.T) {
	want := "val"
	if got := Value(want)(); got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}
