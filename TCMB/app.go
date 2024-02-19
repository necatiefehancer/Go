package tcmb

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
)

type CurrencyDay struct {
	ID         string
	Date       time.Time
	DayNo      string
	Currencies []GoCurrency
}

type GoCurrency struct {
	Code            string
	CrossOrder      int
	Unit            int
	CurrencyNameTR  string
	CurrencyName    string
	ForexBuying     float64
	ForexSelling    float64
	BanknoteBuying  float64
	BanknoteSelling float64
	CrossRateUSD    float64
	CrossRateOther  float64
}

type tarih_Date struct {
	XMLName     xml.Name `xml:"Tarih_Date"`
	Tarih       string   `xml:"Tarih,attr"`
	Date        string   `xml:"Date,attr"`
	Bulten_No   string   `xml:Bulten_No,attr`
	Currencyies []Currency
}

type Currency struct {
	XMLName         xml.Name `xml:"Currency"`
	CrossOrder      string   `xml:"CrossOrder,attr"`
	Kod             string   `xml:"Kod,attr"`
	CurrencyCode    string   `xml:"CurrencyCıde,attr"`
	Unit            string   `xml:"Unit"`
	Isim            string   `xml:"Isim"`
	CurrencyName    string   `xml:"CurrencyName"`
	ForexBuying     string   `xml:"ForexBuying"`
	ForexSelling    string   `xml:"ForexSelling"`
	BanknoteBuying  string   `xml:"BanknoteBuying"`
	BanknoteSelling string   `xml:"BanknoteSelling"`
	CrossRateUSD    string   `xml:"CrossRateUSD`
	CrossRateOther  string   `xml:"CrossRateOther"`
}

func (c *CurrencyDay) GetData(CurrencyDate time.Time) {
	xDate := CurrencyDate
	t := new(tarih_Date)
	currDay := t.getDate(CurrencyDate, xDate)

	for {
		if currDay == nil {
			CurrencyDate = CurrencyDate.AddDate(0, 0, -1)
			currDay := t.getDate(CurrencyDate, xDate)
			if currDay != nil {
				break
			} else {
				break
			}
		}
	}
	SaveToJson("data", currDay)

}

func (t *tarih_Date) getDate(CurrencyDate time.Time, xDate time.Time) *CurrencyDay {
	currDay := new(CurrencyDay)
	var response *http.Response
	var err error
	var url string

	currDay = new(CurrencyDay)
	url = "https://www.tcmb.gov.tr/kurlar/today.xml"
	response, err = http.Get(url)
	fmt.Println(url)
	if err != nil {
		fmt.Println(err)
	} else {
		defer response.Body.Close()
		if response.StatusCode != http.StatusNotFound {
			tarih := new(tarih_Date)
			d := xml.NewDecoder(response.Body)
			marshalErr := d.Decode(&tarih)
			if marshalErr != nil {
				log.Println("Marshal err")
			}

			// c := &tarih_Date{}
			currDay.ID = xDate.Format("20060102")
			currDay.Date = xDate
			currDay.DayNo = tarih.Bulten_No
			currDay.Currencies = make([]GoCurrency, len(tarih.Currencyies))
			for i, curr := range tarih.Currencyies {
				currDay.Currencies[i].Code = curr.CurrencyCode
				currDay.Currencies[i].BanknoteBuying, _ = strconv.ParseFloat(curr.BanknoteBuying, 64)
				currDay.Currencies[i].BanknoteSelling, _ = strconv.ParseFloat(curr.BanknoteSelling, 64)
				currDay.Currencies[i].CrossOrder, _ = strconv.Atoi(curr.CrossOrder)
				currDay.Currencies[i].CrossRateOther, _ = strconv.ParseFloat(curr.CrossRateOther, 64)
				currDay.Currencies[i].CurrencyNameTR = curr.CurrencyName
				currDay.Currencies[i].ForexBuying, _ = strconv.ParseFloat(curr.ForexBuying, 64)
				currDay.Currencies[i].ForexSelling, _ = strconv.ParseFloat(curr.ForexSelling, 64)
				currDay.Currencies[i].Unit, _ = strconv.Atoi(curr.Unit)
				currDay.Currencies[i].CurrencyName = curr.CurrencyName
				currDay.Currencies[i].CrossRateUSD, _ = strconv.ParseFloat(curr.CrossRateUSD, 64)
				fmt.Println(curr)
			}
		} else {
			currDay = nil
		}
	}
	return currDay
}

func SaveToJson(filename string, key interface{}) {
	outFile, err := os.Create("./TCMB/" + filename + ".json")
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(outFile)
	err2 := encoder.Encode(key)
	if err2 != nil {
		log.Fatal(err2)
	}
	log.Println("write Success")
	defer outFile.Close()
}

func Demo1() {
	runtime.GOMAXPROCS(2)
	startTime := time.Now()
	CurrencyDay := new(CurrencyDay)
	CurrencyDate := time.Now()
	CurrencyDay.GetData(CurrencyDate)

	elepsedTime := time.Since(startTime)
	fmt.Printf("Execution Time %s", elepsedTime)

}
