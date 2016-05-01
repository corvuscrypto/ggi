package child

import "net/http"

var handlers = map[string]func(http.ResponseWriter, *http.Request) []byte{}

//AddHandler associates a function with your request object
func AddHandler(route string, f func(http.ResponseWriter, *http.Request) []byte) {
	handlers[route] = f
}
