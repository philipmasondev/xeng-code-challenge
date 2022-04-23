package driver

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// DB stores connections
type DB struct {
	SQL *sql.DB
}

type DB1 struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenConn = 10
const maxIdleConn = 5
const maxTTL = 5 * time.Minute

// ConnectSQL create db
func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		log.Fatal(err)
	}
	d.SetMaxOpenConns(maxOpenConn)
	d.SetConnMaxIdleTime(maxIdleConn)
	d.SetConnMaxLifetime(maxTTL)

	dbConn.SQL = d

	err = testConn(d)

	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

// testConn test db connection by pinging db
func testConn(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}

// NewDatabase create new db
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
