package child

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
)

var pid = os.Getpid()

var connection struct {
	conn   *net.TCPConn
	reader *bufio.Reader
}

func writeErr(i interface{}) {
	f, _ := os.Create("./test_output.txt")
	defer f.Close()
	f.WriteString(fmt.Sprintf("%v", i))
}

func connHandler() {
	for {
		//make the new request
		req, err := http.ReadRequest(connection.reader)
		if err != nil {
			connection.conn.Write([]byte{})
		}
		connection.conn.Write(handleRequest(req))
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
	go connHandler()
}
