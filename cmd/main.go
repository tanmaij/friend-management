package main

import (
	"log"
	"net/http"

	"github.com/tanmaij/friend-management/cmd/router"
	"github.com/tanmaij/friend-management/internal/controller/relationship"
	"github.com/tanmaij/friend-management/internal/handler"
	"github.com/tanmaij/friend-management/internal/repository"
	"github.com/tanmaij/friend-management/pkg/db/sql"
	"github.com/tanmaij/friend-management/pkg/utils/env"
)

type config struct {
	Port  string
	DBUrl string
}

func main() {
	log.Println("FRIEND MANAGEMENT - API")

	config := getConfig()

	log.Println("connecting to database...")
	sqlDB, err := sql.ConnectDB(sql.Postgres, config.DBUrl, sql.ConnectionOption{MaxIdleConnections: 10, MaxOpenConnections: 10})
	if err != nil {
		log.Fatalf("failed to open connection to database: %v", err)
	}
	log.Println("successfully connected to database")

	repositoryInstance := repository.New(sqlDB)
	relationshipCtrlInstance := relationship.New(repositoryInstance)
	handlerInstance := handler.New(relationshipCtrlInstance)

	r := router.InitRouter(handlerInstance)

	log.Println("server is running on port:", config.Port)
	if err := http.ListenAndServe(":"+config.Port, r); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}

func getConfig() config {
	var (
		Port  = env.Get("PORT")
		DBUrl = env.Get("DB_URL")
	)

	if Port == "" {
		log.Fatal("cannot find PORT in environment variables")
	}

	if DBUrl == "" {
		log.Fatal("cannot find DB_URL in environment variables")
	}

	return config{Port: Port, DBUrl: DBUrl}
}
