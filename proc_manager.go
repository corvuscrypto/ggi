package ggi

import (
	"encoding/gob"
	"log"
	"net"
	"os"
	"os/exec"
	"sync"
)

var idSeq = 0
var processManagers = []*PManager{}
var sockPath = "/var/run/"

//PManager is the primary process manager
type PManager struct {
	id        int
	locks     map[string]*sync.Mutex
	processes map[string]*Process
	socket    *net.UnixListener
}

//spawn a process and setup the connection
func (p *PManager) spawnProcess(routeName, path string) *Process {
	log.Println("Spawning process: ")
	var oPath = path
	if path[len(path)-3:] == ".go" {
		path = path[:len(path)-3] + ".ggp"
	} else {
		if path[len(path)-1] == '/' {
			path = path[:len(path)-1] + "/proc.ggp"
		}
	}
	process, err := os.StartProcess(path, nil, &os.ProcAttr{})
	if err != nil {
		//perhaps it's just not compiled yet?
		//attempt to compile the dir/file
		log.Println("attempting compile")
		err = exec.Command("go", "build", "-o", path, oPath).Run()
		if err != nil {
			log.Println(err)
			return nil
		}
		log.Println("Attempting to start process")
		process, err = os.StartProcess(path, nil, &os.ProcAttr{})
		if err != nil {
			log.Println(err)
			return nil
		}
	}

	ggiProc := &Process{}
	ggiProc.managerID = p.id
	ggiProc.process = process
	ggiProc.connection, err = p.socket.AcceptUnix()
	if err != nil {
		log.Println(err)
	}
	ggiProc.encoder = gob.NewEncoder(ggiProc.connection)
	ggiProc.decoder = gob.NewDecoder(ggiProc.connection)
	p.processes[routeName] = ggiProc
	return ggiProc
}

func (s *Server) spawnNewManager() {
	IPCAddress, err := net.ResolveUnixAddr("unix", sockPath+"ggi.sock")
	if err != nil {
		log.Println(err)
	}
	IPCListener, err := net.ListenUnix("unix", IPCAddress)
	if err != nil {
		//likely reached due to an already-existing socket. attempt to remove and
		//create
		os.Remove(sockPath + "ggi.sock")
		IPCListener, err = net.ListenUnix("unix", IPCAddress)
		if err != nil {
			//halt the process and return nothing
			log.Println(err)
			return
		}
		log.Println(err)
	}

	//make the process manager
	pm := &PManager{
		idSeq,
		map[string]*sync.Mutex{},
		map[string]*Process{},
		IPCListener,
	}

	//increment the sequence
	idSeq++

	s.processManagers = append(processManagers, pm)
	// go pm.listen()

}
