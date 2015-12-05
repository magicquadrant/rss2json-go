package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		fmt.Println("Provide a filename on command line")
		os.Exit(1)
	}

	filename := arg[0]
	_, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("file reading failed:", err)
		os.Exit(1)
	}
}
