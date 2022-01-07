package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/ttakuya50/go-architecture-sample/api/domain/repository"
	"github.com/ttakuya50/go-architecture-sample/api/handler"
	"github.com/ttakuya50/go-architecture-sample/api/service"
)

func main() {
	router, cleanupFunc, err := register()
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":8888", router); err != nil {
		cleanupFunc()
		log.Fatal(err)
	}
}

type CleanupFunc func()

func register() (*mux.Router, CleanupFunc, error) {
	conn, err := mysqlClient()
	if err != nil {
		return nil, nil, err
	}

	r := repository.NewRandom()
	db := repository.NewDB(conn)

	userRepo := repository.NewUserRepo()
	taskRepo := repository.NewTaskRepo()
	listRepo := repository.NewListRepo()

	// service
	userService := service.NewUserService(db, userRepo, r, taskRepo, listRepo)

	defaultApiService := handler.NewDefaultApiService(userService)
	defaultApiController := handler.NewDefaultApiController(defaultApiService)
	router := handler.NewRouter(defaultApiController)

	return router, func() {
		conn.Close()
	}, nil
}

func mysqlClient() (*sql.DB, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", "root", "root", "127.0.0.1", "mydb")

	// open database
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, fmt.Errorf("failed to database connection: %v", err)
	}

	return db, nil
}
