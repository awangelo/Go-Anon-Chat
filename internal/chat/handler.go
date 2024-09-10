package chat

import (
	"net/http"
)

// IndexHandler aplica o numero de users conectados no template principal.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/index.html")
}
