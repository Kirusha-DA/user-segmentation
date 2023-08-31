package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Kirusha-DA/user-segmentation/internal/dbs/postgres"
	"github.com/Kirusha-DA/user-segmentation/internal/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("load .env failed")
	}

	db, err := postgres.New(postgres.Config{
		DB:       os.Getenv("DB_NAME"),
		Host:     "localhost",
		Port:     5432,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("psql connection error: %s", err.Error())
	}

	h := handlers.New(db)
	router := mux.NewRouter()

	router.HandleFunc("/segments", h.AddSegment).Methods(http.MethodPost)
	router.HandleFunc("/segments/{slug}", h.DeleteSegment).Methods(http.MethodDelete)
	router.HandleFunc("/users/{id}/segments", h.AddSegmentsToUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}/segments", h.DeleteUserSegments).Methods(http.MethodDelete)
	router.HandleFunc("/users/{id}/segments", h.GetUserSegments).Methods(http.MethodGet)

	log.Println("API is running!")
	http.ListenAndServe(":8080", router)
}
