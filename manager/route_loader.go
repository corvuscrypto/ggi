package manager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var rm = map[string]string{}
var priorityMatcher = []string{}

func loadRoutes() {
	//The file descriptor will always be 4
	f := os.NewFile(4, "")
	f.Seek(0, 0)
	//read the file
	//load the route map into the global var
	data, _ := ioutil.ReadAll(f)
	err := json.Unmarshal(data, &rm)
	//cheap handling of rare race condition
	for err != nil {
		f.Seek(0, 0)
		data, _ = ioutil.ReadAll(f)
		err = json.Unmarshal(data, &rm)
	}
	generatePriorityMatcher()
}

// The idea here is that we don't need a robust multiplexer or matching system.
// We can instead assume that sorting by the number of subspaces in a path (highest first)
// will allow us to iterate through the array will find the first best match for the path
func generatePriorityMatcher() {

	for k := range rm {
		//for each route, we want to split by the / separator.
		//then we just use insertion sort to sort into priorityMatcher for matching
		splitPath := strings.Split(k, "/")
		fmt.Println(k)
		//toss any trailing empty strings
		if splitPath[len(splitPath)-1] == "" {
			splitPath = splitPath[:len(splitPath)-1]
		}
		if len(priorityMatcher) == 0 {
			priorityMatcher = []string{k}
		} else {
			var inserted bool
			for i, v := range priorityMatcher {
				if len(splitPath) >= len(v) {
					//insert into slice
					temp := make([]string, len(priorityMatcher[i:]))
					copy(temp, priorityMatcher[i:])
					priorityMatcher = append(append(priorityMatcher[:i], k), temp...)
					inserted = true
					break
				}
			}
			if !inserted {
				priorityMatcher = append(priorityMatcher, k)
			}
		}
	}
}
