package main

import (
	"database/sql"
	"example/web-service-gin/cmd/api"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/microsoft/go-mssqldb"
)

func InitEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}
}

func GetEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
func main() {

	user := GetEnv("User", "sa")
	passwd := GetEnv("Passwd", "YourStrong!Passw0rd")
	addr := GetEnv("Addr", "127.0.0.1:1433")
	dbName := GetEnv("DBName", "GOTask")

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
