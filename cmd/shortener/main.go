package main

import (
	"github.com/ChristinaFomenko/shortener/configs"
	"github.com/ChristinaFomenko/shortener/internal/app/generator"
	repositoryURL "github.com/ChristinaFomenko/shortener/internal/app/repository/urls"
	serviceURL "github.com/ChristinaFomenko/shortener/internal/app/service/urls"
	"github.com/ChristinaFomenko/shortener/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

//func init() {
//	// loads values from .env into the system
//	if err := godotenv.Load(); err != nil {
//		log.Print("No .env file found")
//	}
//}
//
//type Config struct {
//	ServerAddress string `env:"SERVER_ADDRESS"`
//	BaseURL       string `env:"BASE_URL"`
//}

func main() {

	// Repositories
	repository := repositoryURL.NewRepo()

	// Services
	helper := generator.NewGenerator()
	service := serviceURL.NewService(repository, helper, configs.BaseURL())

	// Route
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	//router.Route("/", func(r chi.Router) {
	router.Post("/", handlers.New(service).Shorten)
	router.Get("/{id}", handlers.New(service).Expand)
	router.Post("/api/shorten", handlers.New(service).APIJSONShortener)
	//})
	port := configs.ServerAddress()

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(port, router))
}
