package xmlandjsonoperations

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func Demo1() {

	jsonstr := `
	{
		"data":{
			"name":"efe",
			"surname":"hancer"
		}
	}
	`

	var myObject map[string]map[string]interface{}

	err := json.Unmarshal([]byte(jsonstr), &myObject)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(myObject["data"]["name"])
	}

	file, err := os.Open("./xmlAndJsonOperations/info.json")
	if err != nil {
		panic(err)
	} else {
		newFile, err := os.Create("./xmlAndJsonOperations/info2.json")
		if err != nil {
			panic(err)
		} else {
			bytesWritten, err := io.Copy(newFile, file)
			if err != nil {
				panic(err)
			} else {
				log.Println("kopaylanan dosyaya şu kadar byte json data yazıldı ", bytesWritten)
				newFile.Sync()
			}
		}

	}

	defer file.Close()
}
