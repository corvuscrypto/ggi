package ggi

import (
	"encoding/gob"
	"net/http"
)

type responsePack struct {
	pid int
	res *http.ResponseWriter
}

func (p *PManager) listenAndServe() {
	for {
		conn, _ := p.socket.AcceptUnix()
		enc := gob.NewEncoder(conn)
		dec := gob.NewDecoder(conn)
		var res responsePack
		//wait for the acknowledgement response
		dec.Decode(&res)
		if &res == nil {
			continue
		}
		if res.pid == 0 {
			continue
		}

		proc := p.processes[res.pid]
		proc.connection = conn
		proc.encoder = enc
		proc.decoder = dec

	}
}

//intermediary function
func (p *PManager) send(r *http.Request) *http.Response {
	return nil
}
