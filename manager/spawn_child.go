package manager

import (
	"encoding/gob"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
)

var childListener *net.TCPListener

//set this to true for now
var watchChanges = true

func makeChildListener() {
	// resolve free tcp address
	addr, err := net.ResolveTCPAddr("tcp", "0")
	childListener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
}

func spawnChildProcess(route string, proc string) (*process, error) {
	//route in this case is the src code route
	//normalize proc
	path := getExecPath(proc)

	process := new(process)
	//store the generating path
	process.srcPath = proc
	process.execPath = path
	process.loadProcess()
	processes[route] = process
	//add the notify instance
	process.newINotifyInstance()
	return process, nil
}

func getExecPath(path string) string {
	// get the last dir in the path
	dir, outfile := filepath.Split(path)
	if filepath.Ext(outfile) != ".go" {
		// normalize the path and filename as necessary
		var filename string
		if outfile == "" {
			_, filename = filepath.Split(dir[:len(dir)-1])
		} else {
			_, filename = filepath.Split(dir)
		}
		outfile = dir + "/" + filename
	} else {
		// strip the extension
		outfile = dir + "/" + outfile[:len(outfile)-3]
	}
	return outfile
}

func (p *process) loadProcess() {
	cmd := exec.Command(p.execPath, childListener.Addr().String())
	cmd.Stderr = os.Stdout
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		p.compile()
		err = cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
	}
	conn, err := childListener.AcceptTCP()
	if err != nil {
		return
	}

	p.proc = cmd.Process
	p.pipe = conn
	p.decoder = gob.NewDecoder(conn)

}

func (p *process) reloadProcess() {
	//kill-wait to release the current process
	p.proc.Kill()
	p.proc.Wait()
	//reload the process into the struct
	p.loadProcess()
}
