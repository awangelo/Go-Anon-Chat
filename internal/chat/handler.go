package chat

import (
	"html/template"
	"log"
	"net/http"
)

// IndexHandler aplica o numero de users conectados no template principal.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// ParseFiles retorna o template do chat com o numero de subscribers.
	tmpl := template.Must(template.ParseFiles("web/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Fatal("Error executing template:", err)
	}
}
