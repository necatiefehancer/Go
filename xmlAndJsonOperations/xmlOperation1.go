package xmlandjsonoperations

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type ObjectSites struct {
	XMLName     xml.Name  `xml:"sites"`
	Versions    string    `xml:"version,attr"`
	Description string    `xml:",innerxml"`
	WebSites    []WebSite `xml:"site"`
}

type WebSite struct {
	XMLName     xml.Name `xml:"site"`
	Name        string   `xml:"Name"`
	Description string   `xml:"Description"`
	Category    string   `xml:"Category"`
}

func Demo3() {

	var (
		xmlFile      *os.File
		err          error
		value        ObjectSites
		readByteData []byte
	)

	xmlFile, err = os.OpenFile("./xmlAndJsonOperations/sites.xml", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	readByteData, err = ioutil.ReadAll(xmlFile)
	if err != nil {
		panic(err)
	}
	err = xml.Unmarshal(readByteData, &value)
	if err != nil {
		panic(err)
	}
	fmt.Println("Succes Writing Object Xml File !")
	fmt.Println(value)
	xmlFile.Close()

	time.Sleep(time.Second * 5)
	fmt.Println("data changed repeat writing data file uptade started !")

	value.WebSites[0].Name = "hancerlerlogistick.com"
	value.WebSites[0].Category = "AKARYAKIT İŞLETMESİ"
	value.WebSites[0].Description = "HANCERLER AKARYAKIT"

	xmlFile, err = os.OpenFile("./xmlAndJsonOperations/sites.xml", os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	readByteData, err = xml.Marshal(value)
	if err != nil {
		panic(err)
	}
	_, err = xmlFile.Write(readByteData)
	if err != nil {
		panic(err)
	}
	fmt.Println("File Data Changed Success")

	xmlFile.Close()

}
