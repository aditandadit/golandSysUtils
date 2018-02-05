package sysUtils

import (
	"fmt"
	"os"
	"path/filepath"
	"flag"
)

func walkFunction(path string, info os.FileInfo, err error) error {
	_,err = os.Stat(path);
	if err != nil {
		return err;
	}
	fmt.Println(path);
	return nil;
}

func walkOnlyDirs(path string, info os.FileInfo, err error) error {
	fileInfo , err := os.Stat(path);
	if err!= nil {
		return err;
	}

	mode := fileInfo.Mode();
	if mode.IsDir() {
		fmt.Println(path);
	}
	return nil;
}

func isExcluded(name string, exclude string) bool {
	if exclude == "" { return false; }
	if filepath.Base(name) == exclude { return true; }
	return false;

}

func main() {
	Path := flag.String("file", "","file" );
	dirOnly := flag.Bool("dir", false, "dirOnly");
	minusX := flag.String("x", "", "Files");
	flag.Parse();

	var err error;
	if (*dirOnly == true) {
		err = filepath.Walk(*Path, walkOnlyDirs);
	} else {
		err = filepath.Walk(*Path, walkFunction);
	}
	if err != nil {
		fmt.Println(err);
		os.Exit(1);
	}
}