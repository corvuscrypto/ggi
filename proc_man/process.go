package main

import (
	"net"
	"net/http"
	"os"
)

type process struct {
	proc *os.Process
	pipe *net.TCPListener
}

func (p *process) handle(r *http.Request) []byte {
	return nil
}
