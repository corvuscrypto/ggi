package main

import (
	"bufio"
	"net"
	"os"
)

func main() {
	//3 will always be the fd of the listener we give to the subprocess
	f := os.NewFile(3, "")
	listener, _ := net.FileListener(f)
	for {
		conn, _ := listener.Accept()
		bufio.NewReader(conn)
		pipe := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
		go handleRequest(pipe)
	}
}
