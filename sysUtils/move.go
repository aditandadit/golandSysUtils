package sysUtils

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	overWrite := flag.Bool("overwrite", false, "overwrite");
	flag.Parse();
	flags := flag.Args();

	if len(flags) < 2 {
		fmt.Println("Please provide two arguments!")
		os.Exit(1)
	}
	source := flags[0];
	destination := flags[1];
	fileinfo, err := os.Stat(source);
	if err == nil {
		mode := fileinfo.Mode();
		if mode.IsRegular() == false {
			fmt.Println("Only Regular Files supported");
			os.Exit(1);
		}
	} else {
		fmt.Println("Error Reading: ",source);
		os.Exit(1);
	}
	newDest := destination;
	destInfo, err := os.Stat(destination);
	if err == nil {
		mode := destInfo.Mode();
		if mode.IsDir() {
			name := filepath.Base(source);
			newDest = destination + "/" + name;
		}
		destination = newDest;
		destInfo, err = os.Stat(destination);
		if err == nil {
			if *overWrite == false {
				fmt.Println("file already exitst");
			}
		}
		err = os.Rename(source, destination);
		if err != nil {
			fmt.Println(err);
			os.Exit(1);
		}
	}
}