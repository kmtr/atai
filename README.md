# Atai (値) [![CircleCI](https://circleci.com/gh/kmtr/atai.svg?style=svg)](https://circleci.com/gh/kmtr/atai)

Atai is a getting value library.

## Description

Someimes we want to get values from any places (environment variable, command line argument, DB, etc...).
There are many library for its purpose, like os.Getenv.
Atai is a wrapper of these function.

### ValueProvider

ValueProvider is the core concept of this library.
It is a type alias of the function (`func() string`).
It can make abstraction that getting a value.

Sample

```go
package main

import (
	"database/sql"
	"database/sql/driver"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kmtr/atai"
)

var db *sql.DB

type DummyDB struct{}

func (db *DummyDB) Open(name string) (driver.Conn, error) {
	return nil, fmt.Errorf("DummyDB: open with `%s`", name)
}

func initialize(dbconn atai.ValueProvider) (err error) {
	sql.Register("dummydb", &DummyDB{})
	db, err = sql.Open("dummydb", dbconn())
	return
}

func main() {
	flag.String("db-conn", "", "Database connection string")
	flag.Parse()

	mvp := atai.MultipleValueProvider(
		atai.ValueFromFlag("db-conn"),
		atai.ValueFromEnv("DB_CONN"),
		atai.Value("default connection string"),
	)
	if err := initialize(mvp); err != nil {
		log.Printf("error: %v", err)
		os.Exit(1)
	}
	log.Printf("info: %v", db.Ping())
}
```

### interface (ValueProviderHolder, KeyHolder, DefaultValueHolder, Explainer)

They are utility interfaces for using ValueProvider.

Sample

```go
package main

import (
	"flag"
	"fmt"

	"github.com/kmtr/atai"
)

type KeyHolderAndExplainer interface {
	atai.KeyHolder
	atai.Explainer
}

func explain(vals ...KeyHolderAndExplainer) {
	for i, val := range vals {
		fmt.Printf("%d: %s, %s\n", i+1, val.Key(), val.Explain())
	}
}

func main() {
	flag.String("db-conn", "", "Database connection string")
	flag.Parse()

	fv := atai.NewFlagValue("db-conn")
	ev := atai.NewEnvValue("DB_CONN", "Database connection string")

	explain(fv, ev)
}
```
