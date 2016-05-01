package manager

import (
	"encoding/gob"
	"net"
	"os"

	"github.com/corvuscrypto/ggi/transport"
)

type process struct {
	proc *os.Process
	pipe *net.TCPConn
}

func (p *process) handle(data []byte) []byte {
	p.pipe.Write(data)
	var res = &transport.Response{}
	dec := gob.NewDecoder(p.pipe)
	dec.Decode(res)
	return res.Data
}
