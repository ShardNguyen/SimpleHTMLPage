package main

import (
	"SimpleHTMLPage/config"
	dbpostgres "SimpleHTMLPage/databases/postgresql"
	dbredis "SimpleHTMLPage/databases/redis"
	"SimpleHTMLPage/routers"
	"context"
)

func main() {
	if err := config.ParseConfig(); err != nil {
		panic(err)
	}
	if err := dbpostgres.UserConnect(); err != nil {
		panic(err)
	}

	defer dbpostgres.CloseUserConnection()

	if err := dbredis.InitTokenStorage(); err != nil {
		panic(err)
	}

	// When server restarts, old tokens are not deleted
	// So have to flush the database to delete all of the active tokens to reset login
	err := dbredis.GetTokenStorage().FlushDB(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	routers.Run()
}
