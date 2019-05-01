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

type FlagValue struct {
	key string
	vp  ValueProvider
}

func NewFlagValue(key string) *FlagValue {
	return &FlagValue{
		key: key,
		vp:  ValueFromFlag(key),
	}
}

func (fv *FlagValue) Key() string {
	return fv.key
}

func (fv *FlagValue) Value() string {
	return fv.vp()
}

func (fv *FlagValue) DefValue() string {
	return flag.Lookup(fv.key).DefValue
}
