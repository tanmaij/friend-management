package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/tanmaij/friend-management/cmd/router"
	relationshipCtrl "github.com/tanmaij/friend-management/internal/controller/relationship"
	userCtrl "github.com/tanmaij/friend-management/internal/controller/user"
	"github.com/tanmaij/friend-management/internal/handler"
	relationshipRepo "github.com/tanmaij/friend-management/internal/repository/relationship"
	userRepo "github.com/tanmaij/friend-management/internal/repository/user"
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

	defer sqlDB.Close()
	log.Println("successfully connected to database")

	relationshipRepoInstance := relationshipRepo.New(sqlDB)
	userRepoInstance := userRepo.New(sqlDB)

	relationshipCtrlInstance := relationshipCtrl.New(relationshipRepoInstance, userRepoInstance)
	userCtrlInstance := userCtrl.New(userRepoInstance)

	handlerInstance := handler.New(relationshipCtrlInstance, userCtrlInstance)

	r := router.InitRouter(handlerInstance)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		log.Printf("received signal: %s. shutting down...\n", sig)
		os.Exit(0)
	}()

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
