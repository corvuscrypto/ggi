package ggi

import (
	"log"
	"net"
	"os"
	"strconv"
)

//Server contains configuration information for the ggi server.
//Some fields are ignored depending on settings.
//E.g. TCPPort is only relevant if UseTCP is true, and ProcessTimeout is only
//relevant if DynamicManagement is true.
type Server struct {
	UseTCP                  bool
	TCPPort                 int
	SockPath                string
	MaxProcessManagers      int
	StartingProcessManagers int
	DynamicManagement       bool
	ProcessTimeout          int

	processManagers []*PManager
}

func (s *Server) setup() {
	//are we using TCP?
	if s.UseTCP {
		//check that port is open
		l, err := net.Listen("tcp", ":"+strconv.Itoa(s.TCPPort))
		if l == nil {
			log.Fatal("Unable to use port ", s.TCPPort, "; ", err)
		}
		l.Close()
	} else {
		//check to see we can make the socket
		addr, err := net.ResolveUnixAddr("unix", s.SockPath)
		if err != nil {
			log.Fatal(err)
		}
		l, err := net.ListenUnix("unix", addr)
		if err != nil {
			//likely reached due to an already-existing socket. attempt to remove and
			//create
			os.Remove(s.SockPath)
			l, err = net.ListenUnix("unix", addr)
			if err != nil {
				//halt the process and return nothing
				log.Fatal("Unable to use socket at ", s.SockPath, "; ", err)
				return
			}
			l.Close()
		}
		l.Close()
	}
	//check that StartingProcessManagers is not > MaxProcessManagers
	if s.StartingProcessManagers > s.MaxProcessManagers {
		log.Fatal("Starting process managers cannot be more than the maximum!")
	}

	if s.DynamicManagement {
		//check to see if processTimeout is > 0
		if s.ProcessTimeout <= 0 {
			log.Fatal("ProcessTimeout cannot be negative or 0")
		}
	}

}

//Serve using the associated Server's configuration
func (s *Server) Serve() {
	s.setup()
	s.spawnNewManager()
}

//default settings
var defaultServer = &Server{
	true,
	17001,
	"",
	10,
	2,
	false,
	0,
	[]*PManager{},
}
