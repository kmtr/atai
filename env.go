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
// If it does not have a value then return a defValue.
func ValueFromEnvWithDefault(envName, defValue string) ValueProvider {
	return func() string {
		value, exists := os.LookupEnv(envName)
		if exists {
			return value
		}
		return defValue
	}
}

type EnvValue struct {
	key      string
	vp       ValueProvider
	defValue string
}

func NewEnvValue(key string) *EnvValue {
	return NewEnvValueWithDefault(key, "")
}

func NewEnvValueWithDefault(key, defValue string) *EnvValue {
	return &EnvValue{
		key:      key,
		vp:       ValueFromEnvWithDefault(key, defValue),
		defValue: defValue,
	}
}

func (ev *EnvValue) Key() string {
	return ev.key
}

func (ev *EnvValue) Value() string {
	return ev.vp()
}

func (ev *EnvValue) DefValue() string {
	return ev.defValue
}
