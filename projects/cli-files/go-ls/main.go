package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		d, err := os.ReadDir(".")
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		fnames := ls(d)
		for _, f := range fnames {
			fmt.Println(f)
		}
	} else {
		dirs := os.Args[1:]
		for _, dname := range dirs {
			fmt.Printf("%s:\n", dname)
			d, err := os.ReadDir(dname)
			if err != nil {
				fmt.Fprint(os.Stderr, err)
			}
			fnames := ls(d)
			for _, f := range fnames {
				fmt.Println(f)
			}
			fmt.Println()
		}

	}

}

func ls(files []fs.DirEntry) []string {
	var fnames []string
	for _, f := range files {
		fnames = append(fnames, f.Name())

	}
	return fnames
}
