package manager

import (
	"net"
	"os"
)

type process struct {
	proc *os.Process
	pipe *net.TCPConn
}

func (p *process) handle(data []byte) []byte {
	p.pipe.Write(data)
	return nil
}
