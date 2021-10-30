package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"todo_app/db"
)

type Server struct {
	db *db.Queries
	ctx context.Context
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	todos, err := s.db.ListTodos(s.ctx)
	if err != nil {
		log.Println(err)
	}
	renderTemplate(w, "index.html", todos)
}

func (s *Server) createTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Failed to parse form on create")
	}
	text := sql.NullString{
		String: r.Form["text"][0],
		Valid: true,
	}
	err = s.db.CreateTodo(s.ctx, text)
	if err != nil {
		log.Printf("Error while creating todo: %e\n", err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *Server) deleteTodo(w http.ResponseWriter, r *http.Request) {
	queryId, err := strconv.Atoi(r.URL.Query()["id"][0])
	if err != nil {
		log.Printf("Error while getting id: %e", err)
		log.Println(fmt.Fprintf(w, "Error, invalid id"))
	}
	id := sql.NullInt32{
		Int32: int32(queryId),
		Valid: true,
	}
	err = s.db.DeleteTodo(s.ctx, id)
	if err != nil {
		log.Printf("Error while deleting todo: %e", err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
