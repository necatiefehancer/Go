package zipandtaroperations

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

var fileFolerPath string = "./ZipAndTarOperations/files/"
var files = []string{fileFolerPath + "demo.go", fileFolerPath + "note1.txt"}

func addFile(fileName string, tw *tar.Writer) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Dosya Açılrıken Bir hata meydana geldi %s %s", fileName, err.Error())
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Dosya Bilgileri Alırken Hata Geldi %s %s ", fileName, err.Error())
	}
	// hdrler dosyanın temel bilgielrini erişip işlem yaptığın alanlar
	hdr := &tar.Header{
		ModTime: stat.ModTime(),
		Name:    stat.Name(),
		Size:    stat.Size(),
		Mode:    int64(stat.Mode().Perm()),
	}
	if err := tw.WriteHeader(hdr); err != nil {
		return fmt.Errorf("Tar Header Yazılırken Hata Meydana Geldi %s %s", fileName, err.Error())
	}
	copied, err := io.Copy(tw, file)
	if err != nil {
		return fmt.Errorf("%s Dosyası tar a yazılırken hata meydana geldi %s", fileName, err.Error())
	}
	if copied < stat.Size() {
		return fmt.Errorf("Eksik Yazma İşlemi olabilir Beklenen %d Kopyalanan", stat.Size(), copied)
	}
	return nil
}

func createArchiveFileTar(archiveFileName string) int {
	if len(archiveFileName) == 0 {
		return -1
	}
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile("./ZipAndTarOperations/"+archiveFileName+".tar", flags, 0644)
	if err != nil {
		log.Fatalf("Tar Dosyasına Yazmak için Açılırken Hata Meydana geldi %s", err.Error())
		return -1
	}
	defer file.Close()
	tw := tar.NewWriter(file)
	defer tw.Close()
	for _, filename := range files {
		if err := addFile(filename, tw); err != nil {
			log.Fatal("%s Dosyası Eklenirken Hta geldi %s", filename, err.Error())
		}

	}
	return 1
}

func Demo1() {

	result := createArchiveFileTar("efe")
	if result > 0 {
		log.Println("Write Success")
	} else {
		log.Println("Write NotSuccess")
	}

}
