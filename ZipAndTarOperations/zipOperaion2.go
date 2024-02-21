package zipandtaroperations

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func Demo4() {

	zipReader, err := zip.OpenReader("zip4458.zip")

	if err != nil {
		fmt.Errorf("Reader Açım Hatası ", err.Error())
	}
	defer zipReader.Close()
	for _, file := range zipReader.Reader.File {
		zippedFile, err := file.Open()
		if err != nil {
			fmt.Errorf("Zip içindeki dosya açım hatası")
		}
		defer zippedFile.Close()
		targetDir := "./"
		extractedFilePath := filepath.Join(targetDir, file.Name)
		if file.FileInfo().IsDir() {
			log.Println("Bu Yapı bir klasör ve oluşturuluyor ")
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			outFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				fmt.Errorf("dosya oluşturulamadı")
			}
			defer outFile.Close()
			_, err = io.Copy(outFile, zippedFile)
			if err != nil {
				fmt.Errorf("dosya kopayalama hatası")
			}

		}
	}

}
