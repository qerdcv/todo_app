package server

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	t, err := template.ParseFiles(fmt.Sprintf("templates/%s", templateName))
	if err != nil {
		fmt.Println("Error while serving index request")
	}
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println("Failed to execute template")
	}
}

