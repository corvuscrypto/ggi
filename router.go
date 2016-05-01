package ggi

import (
	"encoding/json"
	"os"
	"strconv"
)

//RegisterRoute registers a path from a request to a particular executable file
func (s *Server) RegisterRoute(path, filepath string) {
	s.routes[path] = filepath
}

//write the routes to a file in tmp and store the file pointer
func (s *Server) writeRoutes() {
	//create the empty file in /tmp
	f, _ := os.Create("/tmp/ggi_route" + strconv.Itoa(s.ServerPort))
	s.routeFile = f
	data, _ := json.Marshal(s.routes)
	f.Write(data)
	//leave it open and let the process manager close it when done
}
