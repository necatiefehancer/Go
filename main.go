package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/rand"
	zipandtaroperations "myModules/ZipAndTarOperations"
	"myModules/concurrency"
	"myModules/functions"
	"myModules/interfaces"
	"myModules/utils"
	"os"
	"strconv"
	"strings"
	"time"

	///alias
	"github.com/fatih/color"
	loge "github.com/koding/logging"
)

type Languages string

const (
	JAVASCRİPT Languages = "Javascript D"
	C          Languages = "C"
	PYTHON     Languages = "Python"
	JAVA       Languages = "Java"
)

func PrintLanguage(lng Languages) {
	fmt.Println(lng)
}

func main() {

	// var myChan = make(chan bool)

	// go func() {

	// 	time.Sleep(time.Second * 5)
	// 	fmt.Println("MAİN DEĞER GELENE KADAR BEKLEYECEK")
	// 	myChan <- true

	// }()
	// <-myChan

	// xmlandjsonoperations.Demo5()
	// tcmb.Demo1()
	zipandtaroperations.Demo2()
}

func Concurrency() {
	concurrency.ABMain()
}

func structAndObjectInterfaceExample() {
	var ferrari interfaces.Ferrari
	ferrari.SetCarBrand("ferrari")
	ferrari.SetCarModel("F50")
	ferrari.SetCarColor("red")
	ferrari.SetCarSpeed(340)
	ferrari.SetCarPrice(1.98)
	ferrari.SetSpecial(true)
	fmt.Println(ferrari.Info())

	var mercedes interfaces.Mercedes
	mercedes.SetCarBrand("Mercedes")
	mercedes.SetCarModel("AXOR 1840")
	mercedes.SetCarColor("White")
	mercedes.SetCarSpeed(120)
	mercedes.SetCarPrice(1.1)
	fmt.Println(mercedes.Info())
}

func functionsForConsole() {

	fmt.Println(functions.Add(8, 9))
	functions.SayHello1()
	functions.SayHello2("Efe Hançer")

	var statusCode int = 73

	functions.Change1(statusCode)
	fmt.Println(statusCode)

	functions.Change1ByReference(&statusCode)
	fmt.Println(statusCode)
	var a, b string = functions.Swap("necati", "efe")
	fmt.Println(a, b)

	var myArray = []int{8, 16, 24, 36}
	var length, sum int = functions.VariadicSum(myArray...)
	fmt.Printf("length : %v sum : %v \n", length, sum)

	fmt.Println(functions.NamedVariableFunctions(10))

	var x, y int = functions.NamedVariableFunctions(20)
	fmt.Println("geriye değişken dönüşü yapan fonksyion x = ", x, "y = ", y)

	var mySlices = []int{9, 18, 27, 36, 45}
	var k, l int = functions.NamedVariableFunctionsVariadic(mySlices...)
	fmt.Printf("geriye değişken dönüşü yapan variadic fonkiyon sum degeri : %v len degeri : %v \n", k, l)

	functions.VariadicMessageFunctions("hello", "go", "develeopers", "backend")

	fmt.Println(functions.VariadicSumFunctions(9, 10, 98, 84, 73))

	functions.AnonimFunctions()

}

func PackageStatement() {

	//random paketi
	fmt.Println(rand.Intn(10))

	//strings contains

	fmt.Printf("test kelimesinin içinde est stringi %v\n", strings.Contains("test", "est"))

	// strings count
	fmt.Println(strings.Count("sivas", "a"))

	//hasprefix

	fmt.Println(strings.HasPrefix("unit_test", "unit"))

	//hassuffix

	fmt.Println(strings.HasSuffix("efehancer.js", "js"))

	//index

	fmt.Println(strings.Index("efe", "u"))

	//türk geliştirici fatihin paketi

	color.Red("sivas58")

	//türk gelişritici paketi log

	loge.Info("Uygulama sonu alias ile")

	//import alias problem anlatımı

	utils.WritingString("public private büyük kçük kapsamın çalıştı")
}

