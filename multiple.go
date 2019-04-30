package atai

// MultipleValueProvider takes multiple ValueProviders and returns a ValueProvider
// that containes these ValueProviders.
// Each contained ValueProviders is called in order until one of them returns not empty("") value.
// If all ValueProviders return empty(""), finally return empty("").
func MultipleValueProvider(providers ...ValueProvider) ValueProvider {
	return MultipleValueProviderWithDefault("", providers...)
}

// MultipleValueProviderWithDefault takes a defaultValue and multiple ValueProviders
// and returns that included their ValueProviders.
// Each contained ValueProviders is called in order until one of them returns not empty("") value.
// If all ValueProviders return empty(""), finally return defaultValue.
func MultipleValueProviderWithDefault(defaultValue string, providers ...ValueProvider) ValueProvider {
	return ValueProvider(func() string {
		for _, p := range providers {
			if v := p(); v != "" {
				return v
			}
		}
		return defaultValue
	})
}
