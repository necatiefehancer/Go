package fileoperations

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func Demo10() {

	//geçici klasörünü oluştur

	tempDir, err := ioutil.TempDir("./", "tempDir")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Geçici Klasör Oluşturuldu ", tempDir)

	//geçici dosyanı oluştur

	tempFile, err := ioutil.TempFile(tempDir, "tempFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Geçici Dosya Oluşturuldu ", tempFile.Name())

	writeLen, err2 := tempFile.WriteString("efehancer58")
	if err2 != nil {
		log.Fatal(err2)
	}
	log.Printf("Dosyaya Şu Kadar Veri Yazıldı %v", writeLen)

	tempFile.Close()
	time.Sleep(time.Second * 6)

	fmt.Println("Dosyalar Silinecek")

	err3 := os.Remove(tempFile.Name())
	if err3 != nil {
		log.Fatal(err3)
	}
	log.Println("dosya silindi")

	err4 := os.Remove(tempDir)
	if err4 != nil {
		log.Fatal(err4)
	}
	log.Println("dizin Silindi")

}
