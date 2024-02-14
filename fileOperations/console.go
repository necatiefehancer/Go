package fileoperations

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Console() {

	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.TrimSpace(data)

	file, err := os.OpenFile("./fileOperations/folder/console.js", os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
	} else {
		defer file.Close()
		bytesSlice := []byte(data)

		bytesWritten, err := file.Write(bytesSlice)

		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Yazım Tamamlandı Yazılan %d byte \n", bytesWritten)
		}
	}

}
