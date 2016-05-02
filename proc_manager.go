package ggi

import (
	"os"
	"os/exec"
)

var idSeq = 0
var sockPath = "/var/run/"

//PManager is the primary process manager
type PManager struct {
	id      int
	process *os.Process
}

func (s *Server) spawnNewManager() {

	//make the process manager
	pm := &PManager{
		idSeq,
		nil,
	}
	s.processManagers[idSeq] = pm

	proc := exec.Command("./proc_man")
	proc.ExtraFiles = []*os.File{
		s.listenerFile,
		s.routeFile,
	}
	//set the stdout to this process's stdout
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stdout
	//start the command asynchronously
	proc.Start()

	//increment the sequence
	idSeq++

}
