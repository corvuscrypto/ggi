package ggi

import (
	"encoding/gob"
	"net"
	"net/http"
	"os"
)

//Process describes a process to be used for request handling
type Process struct {
	managerID  int
	process    *os.Process
	connection *net.UnixConn
	encoder    *gob.Encoder
	decoder    *gob.Decoder
}

func (p *Process) kill() {
	//attempt to kill just in case
	p.process.Kill()
	//consume Zombie process
	p.process.Wait()
	for _, pm := range processManagers {
		if pm.id == p.managerID {
			for r, proc := range pm.processes {
				if proc == p {
					delete(pm.processes, r)
					p = nil
					return
				}
			}
		}
	}
}

func (p *Process) handleRequest(w http.ResponseWriter, r *http.Request) {
	req := &request{
		nil,
		r.URL.Path,
	}
	err := p.encoder.Encode(req)
	if err != nil {
		p.kill()
		w.WriteHeader(500)
		return
	}
	var res responsePack
	err = p.decoder.Decode(&res)
	if err != nil {
		p.kill()
		w.WriteHeader(500)
		return
	}
	if &res == nil {
		w.WriteHeader(http.StatusBadGateway)
	}
	w.Write(res.Res)
}
