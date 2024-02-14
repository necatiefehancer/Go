package timeanddateoperations

import (
	"fmt"
	"time"
)

func Demo1() {

	fmt.Printf("Şu Anki Tarihin Unix Zamanı %v \n", time.Now().Unix())

	time.Sleep(time.Second * 1)

	fmt.Printf("Şu Anki Tarihin Unix Zamanı %v \n", time.Now().Unix())

	fmt.Println(time.Now())

	fmt.Println("*********************************")

	var myBirthDay time.Time = time.Date(2004, time.September, 6, 0, 0, 0, 0, time.UTC)

	fmt.Printf("Doğum Tarihin : %v \n", myBirthDay)

	fmt.Println("*********************************")

	var nowTime time.Time = time.Now()

	fmt.Println(nowTime)

	fmt.Println("*********************************")

	fmt.Println("Ay : ", nowTime.Month())

	fmt.Println("Gün : ", nowTime.Day())

	fmt.Println("Haftanın Günü : ", nowTime.Weekday())

	var TomorrowTime time.Time = nowTime.AddDate(-1, 1, 1)

	fmt.Println("*********************************")

	fmt.Printf("eklenmiş TÜretilmiş Tarih %v \n", TomorrowTime)

	var longFormat string = "2024 , January 7, Monday"

	fmt.Printf("Tomorrow is long Format : %v\n", TomorrowTime.Format(longFormat))
}
