package manager

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var rm = map[string]string{}

func loadRoutes() {
	//The file descriptor will always be 4
	f := os.NewFile(4, "")
	f.Seek(0, 0)
	//read the file
	data, _ := ioutil.ReadAll(f)
	//load the route map into the global var
	json.Unmarshal(data, &rm)
}
