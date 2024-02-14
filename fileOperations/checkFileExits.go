package fileoperations

import (
	"io/fs"
	"log"
	"os"
)

var (
	fileInfo fs.FileInfo
	myErr    error
)

func Demo4() {

	fileInfo, myErr = os.Stat("./fileOperations/MovedFile.js")

	if myErr != nil {
		if os.IsNotExist(myErr) {
			log.Fatal("File Do Not Exits")
		}
	} else {
		log.Println("File Do Exits")
		log.Printf(fileInfo.Name())
	}

}
