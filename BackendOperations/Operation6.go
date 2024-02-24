package backendoperations

import (
	"bytes"
	"log"
	"net/http"
	"text/template"
)

type Product struct {
	ProductName        string
	ProductId          string
	SellPrice          float64
	BuyPrice           float64
	ID                 string
	ProductDescription string
}

func Op6() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var builder bytes.Buffer
		builder.WriteString("Desni İçme Suyu 500 ml \n")
		builder.WriteString("Erzurum PalanDöken A.ş tesislerinde Üretilmiştir\n")
		builder.WriteString("pH Değeri 7.23")

		var Product1 = Product{
			ProductName:        "Desni Su 0.5L",
			ProductId:          "0168e83",
			SellPrice:          3.68,
			BuyPrice:           5.00,
			ID:                 "86905041728830",
			ProductDescription: builder.String(),
		}

		template, err := template.ParseFiles("./BackendOperations/page.html")
		if err != nil {
			log.Fatal("Template parse etme hatası ", err.Error())
		}
		template.Execute(w, Product1)

	})

	http.ListenAndServe(":8080", nil)

}
