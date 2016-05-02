package manager

import (
	"encoding/gob"
	"net"
	"os"

	"github.com/corvuscrypto/ggi/transport"
)

type process struct {
	proc *os.Process
	srcPath,
	execPath string
	pipe            *net.TCPConn
	decoder         *gob.Decoder
	iNotifyInstance *iNotify
	srcFilepaths    []string
}

func (p *process) handle(data []byte) *transport.Response {
	p.pipe.Write(data)
	var res = &transport.Response{}
	p.decoder.Decode(res)
	return res
}
