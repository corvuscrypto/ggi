package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"strings"
	"unsafe"
)

var processes = map[string]*process{}

func handleRequest(pipe *bufio.ReadWriter) {
	data, _ := ioutil.ReadAll(pipe.Reader)

	//get the route from the http header
	var path = parseHTTPHeadline((uintptr)(unsafe.Pointer(&data[0])), len(data))
	p, ok := processes[path]
	if !ok {
		proc, err := spawnChildProcess(path, rm[path])
		p = proc
		if err != nil {
			log.Fatal(err)
		}
	}
	pipe.Writer.Write(p.handle(data))

}

func parseHTTPHeadline(char uintptr, length int) string {
	path := ""

	//use pointers and save some time here
	for i := 0; i < length; i++ {
		var b = *(*byte)(unsafe.Pointer(char + uintptr(i)))
		if b == '\n' {
			break
		}
	}

	header := strings.Split(path, " ")
	if len(header) > 1 {
		return header[1]
	}
	return ""
}
