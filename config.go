// package mysql
//
// Copyright(C)2020 Micael Treanor
//
// Requirements:
// uses github.com/go-sql-driver/mysql which requires
// MySQL (4.1+), MariaDB, Percona Server, Google CloudSQL or Sphinx (2.2.3+)
//

package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	// mySqlUserVariable and mySqlPassword are the names of the environment variables
	// used to store connection information.
	mySqlUserName = "MYSQL_USERNAME"
	mySqlPassword = "MYSQL_PASSWORD"

	defaultMySQLHost = "localhost" // defaults for localhost are most secure
	defaultMySQLPort = "33060"     // depending on the MySQL version; this may need to be 3306

	// this is the 'driver name' used by helper functions that smooth out connections
	mySqlDriverName = "mysql"
)

// NewDBConfig returns a new MySQL database connection configuration object.
func NewMySQL() (MySQL, error) {
	username := os.Getenv(mySqlUserName)
	if username == "" {
		return nil, fmt.Errorf("environment variable %s for MySQL username not found", mySqlUserName)
	}
	password := os.Getenv(mySqlPassword)
	if password == "" {
		return nil, fmt.Errorf("environment variable %s for MySQL password not found", mySqlPassword)
	}
	d := new(mySQL)
	d.username = username
	d.password = password

	return d, nil
}

// MySQL defines the interface to the MySQL database connection
type MySQL interface {
	Auth() string
	DSN(database string) string
	Open(dbname string) (*sql.DB, error)
	Load(file string) error
	Save(file string) error
}

type mySQL struct {
	username string
	password string
	host     string `default:"localhost"` // defaults for localhost are most secure
	port     string `default:"33060"`     // depending on the MySQL version; this may need to be 3306
	logging  bool   `default:"false"`
}

// Open opens a database specified by its database driver name and a driver-specific data source name, usually consisting of at least a database name and connection information.
//
// Most users will open a database via a driver-specific connection helper function that returns a *DB. No database drivers are included in the Go standard library. See https://golang.org/s/sqldrivers for a list of third-party drivers.
//
// Open may just validate its arguments without creating a connection to the database. To verify that the data source name is valid, call Ping.
//
// The returned DB is safe for concurrent use by multiple goroutines and maintains its own pool of idle connections. Thus, the Open function should be called just once. It is rarely necessary to close a DB.
func (db mySQL) Open(dbname string) (*sql.DB, error) {

	dbconnection, err := sql.Open(mySqlDriverName, db.DSN(dbname))

	if err != nil {
		return nil, err
	}

	// See "Important settings" section.
	dbconnection.SetConnMaxLifetime(time.Minute * 3)
	dbconnection.SetMaxOpenConns(10)
	dbconnection.SetMaxIdleConns(10)
	return dbconnection, nil
}

// DSN returns the entire DSN authentication string including a database name.
// Using "" for the database name will return a generic connection to the server
// that allows listing and choosing different database names.
func (db mySQL) DSN(database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s/%s)", db.username, db.password, db.host, db.port, database)
}

// Load loads the database configuration from a json file
//
// Not Implemented
func (db mySQL) Load(file string) error {
	// load json config file
	return NotImplemented
}

// Load saves the database configuration to a json file
//
// Not Implemented
func (db mySQL) Save(file string) error {
	// save json config file
	return NotImplemented
}

// NotImplemented returns an error if the method is not yet implemented
var NotImplemented error = errors.New("not implemented")
