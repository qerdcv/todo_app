package server

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"todo_app/db"
)

func setupRouter(s *Server) {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// TODO: Use gorilla/mux
	http.HandleFunc(
		"/create",
		filterMethod([]string{http.MethodPost}, s.createTodo))
	http.HandleFunc(
		"/delete",
		filterMethod([]string{http.MethodGet}, s.deleteTodo))
	http.HandleFunc("/",
		filterMethod([]string{http.MethodGet}, s.index))
}

func New() (*Server, error) {
	ctx := context.Background()
	d, err := sql.Open(
		"postgres",
		"user=postgres password=postgres dbname=todo_app host=db sslmode=disable")
	// TODO: Move to the env variables
	if err != nil {
		log.Printf("Error while connecting to db %e", err)
		return nil, err
	}
	queries, err := db.Prepare(ctx, d)
	if err != nil {
		log.Printf("Failed to prepare db %e", err)
		return nil, err
	}
	return &Server{
		db: queries,
		ctx: ctx,
	}, nil
}

func (s *Server) Run() {
	setupRouter(s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