func Maps() {
	//make ile map oluşturdun
	var plates = make(map[int]string)

	//değer atadın
	plates[58] = "Sivas"
	plates[34] = "İstanbul"
	plates[28] = "Giresun"
	plates[44] = "Malatya"
	plates[12] = "Bingöl"

	fmt.Println(plates)

	///delete fonksiyonu ile maptan veri sildin 1 parametre map 2 parmetre key
	delete(plates, 58)

	//foreach ile map dolanım

	for k, v := range plates {
		fmt.Printf("key %v and values %v \n", k, v)
	}

	fmt.Println(plates)

	var keys = make([]int, len(plates))

	var i int = 0

	for k := range plates {
		keys[i] = k
		i++
	}

	fmt.Println(keys)

	for _, v := range keys {
		fmt.Printf("values of map %v \n", plates[v])
	}

	var eleman, err = plates[49]
	fmt.Println(eleman, err)
}

func Slices() {

	var myArray = [5]int{1, 2, 3, 4, 5}
	var mySlice []int = myArray[1:4]
	fmt.Println(mySlice)

	/// sliceler [a:b] a dan a dahil başlar b ye kadar b dahil olmamak üzere alır
	// sliceler arraylardan kopyalar len olarak eleman sayısını alırlar ama
	///sliceler cap() kapasite olarak [a:b] a ile b arasını değil a dan başlayıp arrayin sonuna kadar kapasite alır

	var keyWordArray = [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "ı", "j"}
	var sliceKeyword []string = keyWordArray[2:5]
	fmt.Println(sliceKeyword)
	fmt.Printf("sliceKeyword length %v ", len(sliceKeyword))
	fmt.Printf("sliceKeyword of capacity %v", cap(sliceKeyword))
	///c d e len=3 cap 8

	///sliceler veriyi kopyalamaz bellekten adresini tutarlar yine arrayden çekip işlem yaparlar

	///slice in bir verisini değiştirirsen arrayinde verisi değişir pointer yoluyla çalıştıkları içindir

	var myArray2 = [3]int{8, 16, 24}
	var mySlice2 []int = myArray2[2:3]
	fmt.Println(mySlice2)
	mySlice2[0] = 58
	fmt.Println(myArray2, mySlice2)
	///tanımlanım ve değer değişimi ... operatör ile dizinin elemanlarını sayıp derleyici atayack
	var myArray3 = [...]int{99, 100, 101}
	var mySlice3 []int = myArray3[:]
	fmt.Println("slice", mySlice3)
	fmt.Println("array", myArray3)
	mySlice3[0] = 28
	fmt.Println("slice", mySlice3)
	fmt.Println("array", myArray3)

	///append() function kullanımı

	var mySlice4 = []string{"red", "green", "blue"}
	fmt.Println("myslice4 ", mySlice4)
	mySlice4 = append(mySlice4, "black", "ıdentifier")
	//append ile ilk elemanı silme
	fmt.Println("myslice4 ", mySlice4)
	mySlice4 = append(mySlice4[1:len(mySlice4)])
	fmt.Println("myslice4 ", mySlice4)
	//append ile tek parametre ile son elemanı silme
	mySlice4 = append(mySlice4[:len(mySlice4)-1])
	fmt.Println("myslice4 ", mySlice4)

	///make() function ile kullanım
	/// bana it tipinde 5 elemanlık ve kapasitesi 5 olan slice oluştur
	///ancak eleman ekledikçe derleyici ve go otomatikmen kapasiteyi arttırır
	// kapasite ve len karıştrıma
	mySlice5 := make([]int, 5, 5)
	mySlice5[0] = 10
	fmt.Printf("array %v len %v capaticy %v \n", mySlice5, len(mySlice5), cap(mySlice5))
	mySlice5[1] = 20
	fmt.Printf("array %v len %v capaticy %v \n", mySlice5, len(mySlice5), cap(mySlice5))
	mySlice5[2] = 30
	fmt.Printf("array %v len %v capaticy %v \n", mySlice5, len(mySlice5), cap(mySlice5))
	mySlice5[3] = 40
	fmt.Printf("array %v len %v capaticy %v \n", mySlice5, len(mySlice5), cap(mySlice5))
	mySlice5[4] = 50
	fmt.Printf("array %v len %v capaticy %v \n", mySlice5, len(mySlice5), cap(mySlice5))
	mySlice5 = append(mySlice5, 58)
	fmt.Printf("array %v len %v capaticy %v \n", mySlice5, len(mySlice5), cap(mySlice5))

}

