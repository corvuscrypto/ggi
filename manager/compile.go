package manager

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func compile(path string) {
	commandArgs := []string{"build", "-o"}
	// get the last dir in the path
	dir, outfile := filepath.Split(path)
	if filepath.Ext(outfile) != ".go" {
		// normalize the path and filename as necessary
		var filename string
		if outfile == "" {
			_, filename = filepath.Split(dir[:len(dir)-1])
		} else {
			_, filename = filepath.Split(dir)
		}
		outfile = dir + "/" + filename
		commandArgs = append(commandArgs, outfile)
		path += ""
		//default to get all the .go files in the dir
		files, _ := ioutil.ReadDir(dir)
		for _, file := range files {
			name := file.Name()
			if filepath.Ext(name) == ".go" {
				commandArgs = append(commandArgs, dir+"/"+name)
			}
		}
	} else {
		// strip the extension
		outfile = dir + "/" + outfile[:len(outfile)-3]
		commandArgs = append(commandArgs, outfile, path)
	}

	cmd := exec.Command("go", commandArgs...)
	cmd.Stderr = os.Stdout
	cmd.Run()

}
