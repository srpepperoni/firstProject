package database

import (
	"database/sql"
	logs "firstProject/internal/log"
	_ "github.com/go-sql-driver/mysql"
)

type MySqlClient struct {
	*sql.DB
}

func NewSqlClient(source string) *MySqlClient {
	db, err := sql.Open("mysql", source)

	if err != nil {
		logs.Sugar().Info("cadena: #{source}")
		logs.Log().Error("cannot create db tentat: %s")
		panic("..")
	}

	err = db.Ping()

	if err != nil {
		logs.Log().Info("cannot connect to mysql!")
	}

	return &MySqlClient{db}
}