func Arrays() {

	///array tanım atama yazdırım
	var myArray1 = [3]int{}
	myArray1[0] = 23
	myArray1[1] = 28
	myArray1[2] = 58
	fmt.Println(myArray1)

	//ilk elemanı eksik tanımlayıp sonradan set etme testi
	var myArray2 = [3]int{25, 23}
	myArray2[2] = 44
	fmt.Println(myArray2)

	///foerach ile dolanma
	var colors [3]string
	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "blue"
	fmt.Println(colors)
	for k, v := range colors {
		fmt.Printf("key : %v value : %v \n", k, v)
	}

	///len fonksiyonu ile dizi içindeki eleman sayısnı bulma
	///array oluşturup her indexi kullanıcıdan alıp uzunluk belirleyip diziyi bastırma
	var myArray3 [5]string
	for k := range myArray3 {
		var reader = bufio.NewReader(os.Stdin)
		fmt.Printf("%v indexin karşılığını giriniz \n", k)
		value, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("okuma hatası program sonlandırılıyor")
		} else {
			myArray3[k] = strings.TrimSpace(value)
			fmt.Println("girilen değer atandı")
		}
	}
	fmt.Println(myArray3)
	fmt.Println("myArray3 arrays length", len(myArray3))
}

func Arrays2() {

	// ... operatör ile eleman syısı vermeden tanımlama derleyici kendi sayar

	var myArray = [...]int{10, 200, 2, 75, 4, 57, 8, 8, 45, 5, 45, 49, 118}
	fmt.Println(myArray)
	fmt.Println("length ", len(myArray))

	///cap kapasite fonskiyonu

	var myArray2 [3]string
	myArray2[0] = "scania"
	myArray2[1] = "bmc"
	fmt.Println(len(myArray2))
	fmt.Println(cap(myArray2))

}

func errorTest() (int, error) {
	// fmt.Errorf() ile geriye isteilen veride error dönderilebilir
	var i int = 5
	if i < 0 {
		return 0, fmt.Errorf("Bu bir hata")
	}
	return 1, nil
}

/// her dakika if err!=nil yazmak yerine kontrolcü fonksiyon geliştirmesi

func checkError(err error) {
	if err != nil {
		log.Fatal("Error ", err.Error())
		// fmt.Println("ERROR :", err.Error())
		//programı exit ile kapatırsın 1 true değer görünce kapar
		os.Exit(1)
	}
}

func ErrorStatement() {

	file, error := os.Open("demo.txt")

	//hata var ise
	if error != nil {
		fmt.Println(error.Error())
	} else {
		fmt.Println(file.Name())
	}

	//kendi error hatanı oluştur

	var myError = errors.New("Bu bir hatadır")
	fmt.Println(myError.Error())

}

