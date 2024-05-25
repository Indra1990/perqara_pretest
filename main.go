package main

import (
	"context"
	"log"
	"perqara_testing/api"
	"perqara_testing/db"
)

func main() {
	ctx := context.Background()

	conn, dbErr := db.SetupDatabasePgSQLConnection(ctx)
	if dbErr != nil {
		log.Fatal("cannot connect to db :", dbErr)
	}

	server := api.NewServer(conn)
	if startServerErr := server.Start(ctx); startServerErr != nil {
		log.Fatal("cannot start server :", startServerErr)
	}

	if stopServerErr := server.Stop(ctx); stopServerErr != nil {
		log.Fatal("stop server :", stopServerErr)
	}

}
