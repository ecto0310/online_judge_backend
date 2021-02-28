package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

// Init the connection to the database
func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	Db, err = sql.Open("mysql", dbConnection)
	if err != nil {
		panic(err)
	}
}
