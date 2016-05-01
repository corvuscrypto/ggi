package ggi

import (
	"log"
	"net"
	"os"
	"strconv"
)

//Server contains configuration information for the ggi server.
//Some fields are ignored depending on settings.
//E.g. ProcessTimeout is only relevant if DynamicManagement is true.
type Server struct {
	ServerPort              int
	MaxProcessManagers      int
	StartingProcessManagers int
	DynamicManagement       bool
	ProcessTimeout          int

	listenerFile    *os.File
	routes          map[string]string
	routeFile       *os.File
	processManagers map[int]*PManager
}

func (s *Server) setup() {
	//get the tcp address
	laddr, err := net.ResolveTCPAddr("tcp", ":"+strconv.Itoa(s.ServerPort))
	if err != nil {
		log.Fatal("Unable to resolve an address from the given port ", s.ServerPort, "; ", err)
	}
	l, err := net.ListenTCP("tcp", laddr)
	if l == nil {
		log.Fatal("Unable to use port ", s.ServerPort, "; ", err)
	}

	//set the listenerFile field to the listener's file pointer
	s.listenerFile, _ = l.File()

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
	s.writeRoutes()
	for i := 0; i < s.StartingProcessManagers; i++ {
		s.spawnNewManager()
	}
	select {}
}

//default settings
var defaultServer = &Server{
	8080,
	10,
	2,
	false,
	0,
	nil,
	map[string]string{},
	nil,
	map[int]*PManager{},
}
