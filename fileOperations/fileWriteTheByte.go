package fileoperations

import (
	"log"
	"os"
)

func Demo8() {

	//Dosyanı sadece yazılabilir olarak aç

	var (
		file *os.File
		err  error
	)

	defer file.Close()

	file, err = os.OpenFile("./fileOperations/folder/alert.js", os.O_APPEND, 0666)

	if err != nil {
		if os.IsPermission(err) {
			log.Println("Bu Dosyaya Ekleme Yapılamaz")
		}
	} else {

		byteSlices := []byte("console.log(58)\nalert(`gürün58`)")

		bytesWritten, err := file.Write(byteSlices)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Şu Kadar Byte Veri Yazıldı %d \n", bytesWritten)

	}

}
