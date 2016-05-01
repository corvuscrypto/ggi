package manager

import (
	"bufio"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"unsafe"
)

var processes = map[string]*process{}
var connectionLocker = &sync.Mutex{}

func handleRequest(pipe *bufio.ReadWriter) {
	data, _ := ioutil.ReadAll(pipe.Reader)

	//get the route from the http header
	var path = parseHTTPHeadline((uintptr)(unsafe.Pointer(&data[0])), len(data))

	//first lock the function to prevent other connections
	//from accidentally associating with the new process

	connectionLocker.Lock()
	p, ok := processes[path]
	if !ok {
		proc, err := spawnChildProcess(path, rm[path])
		p = proc
		if err != nil {
			log.Fatal(err)
		}
	}
	connectionLocker.Unlock()

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
		path += string(b)
	}

	header := strings.Split(path, " ")
	if len(header) > 1 {
		return header[1]
	}
	return ""
}
