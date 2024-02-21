package zipandtaroperations

import (
	"archive/tar"
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var appFolderPath string = "./ZipAndTarOperations/applications/"
var appFiles = []string{
	appFolderPath + "data.json",
	appFolderPath + "app.js",
	appFolderPath + "app.go",
}

func createTarFile(tarname string) int {
	if len(tarname) == 0 {
		log.Println("Dosya adı girilmemiş")
		return -1
	}
	flags := os.O_RDONLY | os.O_TRUNC | os.O_CREATE
	tarfile, err := os.OpenFile(tarname+".tar", flags, 0644)
	if err != nil {
		log.Fatal("Tar dosya Açma Hatası")
		return -1
	}
	defer tarfile.Close()
	tarWritter := tar.NewWriter(tarfile)
	defer tarWritter.Close()
	for _, file := range appFiles {
		if err := addTarFile(file, tarWritter); err != nil {
			log.Fatal("Dosya ekleme hatası dosya adı %s", file)
			return -1
		}
	}
	return 1
}

func addTarFile(fileName string, tarWritter *tar.Writer) error {
	if len(fileName) == 0 {
		return fmt.Errorf("Dosya ismi yok hata !")
	}
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Dosya Açma Hatası")
	}
	defer file.Close()
	fileStat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Dosya bilgilerini okuma hatası")
	}
	tarheader := &tar.Header{
		Name:    fileStat.Name(),
		Size:    fileStat.Size(),
		ModTime: fileStat.ModTime(),
		Mode:    int64(fileStat.Mode().Perm()),
	}
	if err := tarWritter.WriteHeader(tarheader); err != nil {
		return fmt.Errorf("dosyayı tar header ` a ekleme hatası")
	}
	copiedData, err := io.Copy(tarWritter, file)
	if err != nil {
		return fmt.Errorf("dosyayı veri yazma hatası")
	}
	if copiedData < fileStat.Size() {
		return fmt.Errorf("Dikkat veri yazılan ve okunan veri boyutları eşlemşmiyor dosyaları kontrol edin")
	}
	return nil
}

func createZipFile(zipname string) int {
	if len(zipname) == 0 {
		log.Fatal("zip ismi girilmemiş 0 karkater")
		return -1
	}
	flags := os.O_RDONLY | os.O_CREATE | os.O_TRUNC
	zipfile, err := os.OpenFile(zipname+".zip", flags, 0644)
	if err != nil {
		log.Println("Zip dosya açım hatası !")
	}
	defer zipfile.Close()
	zipWritter := zip.NewWriter(zipfile)
	defer zipWritter.Close()
	for _, file := range appFiles {
		if err := addzipFile(file, zipWritter); err != nil {
			log.Fatal("Dosya ekleme hatası")
			return -1
		}
	}
	return 1
}

func addzipFile(filename string, zipwritter *zip.Writer) error {
	if len(filename) == 0 {
		return fmt.Errorf("dosya adı yok geçersiz 0 karakter !")
	}
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("dosya açma hatası")
	}
	fileStat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Dosya bilgi Alma hatası")
	}
	defer file.Close()
	zw, err := zipwritter.Create(filename)
	if err != nil {
		return fmt.Errorf("Zip içine dosya oluşturma hatası")
	}
	copiedData, err := io.Copy(zw, file)
	if err != nil {
		return fmt.Errorf("Zip içine dosya yazma hatası !")
	}
	if fileStat.Size() > copiedData {
		return fmt.Errorf("Dikkat veri yazılan ve okunan veri boyutları eşlemşmiyor dosyaları kontrol edin")
	}
	return nil
}

func SelectPackageType() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Oluşturmak İstediğiniz Arşiv Adını Giriniz")
	filename, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("String input okuma hatası")
	}
	filename = strings.TrimSpace(filename)

	fmt.Println("Oluşturmak İstediğiniz Arşiv Tipini Seçiniz (Zip,Tar)")
	data, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("String input okuma hatası")
	}
	data = strings.TrimSpace(data)

	switch data {
	case "Zip":
		createZipFile(filename)
	case "Tar":
		createTarFile(filename)
	default:
		fmt.Println("Geçersiz Kod")
		SelectPackageType()
	}
}
