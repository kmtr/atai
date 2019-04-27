package atai

import "os"

// ValueFromEnv is a function which get a value from env.
func ValueFromEnv(envName string) ValueProvider {
	return func() string {
		value, exists := os.LookupEnv(envName)
		if exists {
			return value
		}
		return ""
	}
}

// ValueFromEnvWithDefault returns a ValueProvider that provide a value from environment.
// If it does not have a value then return a defaultValue.
func ValueFromEnvWithDefault(envName, defaultValue string) ValueProvider {
	return func() string {
		value, exists := os.LookupEnv(envName)
		if exists {
			return value
		}
		return defaultValue
	}
}
