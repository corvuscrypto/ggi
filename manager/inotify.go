package manager

import (
	"fmt"
	"os"
	"syscall"
)

type iNotify struct {
	file       *os.File
	watchdescs []int
}

func (p *process) newINotifyInstance() {
	in := new(iNotify)
	f, err := syscall.InotifyInit()
	if err != nil {
		fmt.Println("Manager: Unable to initialize an INotify instance; ", err)
	}
	in.file = os.NewFile(uintptr(f), "")

	go p.watchForINotifyEvents()
}

func (p *process) watchForINotifyEvents() {
	in := p.iNotifyInstance
	for {
		var b []byte
		//this will block until an event occurs
		in.file.Read(b)
		//dispatch the compile and replace the process
		//Code goes here
	}
}
