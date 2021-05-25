package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/thebigyovadiaz/api-twitter/middlewares"
	"github.com/thebigyovadiaz/api-twitter/routes"
)

/* Handlers: configure the port, cors and serve the API */
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlewares.CheckDB(routes.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
