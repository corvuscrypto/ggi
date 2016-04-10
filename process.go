package ggi

import (
	"encoding/gob"
	"net"
	"net/http"
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

func (p *Process) handleRequest(w http.ResponseWriter, r *http.Request) {
	p.encoder.Encode(r)
	var res responsePack
	p.decoder.Decode(&res)
	if &res == nil {
		w.WriteHeader(http.StatusBadGateway)
	}
	w.Write(res.res)
}
