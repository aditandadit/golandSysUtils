package sysUtils

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args;
	if len(arguments) == 1 {
		fmt.Println("No args")
		os.Exit(1);
	}

	file := arguments[1];
	err := os.Remove(file);
	if err != nil {
		fmt.Println(err);
		return;
	}
}