func Loops() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	///koşulsuz döngü tanımlayıp sonsuz çalıştırma

	// for i := 0; ; i++ {
	// 	fmt.Println(i)
	// }

	///; elleme bu bir kalıptır

	// for i := 0; i < 3;i { yine sonsuz çalışır
	// 	fmt.Println(i)
	// }

	// for yapısında while kullanımı

	// var deger int = 7

	// for deger < 18 {
	// 	fmt.Println(deger)
	// 	deger++
	// }

	///Range kullanımı

	var pow = []int{2, 4, 6, 8, 10, 16}

	//range den value değişkeni atayarak elde etme
	for i, v := range pow {
		fmt.Printf("%v İndisli elemanın karşılığı %v\n", i, v)
	}

	var length = [...]string{"%", "&", "/", "{}", "|}"}

	for i := range length {
		fmt.Println(i, length[i])
	}

	var countryies = map[string]string{"Turkey": "Ankara", "France": "paris", "England": "Tokyo"}

	for k := range countryies {
		fmt.Println(countryies[k])
	}

	for k, v := range countryies {
		fmt.Printf("key: %v value: %v\n", k, v)
	}

}

func StatementBreakContiuneGoto() {
	///break

	// bulunduğu scope kırar ordan çıkarır programı
	var i int = 0
	for {
		if i >= 3 {
			break
		}
		fmt.Println("i values of :", i)
		i++
	}
	fmt.Println("Break döngüyü bulunduğu scope`ı kırdı ve çıktı")
}

func SwitchStatementControl() {

	var a int = 13
	///case sağlasın sağlamasın fallthrough olursa alltaki case de çalıştırılır
	switch a {
	case 9:
		fmt.Println("numeric ==9")
		// fallthrough
	case 10:
		fmt.Println("numeric ==10")
	default:
		fmt.Println("not declared types")
	}

	var score int
	fmt.Println("Please log a score ")
	fmt.Scanf("%v", &score)
	fmt.Printf("%v", score)
	switch {
	case score <= 59:
		fmt.Println("Your Notes F")
	case score <= 69:
		fmt.Println("Your Notes C")
	case score <= 79:
		fmt.Println("Your Notes D")
	case score <= 89:
		fmt.Println("Your Notes B")
	case score <= 100:
		fmt.Println("Your notes A")
	default:
		fmt.Println("Not defined value")

	}

}

func İfStateMentControl() {

	var (
		a   int = 10
		b   int = 10
		foo int = 1
	)

	if a > b {
		fmt.Println("a>b")
	} else if a == b {
		fmt.Println("a=b")
	} else {
		fmt.Println("b>a")
	}

	if foo == 1 {
		fmt.Println("bar")
	}

	//if içinde initilaze kurulum yapıp sorgulama

	if efe := 9; efe == 3 {
		fmt.Println("efe ==3")
	} else {
		fmt.Println("efe!=3")
	}
	//   fmt.Println(efe) hata verir o değişken if içinde tanımlandı scope kapsam hatası
}

func DateAndTimePackages() {

	///istenilen tarihi time tipinde oluşturma
	t := time.Date(2004, time.September, 6, 22, 18, 50, 12, time.UTC)
	fmt.Printf("%T", t)
	fmt.Printf("%s", t)

	///şu anki saat tarih ve zamanı now methodu ile çekme
	//var ile değişken tipi bildirimi ile time tanımlama
	var now time.Time = time.Now()
	fmt.Println(now)

	// time nesnesi içinden istenilen verileri çekme

	fmt.Println(t.Day())
	fmt.Println(t.Weekday())
	fmt.Println(t.Month())
	fmt.Println(t.Year())
	fmt.Println(t.Location())

	///tarihi değiştirip yeni tarih dönderme

	var newTime time.Time = t.AddDate(1, 1, 1)
	fmt.Println(newTime)

	//formatlı gösterim
	// senin tarihini istediğin format şeklinde layouıt alarak geri döndürür
	var longFormat string = "Monday , september 6 , 2004"
	fmt.Println(newTime.Format(longFormat))

	var shortFormat string = "1/2/06"
	fmt.Println(t.Format(shortFormat))
}

func ConsoleInputs() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	fmt.Println("Your Values ", str)

	fmt.Print("Enter number")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("your values %v", f)
	}

}

