package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/itishrishikesh/gofeed/internal/database"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	portNumber := os.Getenv("PORT")
	if portNumber == "" {
		log.Fatal("E#1OLYW6 - PORT is not found in the env variables")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("E#1ORGYA - DB_URL is not found in the env variables")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("E#1ORHEU - Failed to open a connection to db")
	}

	db := database.New(conn)

	apiCfg := apiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute)

	router := chi.NewRouter()

	log.Println("I#1OLYVV - Server starting on port number", portNumber)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", healthCheckHandler)
	v1Router.Get("/err", errorHandler)
	v1Router.Post("/users", apiCfg.createUserHandler)
	v1Router.Get("/users", apiCfg.authMiddleware(apiCfg.getUserHandler))
	v1Router.Post("/feeds", apiCfg.authMiddleware(apiCfg.createFeedHandler))
	v1Router.Get("/feeds", apiCfg.getFeedsHandler)
	v1Router.Post("/feedfollow", apiCfg.authMiddleware(apiCfg.createFeedFollowHandler))
	v1Router.Get("/feedfollow", apiCfg.authMiddleware(apiCfg.getFeedFollowsHandler))
	v1Router.Delete("/feedfollow/{feedFollowId}", apiCfg.authMiddleware(apiCfg.deleteFeedHandler))
	v1Router.Get("/posts", apiCfg.authMiddleware(apiCfg.getPostsForUserHandler))
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portNumber,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln("E#1OLYWJ - Server failed to launch. E:", err)
	}

}
