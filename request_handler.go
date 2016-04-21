package ggi

import (
	"log"
	"net/http"
	"sync"
)

//HandleRequest handles an incoming request and attempts to resolve it by
//delegating it to a managed process for handling
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	//if we don't even have a file/dir to reference just send 404
	filePath, ok := routes[path]
	if !ok {
		w.WriteHeader(404)
		return
	}
	var pm = processManagers[0]

	//get the waitGroup
	lock, ok := pm.locks[path]
	if !ok {
		lock = &sync.Mutex{}
		pm.locks[path] = lock
	}
	lock.Lock()
	proc, ok := pm.processes[path]
	if !ok {
		proc = pm.spawnProcess(path, filePath)
		if proc == nil {
			lock.Unlock()
			log.Println("an Error occurred spawning a process to serve ", path)
			w.WriteHeader(500)
			return
		}
	}
	lock.Unlock()
	go proc.handleRequest(w, r)
}
