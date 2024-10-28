package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tanmaij/friend-management/cmd/server/router"
	relationshipCtrl "github.com/tanmaij/friend-management/internal/controller/relationship"
	userCtrl "github.com/tanmaij/friend-management/internal/controller/user"
	"github.com/tanmaij/friend-management/internal/handler"
	relationshipRepo "github.com/tanmaij/friend-management/internal/repository/relationship"
	userRepo "github.com/tanmaij/friend-management/internal/repository/user"
	"github.com/tanmaij/friend-management/pkg/db/sql"
	"github.com/tanmaij/friend-management/pkg/utils/env"
)

type serverConfig struct {
	Port  string
	DBUrl string
}

func initializeServerConfig() serverConfig {
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

	return serverConfig{Port: Port, DBUrl: DBUrl}
}

func main() {
	log.Println("FRIEND MANAGEMENT - API")

	config := initializeServerConfig()

	log.Println("connecting to database...")
	sqlDB, err := sql.ConnectDB(sql.Postgres, config.DBUrl, sql.ConnectionOption{MaxIdleConnections: 10, MaxOpenConnections: 10})
	if err != nil {
		log.Fatalf("failed to open connection to database: %v", err)
	}
	defer sqlDB.Close()

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	if err := sqlDB.PingContext(ctx); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	log.Println("successfully connected to database")

	relationshipRepoInstance := relationshipRepo.New(sqlDB)
	userRepoInstance := userRepo.New(sqlDB)

	relationshipCtrlInstance := relationshipCtrl.New(relationshipRepoInstance, userRepoInstance)
	userCtrlInstance := userCtrl.New(userRepoInstance)

	handlerInstance := handler.New(relationshipCtrlInstance, userCtrlInstance)
	r := router.InitRouter(handlerInstance)

	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: r,
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		log.Printf("received signal: %s. shutting down...\n", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("server shutdown failed: %v", err)
		}

		log.Println("server stopped gracefully")
	}()

	log.Println("server is running on port:", config.Port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error starting server: %v", err)
	}
}
