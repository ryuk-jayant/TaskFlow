package api

import (
	"context"
	"database/sql"
	"example/web-service-gin/service/project"
	"example/web-service-gin/service/user"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	Addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		Addr: addr,
		db:   db,
	}
}

// func (s *ApiServer) Run() error {
// 	router := mux.NewRouter()
// 	subRouter := router.PathPrefix("/api/v1").Subrouter()
// 	userStore := user.NewStore(s.db)
// 	userRoutes := user.NewHandler(userStore)
// 	userRoutes.RegisterRoutes(subRouter)
// 	projectStore:=project.NewStore(s.db)
// 	projectRoutes:=project.NewHandler(projectStore)
// 	projectRoutes.RegisterRoutes(subRouter)


// 	log.Printf("Server running on %v \n", s.Addr)
// 	return http.ListenAndServe(s.Addr, router)
// }
func (s *ApiServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userRoutes := user.NewHandler(userStore)
	userRoutes.RegisterRoutes(subRouter)

	projectStore := project.NewStore(s.db)
	projectRoutes := project.NewHandler(projectStore)
	projectRoutes.RegisterRoutes(subRouter)

	// Create server instance
	server := &http.Server{
		Addr:    s.Addr,
		Handler: router,
	}

	// Start server in goroutine
	go func() {
		log.Printf("Server running on %v\n", s.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen error: %v", err)
		}
	}()

	// Listen for shutdown signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop // block until signal received

	log.Println("Shutting down server...")

	// Timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Forced shutdown: %v", err)
		return err
	}

	// Close DB connection
	if err := s.db.Close(); err != nil {
		log.Printf("Error closing DB: %v", err)
	}

	log.Println("Server exited cleanly")
	return nil
}
