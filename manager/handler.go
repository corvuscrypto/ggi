package manager

import (
	"log"
	"net/http"
	"net/http/httputil"
	"sync"
)

var processes = map[string]*process{}
var connectionLocker = &sync.Mutex{}

func (c *connHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//get the route from the http header
	path := r.URL.Path

	//if we don't track this route, 404 it
	if _, ok := rm[path]; !ok {
		w.WriteHeader(404)
		return
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
	w.Write(res)
}
