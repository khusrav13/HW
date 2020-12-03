package db

import (
	"database/sql"
	"log"
)

func Dbinit(db *sql.DB)  {
	DDLs := []string{CreateUsersAccount, CreateATMsTable}
	for _, ddl := range DDLs {
		_, err := db.Exec(ddl)
		if err != nil {
			log.Fatalf("Cannot init %s, err is %e", ddl, err)
		}
	}
}