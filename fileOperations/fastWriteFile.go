package fileoperations

import (
	"io/ioutil"
	"log"
)

func Demo9() {

	err := ioutil.WriteFile("test.go", []byte("fmt.Println(52328)"), 0666)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Writing File Success")
	}

}
