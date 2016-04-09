package ggi

var routes map[string]string

//RegisterRoute registers a path from a request to a particular file
func RegisterRoute(path, filepath string) {
	routes[path] = filepath
}
