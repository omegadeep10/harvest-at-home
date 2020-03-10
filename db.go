import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var database

func GetDb() {
	database, _ = sql.Open("sqlite3", "./h.db")
	return database
}