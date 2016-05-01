package child

import "net/http"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if v, ok := handlers[r.URL.Path]; ok {
		v(w, r)
	}
}
