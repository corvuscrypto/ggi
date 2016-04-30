package manager

import (
	"log"
	"net"
	"os"
	"sync"
)

var connectionLocker = &sync.Mutex{}

var childListener *net.TCPListener

func makeChildListener() {
	// resolve free tcp address
	addr, err := net.ResolveTCPAddr("tcp", "0")
	childListener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
}

func spawnChildProcess(route string, proc string) (*process, error) {
	//first lock the function to prevent other connections
	//from accidentally associating with the new process
	connectionLocker.Lock()
	p, err := os.StartProcess(proc, []string{childListener.Addr().String()}, &os.ProcAttr{})
	if err != nil {
		log.Fatal(err)
	}
	conn, err := childListener.AcceptTCP()
	if err != nil {
		return nil, err
	}
	process := &process{
		p,
		conn,
	}
	processes[route] = process
	return process, nil
}
