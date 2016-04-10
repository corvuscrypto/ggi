package ggi

import "net/http"

//HandleRequest handles an incoming request and attempts to resolve it by
//delegating it to a managed process for handling
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	//if we don't even have a file/dir to reference just send 404
	_, ok := routes[path]
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
	proc := pm.processes[path]
	proc.handleRequest(w, r)

}
