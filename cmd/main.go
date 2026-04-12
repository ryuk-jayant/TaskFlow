package main

import (
	"database/sql"
	"example/web-service-gin/cmd/api"
	"example/web-service-gin/cmd/config"
	"fmt"
	"log"
	"time"

	_ "github.com/microsoft/go-mssqldb"
)


func main() {
	config.InitEnv()
	user := config.GetEnv("User","fall")
	passwd := config.GetEnv("Passwd", "fallbackPassw0rd")
	addr := config.GetEnv("Addr", "localhost")
	dbName := config.GetEnv("DBName", "GOTask")

	log.Println("Config:", user, dbName)
	log.Println("Connecting to SQL Server...")

	connString := fmt.Sprintf(
		"sqlserver://%s:%s@%s?database=%s",
		user,
		passwd,
		addr, // example: localhost:1433
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
	server := api.NewApiServer("localhost:9020", DB)
	if err := server.Run(); err != nil {
		log.Fatal("Server Stopped")
	}
}
