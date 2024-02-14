package fileoperations

import (
	"log"
	"os"
)

func Demo6() {

	// file, err := os.OpenFile("demo.txt", os.O_WRONLY, 0666)

	// if err != nil {
	// 	if os.IsPermission(err) {
	// 		log.Println("Yazma İzni Yok Bu Dosyanın")
	// 	}
	// }

	// file.Close()

	file, err := os.OpenFile("demo.txt", os.O_RDONLY, 0666)

	if err != nil {
		if os.IsPermission(err) {
			log.Println("Okuma İzni Yok Bu Dosyanın")
		}
	}

	file.Close()

}
