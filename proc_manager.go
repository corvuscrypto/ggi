package ggi

import (
	"encoding/gob"
	"log"
	"net"
	"os"
	"os/exec"
)

var processManagers = []*PManager{}

//PManager is the primary process manager
type PManager struct {
	processes map[string]*Process
	socket    *net.UnixListener
}

//spawn a process and setup the connection
func (p *PManager) spawnProcess(routeName, path string) *Process {
	var oPath = path
	if path[len(path)-3:] == ".go" {
		path = path[len(path)-3:] + ".a"
	} else {
		if path[len(path)-1] == '/' {
			path = path[:len(path)-1] + "/proc.a"
		}
	}
	process, err := os.StartProcess(path, nil, &os.ProcAttr{})
	if err != nil {
		//perhaps it's just not compiled yet?
		//attempt to compile the dir/file
		err = exec.Command("go", "build", "-o", path, oPath).Run()
		if err != nil {
			log.Println(err)
			return nil
		}
		process, err = os.StartProcess(path, nil, &os.ProcAttr{})
		if err != nil {
			log.Println(err)
			return nil
		}
	}

	ggiProc := &Process{}
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

func spawnNewManager() *PManager {
	IPCAddress, err := net.ResolveUnixAddr("unixpacket", "/var/run/ggi.sock")
	if err != nil {
		log.Println(err)
	}
	IPCListener, err := net.ListenUnix("unixpacket", IPCAddress)
	if err != nil {
		//likely reached due to an already-existing socket. attempt to remove and
		//create
		os.Remove("/var/run/ggi.sock")
		IPCListener, err = net.ListenUnix("unixpacket", IPCAddress)
		if err != nil {
			//halt the process and return nothing
			log.Println(err)
			return nil
		}
		log.Println(err)
	}

	pm := &PManager{
		map[string]*Process{},
		IPCListener,
	}
	processManagers = append(processManagers, pm)

	return pm
}
