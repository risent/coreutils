// Origin: https://groups.google.com/d/topic/golang-nuts/aPWnfWoTVec/discussion

package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

const (
	Newline = "\n"
)

var longFormat bool

func init() {
	flag.BoolVar(&longFormat, "l", false, "Long Format")
}

func ls(filename string, longFormat bool) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if f == nil {
		fmt.Fprintf(os.Stderr, "ls: cannot access %s: %s\n", filename,
			err.Error())
		os.Exit(1)
	}
	files, err := f.Readdir(-1)
	if files == nil {
		fmt.Fprintf(os.Stderr, "ls: could not get files in %s: %s\n",
			filename, err.Error())
		os.Exit(1)
	}
	for j := 0; j < len(files); j++ {
		if longFormat {
			timeLayout := "Jan 2 15:04"
			gid := files[j].Sys().(*syscall.Stat_t).Gid
			uid := files[j].Sys().(*syscall.Stat_t).Uid
			group, _ := user.LookupId(strconv.Itoa(int(gid)))
			user, _ := user.LookupId(strconv.Itoa(int(uid)))

			fmt.Printf("%s\t%s\t%s\t%d\t%s\t%s\n", files[j].Mode(), group.Username, user.Username, files[j].Size(), files[j].ModTime().Format(timeLayout), files[j].Name())
		} else {

			fmt.Printf("%s\n", files[j].Name())
		}
	}
	f.Close()
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		ls(".", longFormat)
	} else {
		for i := 0; i < flag.NArg(); i++ {
			ls(flag.Arg(i), longFormat)
		}
	}
}
