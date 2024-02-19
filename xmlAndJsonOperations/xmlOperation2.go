package xmlandjsonoperations

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Person struct {
	XMLName   xml.Name `xml:"person"`
	ID        int      `xml:"id"`
	FirstName string   `xml:"firstname"`
	LastName  string   `xml:"lastname"`
	UserName  string   `xml:"username"`
}

type Company struct {
	XMLName  xml.Name `xml:"company"`
	Persons  []Person `xml:"person"`
	Version  string   `xml:"verison,attr"`
	Encoding string   `xml:"encoding,attr"`
}

type JsonPerson struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
}

func (p Person) GetPersonİnfo() string {
	return fmt.Sprintf("/t ID : %d FirstName : %s LastName: %s UserName: %s", p.ID, p.FirstName, p.LastName, p.UserName)
}

func Demo4() {

	var (
		xmlFile     *os.File
		jsonFile    *os.File
		err         error
		dataBytes   []byte
		company     Company
		jsonEncoder *json.Encoder
		jsonPerson  JsonPerson
		jsonPersons []JsonPerson
	)

	xmlFile, err = os.OpenFile("./xmlAndJsonOperations/Employees.xml", os.O_RDONLY, 0666)
	defer xmlFile.Close()
	if err != nil {
		panic(err)
	}
	dataBytes, err = ioutil.ReadAll(xmlFile)
	if err != nil {
		panic(err)
	}
	err = xml.Unmarshal(dataBytes, &company)
	if err != nil {
		panic(err)
	}
	log.Println("File Read Succees !")
	fmt.Println(company.Persons)

	time.Sleep(time.Second * 5)
	fmt.Println("Xml To Json Convert Started !")

	jsonFile, err = os.Create("./xmlAndJsonOperations/Employees.json")
	if err != nil {
		panic(err)
	}

	for _, p := range company.Persons {
		jsonPerson.FirstName = p.FirstName
		jsonPerson.LastName = p.LastName
		jsonPerson.ID = p.ID
		jsonPerson.UserName = p.UserName
		jsonPersons = append(jsonPersons, jsonPerson)
	}
	jsonEncoder = json.NewEncoder(jsonFile)
	err = jsonEncoder.Encode(&jsonPersons)
	if err != nil {
		panic(err)
	}
	log.Println("translate success !")

}
