package manager

import (
	"encoding/gob"
	"os"
)

var rm = map[string]string{}

func loadRoutes() {
	//The file descriptor will always be 4
	f := os.NewFile(4, "")
	f.Seek(0, 0)
	//load the route map into the global var
	gob.NewDecoder(f).Decode(&rm)
}
