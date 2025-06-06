package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	cwd, _ := os.Getwd()
	fmt.Println("Current directory:", cwd)

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file", envErr)
	}

	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var version string
	if err := db.QueryRow("select version()").Scan(&version); err != nil {
		panic(err)
	}

	fmt.Printf("version=%s\n", version)

	fmt.Printf("change")
}
