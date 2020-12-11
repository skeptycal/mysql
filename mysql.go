package mysql

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func dbConnect(dbconfig dbConfig, database string) (*sql.DB, error) {

	// mysqlConnectionString, err := getEnvConnectionString(mySqlUserVariable)
	mysqlConnectionString := dbconfig.Auth()

	//!   INSECURE TODO  ------------------------------------> INSECURE ------>> REMOVE
	log.Info("mysql username: ", mysqlConnectionString)

	// Open database connection.
	db, err := sql.Open("mysql", mysqlConnectionString)
	if err != nil {
		return nil, err
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}

// Check performs a connection check on the mysql database connection
func Check() error {
	dbconfig, err := NewDBConfig()
	if err != nil {
		return err
	}

	// err = dbconfig.Save()
	// if err != nil {
	// 	return err
	// }

	db, err := dbConnect(dbconfig, "")
	if err != nil {
		return err
	}

	// defer the close until  the main function is done
	// This is not normal in a running application. It is rare to Close a DB, as the DB handle is meant to be long-lived and shared between many goroutines.
	defer db.Close()

	// perform query test
	response, err := db.Query("SHOW DATABASES;")
	if err != nil {
		return err
	}

	log.Info("MySQL query response: ", response)
	return nil

}

// const (
// 	username = "root"
// 	password = "password"
// 	hostname = "127.0.0.1:3306"
// 	dbname   = "test"
// )

// Notes: DSN (Data Source Name)
/* The Data Source Name has a common format, like e.g. PEAR DB uses it, but without type-prefix (optional parts marked by squared brackets):
 */