func ConsoleOutputs() {
	var (
		text              string = "Necati Efe Hancer"
		stringLength, err        = fmt.Println(text)
		aNumber           int    = 28
		isBool            bool   = true
	)
	fmt.Println(stringLength, err)
	fmt.Printf("Value of number %v \n", aNumber)
	fmt.Printf("Value of boolean %v\n", isBool)
	fmt.Printf("Value of converted float %.2f\n", float64(aNumber))
	fmt.Printf("Data Types %T %T %T %T %T %T", text, stringLength, err, aNumber, isBool, 15.8)
}

func convertionVariables() {

	var (
		deger1 string  = "8"
		deger2 int     = 10
		deger3 float64 = 2.8
	)

	convertedNumber1, err := strconv.Atoi(deger1)

	fmt.Println(deger1, deger2, deger3)

	fmt.Println(convertedNumber1, err)

	var convertedNumber2 string = strconv.Itoa(deger2)

	fmt.Println(convertedNumber2 + "sivas")

	var i int = 80
	var j float64 = float64(i)
	var k string = "75"
	fmt.Println(i, j, k)
}

func ConstantState() {
	PrintLanguage(JAVASCRİPT)
	PrintLanguage(JAVA)
	PrintLanguage(PYTHON)
	PrintLanguage(C)
}

func BlcakIdentifier() {
	var _, x, _ int = 0, 9, 8
	fmt.Println(x)
}

func letterUpperKeywords() {

	fmt.Println("hello planet")
	fmt.Println("www.heraSystem.com")

	// public private kapsamı kapısına çıkan isimlendirmeler

	var projectName string = "GoBackendServer" //dış dosyalardan buna erişilmesin private yapısı
	var ProjectName string = "GOBackendServer" // dış dosyalardan buna erişilsin public yapısı

	fmt.Println(projectName, ProjectName)
}

func areaVariables() {

	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	var uName string = os.Getenv("USERNAME")
	var uDomain string = os.Getenv("USERDOMAIN")
	var processorArthitecture string = os.Getenv("PROCESSOR_ARCHITECTURE")
	var processorIdentifier string = os.Getenv("PROCESSOR_IDENTIFIER")
	var processorLevel string = os.Getenv("PROCESSOR_LEVEL")
	var goRoot string = os.Getenv("GOROOT")
	var goPath string = os.Getenv("GOPATH")
	var homePath string = os.Getenv("HOMEPATH")

	fmt.Println("username" + uName)
	fmt.Println("Domain" + uDomain)
	fmt.Println("İşlemci Mimarisi " + processorArthitecture)
	fmt.Println(processorIdentifier)
	fmt.Println("İşlemci seviyesi " + processorLevel)
	fmt.Println("GO ROOT ALANI" + goRoot)
	fmt.Println("GO PATH ALANI" + goPath)
	fmt.Println("HOME PATH ALANI" + homePath)
}

func Variables() {
	// const text string = "efe"

	// var (
	// 	degisken1 string = "efe"
	// 	degisken2 int    = 23
	// )

	// var name string
	// name = "necatiefe"
	// fmt.Println(name)

	// var a, b, c int
	// var a, b, c int = 3, 4, 5

	// fmt.Println(a, b, c)

	// var message = "HelloPlanet"
	// var a, b, c = 3, true, 4.8
	// fmt.Println(message, a, b, c)

	// var k, o string = "abc", "xyz"

	// var p = 42
	// var s, b = "xyz", true

	// n := 28
	// name, age := "necati efe", 20
	// fmt.Println(name, age)

	// a := "Text"
	// b := 'T'
	// c := `Text`

	// fmt.Println(a, b, c)

	// var myFloat32 float32 = 58.
	// myComplex := complex(5, 6)
	// println(myFloat32)
	// println(myComplex)

}

func Operators() {
	var (
		a     int = 10
		b     int = 20
		total int
	)
	total = a + b
	total = total - 5

	total *= 20

	total++

	total--
	fmt.Println(total)
}
