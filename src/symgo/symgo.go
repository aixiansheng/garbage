package main

import (
	"flag"
	"fmt"
	"os"
	"symcache"
	"path/filepath"
)

func VisitFile(path string, visitor symcache.MachFileVisitor) error {
	if fh, err := os.Open(path); err != nil {
		return nil
	} else {
		defer fh.Close()
		if err = visitor.VisitFile(path, fh); err != nil {
			fmt.Printf("file error: %v\n", err)
		}
		
		return nil
	}
}

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("%s [OPTIONS] DIR [DIR...]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		fmt.Errorf("You must specify dirs to search")
		os.Exit(1)
	}

	dirs := flag.Args()
	visitor := &symcache.SymbolPrinter{}

	for _, dir := range(dirs) {
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				return nil
			}

			return VisitFile(path,  visitor)
		})
	}
}
