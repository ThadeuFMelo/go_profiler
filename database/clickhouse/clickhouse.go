package clickhouse

import (
	"time"

	"github.com/uptrace/go-clickhouse/ch"
	"github.com/uptrace/go-clickhouse/chdebug"
)

// Clickhouse struct
type Clickhouse struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

// Connect to clickhouse
func Connect() *ch.DB {
	db := ch.Connect(
		ch.WithAddr("localhost:9000"),
		ch.WithDatabase("helloworld"),
		ch.WithTimeout(5*time.Second),
		ch.WithDialTimeout(5*time.Second),
		ch.WithReadTimeout(5*time.Second),
		ch.WithWriteTimeout(5*time.Second),
	)

	db.AddQueryHook(chdebug.NewQueryHook(
		chdebug.WithVerbose(true),
		chdebug.FromEnv("CHDEBUG"),
	))
	return db
}
