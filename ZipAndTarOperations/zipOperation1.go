package zipandtaroperations

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

func Demo3() {

	var result int = createZipArchiveFile("efehancer58")
	if result > 0 {
		log.Println("success")
	} else {
		log.Println("fatal")
	}

}

var fileFolderPath string = "./ZipAndTarOperations/files/"
var filesArr = []string{fileFolderPath + "demo.go", fileFolderPath + "note1.txt"}

func createZipArchiveFile(zipName string) int {
	if len(zipName) == 0 {
		log.Println("Zip Adı Eksik Hata")
		return -1
	}
	flag := os.O_RDONLY | os.O_CREATE | os.O_TRUNC
	zipFile, err := os.OpenFile(zipName+".zip", flag, 0644)
	if err != nil {
		log.Println("Zip oluşturma veyahut açma hatası")
		return -1
	}
	defer zipFile.Close()
	zw := zip.NewWriter(zipFile)
	defer zw.Close()
	for _, file := range filesArr {
		if err := addZipFile(file, zw); err != nil {
			log.Fatal("Zip ekleme hatası")
			return -1
		}
	}

	return 1
}

func addZipFile(filename string, zipWritter *zip.Writer) error {

	if len(filename) == 0 {
		return fmt.Errorf("Dosya Adı Girilmemiş hata")
	}

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Dosya Açma Hatası")
	}
	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Dosya Bilgi Alma Hatası")
	}
	defer file.Close()
	writter, err := zipWritter.Create(filename)
	if err != nil {
		return fmt.Errorf("Zip içine dosya oluşturma hatası")
	}
	copiedLen, err := io.Copy(writter, file)
	if err != nil {
		return fmt.Errorf("Dosya yazma kopyalama hatası")
	}
	if stat.Size() > copiedLen {
		return fmt.Errorf("Dikkat yazım var Ama veri boyutu eksik lütfen kontrol edin")
	}
	return nil
}
