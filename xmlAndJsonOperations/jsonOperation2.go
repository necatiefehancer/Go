package xmlandjsonoperations

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Product struct {
	ProductName string
	ProductCode int
	SellPrice   float64
	BuyPrice    float64
	CreatorFirm []Firm
}
type Products struct {
	Products []Product
}

type Firm struct {
	ID         int
	FirmName   string
	FirmEmail  []Email
	FirmAdress []Adress
}

type Adress struct {
	Adress string
}

type Email struct {
	Email string
}

func (p *Product) GetProductInfo() string {
	return p.ProductName + " " + strconv.Itoa(p.ProductCode) + "Buy Price " + strconv.FormatFloat(p.BuyPrice, 'f', 2, 64) + " Sell Price " + strconv.FormatFloat(p.SellPrice, 'f', 2, 64)
}

func (f *Firm) GetFirmİnfo(index int) string {
	return f.FirmName + " " + strconv.Itoa(f.ID) + " " + f.FirmAdress[index].Adress + f.FirmEmail[index].Email
}

func CheckError(err error) (state bool) {
	if err != nil {
		state = true
		panic(err)
	} else {
		state = false
	}
	return
}

func WriteToDataJson(filePath string, object interface{}) {
	newDataFile, err := os.Create("./xmlAndJsonOperations/" + filePath)
	if !CheckError(err) {
		jsonEncoding := json.NewEncoder(newDataFile)
		err2 := jsonEncoding.Encode(object)
		if !CheckError(err2) {
			log.Println("Write success")
		}
	}
	defer newDataFile.Close()

}

func ReadToDataJson(path string) (DataProduct Products) {
	dataFile, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if !CheckError(err) {
		decoder := json.NewDecoder(dataFile)
		err2 := decoder.Decode(&DataProduct)
		if !CheckError(err2) {
			log.Println("Read Success")
		}
	}
	defer dataFile.Close()
	return
}

func Demo2() {

	var ürünler = Products{Products: []Product{
		Product{
			ProductName: "Damla Su 500ML",
			ProductCode: 8692000158,
			SellPrice:   3.68,
			BuyPrice:    5.00,
			CreatorFirm: []Firm{
				Firm{
					ID:       18,
					FirmName: "CocaCola A.ş",
					FirmEmail: []Email{
						Email{Email: "cooke@gmail.com"},
						Email{Email: "anadoluicecekdestek@hotmail.com.tr"},
					},
					FirmAdress: []Adress{
						Adress{Adress: "Eskişehir/Türkiye"},
						Adress{Adress: "İstanbul Avrupa Beşiktaş"},
					},
				},

				Firm{
					ID:       19,
					FirmName: "PhilipMorsisA.ş",
					FirmEmail: []Email{
						Email{Email: "Phs@gmail.com"},
						Email{Email: "phs.gov.tr"},
					},
					FirmAdress: []Adress{
						Adress{Adress: "İstanbul Avrupa"},
						Adress{Adress: "Ankara Keçiören"},
					},
				},
			},
		},
		Product{
			ProductName: "EfeGabcer",
			ProductCode: 8692000158,
			SellPrice:   3.68,
			BuyPrice:    5.00,
			CreatorFirm: []Firm{
				Firm{
					ID:       18,
					FirmName: "CocaCola A.ş",
					FirmEmail: []Email{
						Email{Email: "cooke@gmail.com"},
						Email{Email: "anadoluicecekdestek@hotmail.com.tr"},
					},
					FirmAdress: []Adress{
						Adress{Adress: "Eskişehir/Türkiye"},
						Adress{Adress: "İstanbul Avrupa Beşiktaş"},
					},
				},

				Firm{
					ID:       19,
					FirmName: "PhilipMorsisA.ş",
					FirmEmail: []Email{
						Email{Email: "Phs@gmail.com"},
						Email{Email: "phs.gov.tr"},
					},
					FirmAdress: []Adress{
						Adress{Adress: "İstanbul Avrupa"},
						Adress{Adress: "Ankara Keçiören"},
					},
				},
			},
		},
	}}
	fmt.Println(ürünler)

	// WriteToDataJson("products.json", ürünler.products)

	time.Sleep(time.Second * 5)

	data := ReadToDataJson("./xmlAndJsonOperations/products.json")
	fmt.Println(data)

}
