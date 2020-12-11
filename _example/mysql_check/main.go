package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := mysql.Check()
	if err != nil {
		log.Fatal(err)
	}
}
