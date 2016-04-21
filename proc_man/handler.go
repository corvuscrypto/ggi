package main

import (
	"bufio"
	"net/http"
	"os"
)

var processes = map[string]*process{}

func handleRequest(pipe *bufio.ReadWriter) {
	req, err := http.ReadRequest(pipe.Reader)
	if err != nil {
		os.Exit(1)
	}
	path := req.URL.Path
	pipe.Writer.Write(processes[path].handle(req))
}
