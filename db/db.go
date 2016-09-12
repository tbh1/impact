package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/Sirupsen/logrus"
)

type DBType int
type Config struct {
	Url  string
	Type DBType
}

func (e DBType) String() string {
	switch e {
	case MySQLType:
		return "mysql"
	case PostgreSQLType:
		return "postgres"
	case SQLite3Type:
		return "sqlite3"
	}
	return ""
}

const (
	MySQLType DBType = iota
	PostgreSQLType
	SQLite3Type
)

var (
	config Config
	dbConn *sql.DB
)

func Initialize(c Config) {
	// Prepare DB connection
	config = c
	db, err := sql.Open(c.Type.String(), c.Url)
	if err != nil {
		return "", err
	}
	dbConn = db
	// Test DB connection immediately and exit on error
	msg, err := HealthCheck()
	if err != nil {
		log.Fatal(err)
	}
	log.Info(msg)
}

func HealthCheck() (string, error) {
	err := dbConn.Ping()
	if err != nil {
		return "", err
	}
	return "Connection successful", nil
}
