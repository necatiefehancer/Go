package zipandtaroperations

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

func AddFile(filepath string, tarWritter *tar.Writer) error {
	if len(filepath) == 0 {
		return fmt.Errorf("Dosya İsim Boş Geçilmiş Hata ")
	}
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("Dosya Açma Hatası")
	}
	defer file.Close()
	fileStat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Dosya Bilgi Okuma Hatası")
	}
	tarheader := &tar.Header{
		Name:    fileStat.Name(),
		ModTime: fileStat.ModTime(),
		Mode:    int64(fileStat.Mode().Perm()),
		Size:    fileStat.Size(),
	}
	if err := tarWritter.WriteHeader(tarheader); err != nil {
		return fmt.Errorf("Yazma Hatası")
	}
	copiedData, err := io.Copy(tarWritter, file)
	if err != nil {
		return fmt.Errorf("copy hata")
	}
	if copiedData < fileStat.Size() {
		return fmt.Errorf("Eksik Yzma hatası !")
	}
	return nil
}

func createTar(tarName string) int {
	if len(tarName) == 0 {
		return -1
	}
	flag := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(tarName, flag, 0644)
	if err != nil {
		return -1
	}
	defer file.Close()
	tarWritter := tar.NewWriter(file)
	defer tarWritter.Close()
	if AddFile("./ZipAndTarOperations/files/demo.go", tarWritter) != nil {
		log.Fatal("hata 1")
	}
	if AddFile("./ZipAndTarOperations/files/note1.txt", tarWritter) != nil {
		log.Fatal("hata 2")
	}
	return 1
}
func Demo2() {

	if createTar("./ZipAndTarOperations/data.tar") > 0 {
		log.Println("success")
	} else {
		log.Println("Fatall")
	}

}
