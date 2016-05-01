package manager

import (
	"net"
	"net/http"
	"os"
)

type connHandler struct{}

func init() {
	loadRoutes()
	makeChildListener()
	//3 will always be the fd of the listener we give to the subprocess
	f := os.NewFile(3, "")
	listener, _ := net.FileListener(f)
	http.Serve(listener, &connHandler{})
}
