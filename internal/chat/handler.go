package chat

import (
	"html/template"
	"log"
	"net/http"
)

// Pode ser apenas um map por enquanto
// type chatServer struct {
var subscribers map[*subscriber]struct{}

// Adcionar tempo limite? (1 dia)
// }

type subscriber struct {
	ip    string
	color int
}

// IndexHandler aplica o numero de users conectados no template principal.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	numUsers := getNumberOfUsers()

	tmpl := template.Must(template.ParseFiles("web/index.html"))
	err := tmpl.Execute(w, numUsers)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Fatal("Error executing template:", err)
	}
}

func getNumberOfUsers() int {
	return 2
}
