package ggi

import (
	"encoding/gob"
	"log"
	"net/http"
)

type requestPack struct {
	req *http.Request
}

type responsePack struct {
	Pid  int
	Code int
	Res  []byte
}

func (p *PManager) listen() {
	for {
		conn, _ := p.socket.AcceptUnix()
		enc := gob.NewEncoder(conn)
		dec := gob.NewDecoder(conn)
		var res = &responsePack{}
		//wait for the acknowledgement response
		dec.Decode(&res)
		if &res == nil {
			continue
		}
		log.Println(res)
		if res.Pid == 0 {
			continue
		}
		for _, proc := range p.processes {
			if proc.process.Pid == res.Pid {
				proc.connection = conn
				proc.encoder = enc
				proc.decoder = dec
				break
			}
		}
	}
}

//intermediary function
func (p *PManager) send(r *http.Request) *http.Response {
	return nil
}
