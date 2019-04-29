package atai

func MultipleValueProvider(providers ...ValueProvider) ValueProvider {
	return ValueProvider(func() string {

		for _, p := range providers {
			if v := p(); v != "" {
				return v
			}
		}
		return ""
	})
}
