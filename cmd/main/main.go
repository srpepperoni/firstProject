package main

import (
	players "firstProject/basket/players/web"
	"firstProject/internal/database"
	logs "firstProject/internal/log"
	"flag"
	"fmt"
	_ "github.com/golang-migrate/migrate/source/file"
	"os"
)

const mysqlConnStr = "%s:%s@tcp(%s:%s)/%s?parseTime=true"

func main() {
	environment := initEnv()
	logs.InitDefault(environment.env)

	sqlClient := database.NewSqlClient(fmt.Sprintf(mysqlConnStr, environment.dbUser, environment.dbPass, environment.dbUrl,
		environment.dbPort, environment.dbName))

	mux := routes(
		players.NewPlayerHandler(sqlClient),
	)

	server := NewServer(mux)
	server.Run()
}

type envSetup struct {
	dbUser string
	dbPass string
	dbUrl  string
	dbName string
	dbPort string
	port   string
	env    string
}

func initEnv() *envSetup {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbURL := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	env := flag.String("E", "local", "the current execution environment")
	flag.Parse()

	if dbPass == "" && dbUser == "" && dbURL == "" && dbName == "" && dbPort == "" {
		dbUser, dbPass, dbURL, dbName, dbPort = "root", "root-go", "localhost", "basket", "33006"
	}

	serverPort := os.Getenv("PORT")

	if serverPort == "" {
		serverPort = "9000"
	}

	return &envSetup{
		dbUser: dbUser,
		dbPass: dbPass,
		dbUrl:  dbURL,
		dbName: dbName,
		dbPort: dbPort,
		port:   serverPort,
		env:    *env,
	}
}
