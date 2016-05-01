package child

import (
	"bufio"
	"encoding/gob"
	"net"
	"net/http"
	"os"

	"github.com/corvuscrypto/ggi/transport"
)

var pid = os.Getpid()

var connection struct {
	conn    *net.TCPConn
	reader  *bufio.Reader
	encoder *gob.Encoder
}

func connHandler() {
	for {
		var res = new(transport.Response)
		//make the new request
		req, err := http.ReadRequest(connection.reader)
		if err == nil {
			handleRequest(res, req)
		}
		connection.encoder.Encode(res)
	}
}

func init() {
	addrStr := os.Args[1]
	addr, err := net.ResolveTCPAddr("tcp", addrStr)
	if err != nil {
		//if there's an error, kill start the process
		os.Exit(1)
	}
	connection.conn, err = net.DialTCP("tcp", nil, addr)
	if err != nil {
		os.Exit(1)
	}
	//make a bufio reader from the connection and store it in connection
	connection.reader = bufio.NewReader(connection.conn)
	//make gob encoder
	connection.encoder = gob.NewEncoder(connection.conn)
	go connHandler()
}
