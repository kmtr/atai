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

var db *sql.DB

type DummyDB struct{}

func (db *DummyDB) Open(name string) (driver.Conn, error) {
	return nil, fmt.Errorf("DummyDB: open with `%s`", name)
}

func initialize(mvp atai.ValueProvider) error {
	dbconn := mvp()
	sql.Register("dummydb", &DummyDB{})
	var err error
	db, err = sql.Open("dummydb", dbconn)
	if err != nil {
		return err
	}
	return nil
}
