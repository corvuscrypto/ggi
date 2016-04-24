package main

import (
	"bufio"
	"io/ioutil"
	"strings"
	"unsafe"
)

var processes = map[string]*process{}

func handleRequest(pipe *bufio.ReadWriter) {
	data, _ := ioutil.ReadAll(pipe.Reader)

	//get the route from the http header
	var path = parseHTTPHeadline((uintptr)(unsafe.Pointer(&data[0])), len(data))

	pipe.Writer.Write(processes[path].handle(data))
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
