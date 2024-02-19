package xmlandjsonoperations

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Demo5() {

	csvFile, err := os.OpenFile("./xmlAndJsonOperations/data.csv", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	csvReader := csv.NewReader(csvFile)
	csvReader.Read()
	csvRecordData, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, dataArray := range csvRecordData {
		for i, v := range dataArray {
			fmt.Printf("[%d] %s \n", i, v)
		}
	}

}
