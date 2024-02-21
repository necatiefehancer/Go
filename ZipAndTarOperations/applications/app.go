package applications

import (
	"fmt"
	"log"
	"os"
)

func app() {
	f, err := os.Create("efh.json")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(f.Name())
	}
}
