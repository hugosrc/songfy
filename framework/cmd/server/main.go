package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/hugosrc/songfy/framework/storage"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	repos, err := storage.NewRepositories(os.Getenv("APP_ENVIRONMENT"))
	if err != nil {
		panic(err)
	}
	defer repos.Close()
	repos.AutoMigrate()

	router := mux.NewRouter()

	server := &http.Server{
		Addr:         "localhost:3000",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
