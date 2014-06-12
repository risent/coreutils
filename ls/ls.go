// Origin: https://groups.google.com/d/topic/golang-nuts/aPWnfWoTVec/discussion

package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	Newline = "\n"
)

func ls(filename string) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if f == nil {
		fmt.Fprintf(os.Stderr, "ls: cannot access %s: %s\n", filename,
			err.Error())
		os.Exit(1)
	}
	files, err := f.Readdirnames(-1)
	if files == nil {
		fmt.Fprintf(os.Stderr, "ls: could not get files in %s: %s\n",
			filename, err.Error())
		os.Exit(1)
	}
	for j := 0; j < len(files); j++ {
		fmt.Printf("%s\n", files[j])
	}
	f.Close()
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		ls(".")
	} else {
		for i := 0; i < flag.NArg(); i++ {
			ls(flag.Arg(i))
		}
	}
}
