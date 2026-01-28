package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/karsinkk/108-Hackathon-Backend/controllers/admincontroller"
	"github.com/karsinkk/108-Hackathon-Backend/controllers/usercontroller"
	"github.com/karsinkk/108-Hackathon-Backend/controllers/vehiclecontroller"
	"github.com/karsinkk/108-Hackathon-Backend/dif"
	"github.com/karsinkk/108-Hackathon-Backend/helpers"
)

func main() {
	// Initialize structured logging
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	// Load configuration
	config := dif.GetConfig()

	// Initialize database connection pool
	_ = dif.GetDB()
	defer dif.CloseDB()

	// Create router
	router := chi.NewRouter()

	// Middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	// CORS configuration
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{config.AllowedOrigins},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link", "108"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Health check endpoint
	router.Get("/health", healthCheck)

	// Vehicle routes
	router.Route("/vehicle", func(r chi.Router) {
		r.Post("/login", vehiclecontroller.LoginVehicle)
		r.Get("/update", vehiclecontroller.UpdateVehicle)
		r.Get("/notification", vehiclecontroller.UpdateVehicle)
		r.Post("/finish", vehiclecontroller.Finish)
	})

	// User routes
	router.Route("/user", func(r chi.Router) {
		r.Post("/emergency", usercontroller.Emergency)
	})

	// Admin routes
	router.Route("/admin", func(r chi.Router) {
		r.Post("/register", admincontroller.Register)
		r.Get("/seen", admincontroller.ModifySeen)
		r.Post("/addvehicle", admincontroller.AddVehicle)
		r.Post("/login", admincontroller.Login)
		r.Get("/notification", admincontroller.Notification)
		r.Get("/emergency", admincontroller.DisplayEmergency)
		r.Get("/emergencycount", admincontroller.CountEmergency)
		r.Post("/status", admincontroller.Status)
		r.Post("/dismiss", admincontroller.DismissEmergency)
		r.Post("/dismissvehicle", admincontroller.DismissVehicle)
		r.Get("/ambulance", admincontroller.DisplayAmbulance)
		r.Get("/firepolice", admincontroller.DisplayFirePolice)
	})

	// Create server with graceful shutdown
	server := &http.Server{
		Addr:         ":" + config.ServerPort,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in goroutine
	go func() {
		log.Info().
			Str("port", config.ServerPort).
			Msg("Starting server")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Server failed to start")
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server exited gracefully")
}

// healthCheck handles the health check endpoint
func healthCheck(w http.ResponseWriter, r *http.Request) {
	db := dif.GetDB()
	dbStatus := "healthy"

	if err := db.Ping(); err != nil {
		dbStatus = "unhealthy"
	}

	response := helpers.HealthResponse{
		Status:    "ok",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Database:  dbStatus,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
