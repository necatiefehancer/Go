package backendoperations

import (
	"fmt"
	"log"
	"net/http"
)

type Human struct {
	Name    string
	Surname string
	Age     int
}

func (h Human) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Name = "EFE"
	h.Surname = "HANÇER"
	h.Age = 20

	//FORMU PARSE ET
	r.ParseForm()

	// formu yaz
	fmt.Println(r.Form)

	table := fmt.Sprintf("<div> Name %s</div>\n"+
		"<div> Surname %s</div>\n"+
		"<div> Age %d</div>\n"+
		"<div> Path %s</div>", h.Name, h.Surname, h.Age, r.URL.Path[1:])

	fmt.Println(r.URL.Path[1:])

	fmt.Fprintf(w, table)
}

func Op2() {

	var hum Human
	err := http.ListenAndServe("localhost:9000", hum)
	if err != nil {
		log.Fatal(err)
	}

}
