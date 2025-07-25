package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/a-rezk/rssagg/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load("./.env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the enviroment!")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the enviroment")
	}

	conn, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database!: ", err)
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: false,
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		MaxAge:           300,
	}))

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErrors)
	v1Router.Post("/users", apiCfg.createUserHandler)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.createFeedHandler))
	v1Router.Get("/feeds", apiCfg.GetFeedsHandler)
	v1Router.Delete("/feeds/{feedID}", apiCfg.middlewareAuth(apiCfg.DeleteFeedHandler))
	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.createFeedFollowHandler))
	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.GetFeedFollowsHandler))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.DeleteFeedFollowHandler))
	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetPostForUser))
	router.Mount("/v1", v1Router)
	log.Printf("Server Starting on Port: %s", portString)
	sErr := srv.ListenAndServe()
	if sErr != nil {
		log.Fatal(sErr)
	}

}
