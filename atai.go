package atai

// ValueProvider is an alias of `func() string`
type ValueProvider func() string

// Value returns a ValueProvider that provide a value of argument.
func Value(v string) ValueProvider {
	return func() string {
		return v
	}
}

type ValueHolder interface {
	Value() string
}

type KeyHolder interface {
	Key() string
}

type DefaultValueHolder interface {
	DefValue() string
}

type Explainer interface {
	Explain() string
}
