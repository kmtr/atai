package atai

// ValueProvider is an alias of `func() string`
type ValueProvider func() string

// Value returns a ValueProvider that provide a value of argument.
func Value(v string) ValueProvider {
	return func() string {
		return v
	}
}

// ValueProviderHolder is the interface that returns a ValueProvider.
type ValueProviderHolder interface {
	ValueProvider() ValueProvider
}

// KeyHolder is the interface that returns the key of the value.
// For example, command line argument, environment value name.
type KeyHolder interface {
	Key() string
}

// DefaultValueHolder is the interface that returns a default value.
type DefaultValueHolder interface {
	DefValue() string
}

// Explainer is the interface that returns an explain of the value.
type Explainer interface {
	Explain() string
}
