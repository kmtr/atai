# Atai (値) [![CircleCI](https://circleci.com/gh/kmtr/atai.svg?style=svg)](https://circleci.com/gh/kmtr/atai)

Atai is a getting value library.

## Description

Someimes we want to get values from any places (environment variable, command line argument, DB, etc...).
There are many library for its purpose, like os.Getenv.
Atai is a wrapper of these function.

### ValueProvider

ValueProvider is the core concept of this library.
Its definition is very simple.

```go
type ValueProvider func() string
```

Atai package has some types of ValueProvider.
However it is easy to make your ValueProvider.
For example.

```go
yourValueProvider := ValueProvider(func() string {
    return "This is my ValueProvider"
})
```

### Value

`Value` is very simple. Its returns a argument of its.

```go
fmt.Print(Value("val")()) // val 
```

### ValueFromEnv (ValueFromEnvWithDefault)

`ValueFromEnv` is for getting value from environment variables.
`ValueFromEnvWithDefault` is for too and this can be set default value.

```go
os.Setenv("VALUE", "env val")
fmt.Print(ValueFromEnv("VALUE")()) // env val 
```
### ValueFromFlag

`ValueFromFlag` is for getting value from command line argument with `flag` package.

```go
flag.String("value", "", "command line argument --value")
flag.Set("value", "flag val")
flag.Parse()
fmt.Print(ValueFromFlag("value")()) // flag val
```
