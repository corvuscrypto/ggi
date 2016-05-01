package manager

import (
	"os/exec"
	"path/filepath"
)

func compile(path string) {
	// get the last dir in the path
	dir, outfile := filepath.Split(path)
	if filepath.Ext(outfile) != ".go" {
		// normalize the path and filename as necessary
		_, filename := filepath.Split(dir)
		outfile = dir + "/" + filename
		//adjust the path to have a * suffix
		path += "/*"
	} else {
		// strip the extension
		outfile = dir + "/" + outfile[:len(outfile)-3]
	}

	cmd := exec.Command("go", "build", "-o", outfile, path)
	cmd.Run()

}
