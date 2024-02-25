package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "modernc.org/sqlite"
)

func main() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static/"))

	conn, _ := sql.Open("sqlite", "database.sqlite3")
	db := DatabaseWrapper{conn}

	db.CreateTable()

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from router!")
	}).Methods("GET")

	api.HandleFunc("/event", func(w http.ResponseWriter, r *http.Request) {
		var event Event
		json.NewDecoder(r.Body).Decode(&event)
		db.CreateEvent(event)
	}).Methods("POST")

	api.HandleFunc("/event", func(w http.ResponseWriter, r *http.Request) {
		events := db.ReadEvents()
		json.NewEncoder(w).Encode(events)
	}).Methods("GET")

	api.HandleFunc("/event/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		event := db.ReadEventById(id)
		json.NewEncoder(w).Encode(event)
	}).Methods("GET")

	api.HandleFunc("/event/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		var event Event
		json.NewDecoder(r.Body).Decode(&event)
		db.UpdateEventById(id, event)
	}).Methods("PUT")

	api.HandleFunc("/event/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		db.DeleteEventById(id)
	}).Methods("DELETE")

	r.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	http.ListenAndServe(":8080", handlers.CORS(origins, methods)(r))
}
