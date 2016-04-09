package child

import (
	"encoding/gob"
	"log"
	"net"
	"net/http"
	"os"
)

var pid = os.Getpid()

type responsePack struct {
	pid int
	res *http.ResponseWriter
}

var connection struct {
	conn    *net.UnixConn
	encoder *gob.Encoder
	decoder *gob.Decoder
}

func connHandler() {
	for {
		var r http.Request
		connection.decoder.Decode(&r)
		if &r != nil {
			connection.encoder.Encode(handleRequest(&r))
		} else {
			connection.encoder.Encode(&responsePack{
				pid,
				nil,
			})
		}
	}
}

func init() {
	addr, err := net.ResolveUnixAddr("unix", "/var/run/ggi.sock")
	if err != nil {
		//if there's an error, kill start the process
		log.Fatal(err)
	}
	connection.conn, err = net.DialUnix("unix", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	connection.encoder = gob.NewEncoder(connection.conn)
	connection.decoder = gob.NewDecoder(connection.conn)
	go connHandler()
}
