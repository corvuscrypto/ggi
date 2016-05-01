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

func makeChildListener() {
	// resolve free tcp address
	addr, err := net.ResolveTCPAddr("tcp", "0")
	childListener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
}

func spawnChildProcess(route string, proc string) (*process, error) {
	//normalize proc
	path := getExecPath(proc)
	cmd := exec.Command(path, childListener.Addr().String())
	cmd.Stderr = os.Stdout
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		compile(proc)
		err = cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
	}
	conn, err := childListener.AcceptTCP()
	if err != nil {
		return nil, err
	}
	process := &process{
		cmd.Process,
		conn,
		gob.NewDecoder(conn),
	}
	processes[route] = process
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
