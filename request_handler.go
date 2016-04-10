package ggi

import (
	"log"
	"net/http"
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
	var pm *PManager
	if len(processManagers) == 0 {
		pm = spawnNewManager()
	} else {
		pm = processManagers[0]
	}
	proc, ok := pm.processes[path]
	if !ok {
		proc = pm.spawnProcess(path, filePath)
		if proc == nil {
			log.Println("an Error occurred spawning a process", path, filePath)
			w.WriteHeader(500)
			return
		}
	}
	log.Println("handling request")
	proc.handleRequest(w, r)
	log.Println("done")
}
