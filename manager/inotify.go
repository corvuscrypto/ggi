package manager

import (
	"fmt"
	"syscall"
)

const iNotifyEvents = syscall.IN_CLOSE_WRITE

type iNotify struct {
	fd         int
	watchdescs []int
}

func (p *process) newINotifyInstance() {
	in := new(iNotify)
	f, err := syscall.InotifyInit()
	if err != nil {
		fmt.Println("Manager: Unable to initialize an INotify instance; ", err)
	}
	in.fd = f
	in.addWatchers(p)
	p.iNotifyInstance = in
	go p.watchForINotifyEvents()
}

func (i *iNotify) addWatchers(p *process) {
	i.watchdescs = []int{}
	for _, f := range p.srcFilepaths {
		wd, err := syscall.InotifyAddWatch(i.fd, f, iNotifyEvents)
		if err != nil {
			fmt.Println(err)
			continue
		}
		i.watchdescs = append(i.watchdescs, wd)
	}
}

func (p *process) watchForINotifyEvents() {

	for {
		var b = make([]byte, 2<<16)
		//this will block until an event occurs
		syscall.Read(p.iNotifyInstance.fd, b)

		//dispatch the compile and replace the process
		p.compile()
		p.reloadProcess()
	}
}
