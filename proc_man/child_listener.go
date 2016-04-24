package main

import (
	"log"
	"net"
	"os"
)

func makeChildListener() *net.TCPListener {
	// resolve free tcp address
	addr, err := net.ResolveTCPAddr("tcp", "0")
	childListener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	return childListener
}

func spawnChildProcess() *os.Process {

}
