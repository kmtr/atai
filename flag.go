package atai

import (
	"flag"
	"fmt"
)

// ValueFromFlag returns a ValueProvider that provide a value from `flag`.
func ValueFromFlag(name string) ValueProvider {
	return func() string {
		f := flag.Lookup(name)
		if f == nil {
			return ""
		}
		return fmt.Sprint(f.Value)
	}
}
