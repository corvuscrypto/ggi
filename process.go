package ggi

import (
	"encoding/gob"
	"net"
	"os"
)

//Process describes a process to be used for request handling
type Process struct {
	process    *os.Process
	connection *net.UnixConn
	encoder    *gob.Encoder
	decoder    *gob.Decoder
}

func (p *Process) kill() {
	p.process.Kill()
}
