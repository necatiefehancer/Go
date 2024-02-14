package fileoperations

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

var (
	fileİnfo fs.FileInfo
	myerr    error
)

func Demo2() {

	fileİnfo, myerr = os.Stat("file.js")

	if myerr != nil {
		log.Fatal(myerr)
	} else {

		fmt.Println("File Name ", fileİnfo.Name())

		fmt.Println("Filse Size ", fileİnfo.Size())

		fmt.Println("File Permissions ", fileİnfo.Mode())

		fmt.Println("Last Modified ", fileİnfo.ModTime())

		fmt.Println("The Path İS DİR ? ", fileİnfo.IsDir())

		fmt.Println("System İnterface Type ",fileİnfo.Sys())
	}

}
