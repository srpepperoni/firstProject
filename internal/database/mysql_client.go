package database

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type MySqlClient struct {
	*sql.DB
}

func NewSqlClient(source string) *MySqlClient {
	db, err := sql.Open("mysql", source)

	if err != nil {
		log.Printf("cadena: %s", source)
		log.Fatal("cannot create db tentat: %s", err.Error())
		panic("..")
	}

	err = db.Ping()

	if err != nil {
		log.Println("cannot connect to mysql!")
	}

	return &MySqlClient{db}
}
