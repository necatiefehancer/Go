package fileoperations

import (
	"io"
	"log"
	"os"
)

func Demo7() {

	var (
		originalFile *os.File
		newFile      *os.File
		err          error
		bytesWritten int64
	)

	originalFile, err = os.Open("./fileOperations/folder/app.js")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	newFile, err = os.Create("./fileOperations/folder/copy_app.js")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	//Byteları kaynaktan kopyala

	bytesWritten, err = io.Copy(newFile, originalFile)

	log.Printf("Şu Kadar byte Kopyalandı %d \n", bytesWritten)

	// bellekteki verileri fiziksel olarak diske boşalt

	err = newFile.Sync()

}
