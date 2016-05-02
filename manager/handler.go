package manager

import (
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"sync"
)

var processes = map[string]*process{}
var connectionLocker = &sync.Mutex{}

func (c *connHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//get the route from the http header
	path := r.URL.Path

	if _, ok := rm[path]; !ok {
		//if we don't track this route, attempt to find the best match
		path = findRoute(path)
		// if we still haven't found a match, 404 it yo
		if path == "" {
			w.WriteHeader(404)
			return
		}
	}
	//first lock the function to prevent other connections
	//from accidentally associating with the new process

	connectionLocker.Lock()
	p, ok := processes[path]
	if !ok {
		proc, err := spawnChildProcess(path, rm[path])
		p = proc
		if err != nil {
			log.Fatal(err)
		}
	}
	connectionLocker.Unlock()

	data, _ := httputil.DumpRequest(r, true)

	res := p.handle(data)
	for k, v := range res.Headers {
		for _, sv := range v {
			w.Header().Add(k, sv)
		}
	}
	w.WriteHeader(res.StatusCode)
	w.Write(res.Data)
}

func findRoute(path string) string {
	var result string
	for _, v := range priorityMatcher {
		if strings.Index(path, v) == 0 {
			return v
		}
	}
	return result
}
