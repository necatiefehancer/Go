package fileoperations

import (
	"fmt"
	"log"
	"os"
)

func Demo3() {

	var (
		originalPath string = "file.js"
		newPath      string = "./fileOperations/MovedFile.js"
		visitErr     error
	)

	visitErr = os.Rename(originalPath, newPath)

	if visitErr != nil {
		log.Fatal(visitErr)
	} else {
		fmt.Println("Başarıyla dosya taşındı")
	}

}
