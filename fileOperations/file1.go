package fileoperations

import (
	"log"
	"os"
)

var (
	myFile *os.File
	err    error
)

func Demo1() {

	myFile, err = os.Create("file.js")

	if err != nil {
		log.Fatal(err)
	}

}
