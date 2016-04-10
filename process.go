package ggi

import (
	"encoding/gob"
	"log"
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
	req := &request{
		nil,
		r.URL.Path,
	}
	err := p.encoder.Encode(req)
	if err != nil {
		log.Println(err)
	}
	var res responsePack
	err = p.decoder.Decode(&res)
	if err != nil {
		log.Println(err)
	}
	if &res == nil {
		w.WriteHeader(http.StatusBadGateway)
	}
	w.Write(res.Res)
}
