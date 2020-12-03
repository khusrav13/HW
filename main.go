package main

import (
	"Homework/db"
	"Homework/pkg/core/services"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	DB, err := sql.Open("sqlite3", "test")
		if err != nil {
			fmt.Println("Error", err)
		}
	db.Dbinit(DB)
	Start(DB)

}

func Start(database *sql.DB) {
	for  {
		login, password := services.Authorization(database)
		services.Login(database,login,password)
	}
}