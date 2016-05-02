package child

import "net/http"

//Handle is a convenience function for attaching a custom handler to
//the DefaultServeMux from the http package
func Handle(path string, h http.Handler) {
	http.DefaultServeMux.Handle(path, h)
}
