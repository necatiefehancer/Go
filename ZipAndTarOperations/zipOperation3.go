package zipandtaroperations

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func OpenZipFile() {

	ZipReader, err := zip.OpenReader("zip4458.zip")
	if err != nil {
		log.Fatal("Zip Açma Hatası ", err.Error())
	}
	defer ZipReader.Close()
	for _, file := range ZipReader.Reader.File {
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal("Zip içi dosya Açma Hatası ", err.Error())
		}
		defer zippedFile.Close()
		dirPath := "./"
		exrtactedFilePath := filepath.Join(dirPath, file.Name)
		if file.FileInfo().IsDir() {
			log.Println("Bir klasör tespit edildi adı ", file.Name)
			err = os.MkdirAll(exrtactedFilePath, file.Mode())
			if err != nil {
				log.Println("Klasör oluşturulamadı !")
			}
		} else {
			outFile, err := os.OpenFile(
				exrtactedFilePath,
				os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
				file.Mode(),
			)
			if err != nil {
				log.Println("Çıkarılan Dosya Oluşturma Hatası !")
			}
			defer outFile.Close()
			copiedData, err := io.Copy(outFile, zippedFile)
			if err != nil {
				log.Fatal("Dosya kopyalama hatası")
			}
			log.Println("Kopyalanan veri ", copiedData)
		}
	}

}
