package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"tech-test-azura-lab/repository"
	"tech-test-azura-lab/router"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "azura_test"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error while loading .env file")
	}

	PG_PASS := os.Getenv("POSTGRES_PASS")

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, PG_PASS, dbname)
	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		log.Fatalln("Error while connecting to DB : ", err.Error())
	}

	productRepo := repository.NewProductRepostory(db)
	mux := router.GenerateMux(productRepo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		return
	}
	log.Println("Server run on port ", port)
	http.ListenAndServe(":"+port, mux)
}
