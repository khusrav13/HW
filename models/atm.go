package models

import (
	"Homework/db"
	"database/sql"
	"log"
)

type ATMs struct {
	ID int64
	Name string
	Status bool
}

func AddATM(DB *sql.DB, address string) (ok bool, err error) {

	_, err = DB.Exec(db.AddNewATM, address)
	if err != nil {
		log.Println("Cannot insert to ATMs table new address, error is ", err)
		return false, err
	}
	return true, nil
}

