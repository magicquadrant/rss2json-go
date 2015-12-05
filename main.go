package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := "html5rocks.xml"
	_, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("file reading failed:", err)
		os.Exit(1)
	}
}
