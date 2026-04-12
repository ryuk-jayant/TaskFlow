package main

import (
	"database/sql"
	"example/web-service-gin/cmd/api"
	"example/web-service-gin/cmd/config"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/microsoft/go-mssqldb"
)

func main() {
	// config.InitEnv()//for local env
	// log.Println("DB_HOST:", os.Getenv("DB_HOST"))
	log.Println("DB_PORT:", os.Getenv("DB_PORT"))
	user := config.GetEnv("DB_USER", "fall")
	passwd := config.GetEnv("DB_PASSWORD", "fallbackPassw0rd")
	addr := config.GetEnv("DB_ADRR", "localhost")
	dbName := config.GetEnv("DB_NAME", "GOTask")

	log.Println("Config:", user, dbName)
	log.Println("Connecting to SQL Server...")

	connString := fmt.Sprintf(
		"sqlserver://%s:%s@%s?database=%s",
		user,
		passwd,
		addr,
		dbName,
	)

	DB, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error opening DB:", err)
	}

	// Optional but recommended: connection pooling
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)

	// Ping to verify connection
	if err := DB.Ping(); err != nil {
		log.Fatal("DB connection failed:", err)
	}

	log.Println("Connected to SQL Server!")
	server := api.NewApiServer("localhost:8080", DB)
	if err := server.Run(); err != nil {
		log.Fatal("Server Stopped")
	}
}
