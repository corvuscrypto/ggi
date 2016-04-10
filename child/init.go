package child

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
)

var pid = os.Getpid()

type Request struct {
	Params []string
	Path   string
}

type responsePack struct {
	Pid  int
	Code int
	Res  []byte
}

var connection struct {
	conn    *net.UnixConn
	encoder *gob.Encoder
	decoder *gob.Decoder
}

func writeErr(i interface{}) {
	f, _ := os.Create("./test_output.txt")
	defer f.Close()
	f.WriteString(fmt.Sprintf("%v", i))
}

func connHandler() {
	for {
		var r = &Request{}
		err := connection.decoder.Decode(r)
		if err != nil {
			writeErr(err)
		}
		if r == nil {
			continue
		}
		writeErr("received and handled:")
		writeErr(r)
		connection.encoder.Encode(handleRequest(r))
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
}

func Start() {
	connHandler()
}
