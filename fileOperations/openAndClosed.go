package fileoperations

import (
	"log"
	"os"
)

func Demo5() {

	// file, err := os.Open("./fileOperations/MovedFile.js")

	// defer file.Close()

	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	//üzerine veri yazdın
	// }

	file, err := os.OpenFile("./fileOperations/MovedFile.js", os.O_APPEND, 0666)

	if err != nil {
		log.Fatal(err)
	}
	file.Close()

}
