package child

import "net/http"

func handleRequest(r *http.Request) []byte {

	if v, ok := handlers[r.URL.Path]; ok {
		return v(r)
	}
	return []byte{}
}
