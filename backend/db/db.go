// db/db.go
package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB(dataSourceName string) {
	var err error
	DB, err = sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		log.Fatalln("Database connection failed:", err)
	}
}
