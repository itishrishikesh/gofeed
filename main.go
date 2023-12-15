package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/itishrishikesh/gofeed/controller"
	"github.com/itishrishikesh/gofeed/internal/database"
	"github.com/itishrishikesh/gofeed/middleware"
	"github.com/itishrishikesh/gofeed/utils"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	portNumber := os.Getenv("PORT")
	if portNumber == "" {
		log.Fatal("E#1OLYW6 - PORT is not found in the env variables")
	}

	dbUrl := os.Getenv("GO_FEED_DB_URL")
	if dbUrl == "" {
		log.Fatal("E#1ORGYA - DB_URL is not found in the env variables")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("E#1ORHEU - Failed to open a connection to db")
	}

	db := database.New(conn)

	apiCfg := middleware.ApiConfig{
		DB: db,
	}

	go utils.StartScraping(db, 10, time.Minute)

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

	config := controller.ApiConfig{
		DB: db,
	}

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", controller.HealthCheckHandler)
	v1Router.Get("/err", controller.ErrorHandler)
	v1Router.Post("/users", config.CreateUserHandler)
	v1Router.Get("/users", apiCfg.AuthMiddleware(config.GetUserHandler))
	v1Router.Post("/feeds", apiCfg.AuthMiddleware(config.CreateFeedHandler))
	v1Router.Get("/feeds", config.GetFeedsHandler)
	v1Router.Post("/feedfollow", apiCfg.AuthMiddleware(config.CreateFeedFollowHandler))
	v1Router.Get("/feedfollow", apiCfg.AuthMiddleware(config.GetFeedFollowsHandler))
	v1Router.Delete("/feedfollow/{feedFollowId}", apiCfg.AuthMiddleware(config.DeleteFeedHandler))
	v1Router.Get("/posts", apiCfg.AuthMiddleware(config.GetPostsForUserHandler))
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
