package main

import (
	"firstProject/internal/database"
	"fmt"
	players "firstProject/basket/players/web"
)

const (
	mysqlConnStr = "%s:%s@tcp(%s:33006)/%s?parseTime=true"
)

func main() {

	sqlClient := database.NewSqlClient(fmt.Sprintf(mysqlConnStr, "root", "root-go", "localhost",
		"basket"))

	mux := routes(
		players.NewPlayerHandler(sqlClient),
	)

	server := NewServer(mux)
	server.Run()
}
