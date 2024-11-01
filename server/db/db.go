package db

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)




func Connect () (*sql.DB,error){

	db, err := sql.Open (
		"mysql",
		"root:ianbore12@tcp(127.0.0.1:3306)/missing_persons",
	)

	if err != nil {
		panic( err.Error())
	}

	err =db.Ping()


	if err != nil {

		return nil , err
	}


	
	fmt.Println ("Successfully connected to the database")


	return db, nil
	
}