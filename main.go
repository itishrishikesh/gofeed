package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	portNumber := os.Getenv("PORT")
	if portNumber == "" {
		log.Fatal("E#1OLYW6 - PORT is not found in the env variables")
	}

	router := chi.NewRouter()

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portNumber,
	}

	log.Println("I#1OLYVV - Server starting on port number", portNumber)

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Mount("/v1", router)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalln("E#1OLYWJ - Server failed to launch. E:", err)
	}

}
