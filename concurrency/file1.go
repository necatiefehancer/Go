package concurrency

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func Xfunc() {

	for l := byte('a'); l <= byte('z'); l++ {
		fmt.Println(string(l))
	}
}

func Fn1() {
	runtime.GOMAXPROCS(7)
	go Xfunc()
	time.Sleep(time.Hour * 5)
}

func Sum(arr []int, ch chan int) {
	var sum int = 0
	for _, v := range arr {
		sum += v
	}
	ch <- sum
}

func Fn2() {

	var myArray = []int{1, 2, 3, 4, 5, 6}
	var myChan = make(chan int)
	go Sum(myArray, myChan)
	var x int = <-myChan
	fmt.Printf("beklendi ve alınan değer %v", x)

}

func Afunc(mych chan string) {

	for x := byte('a'); x <= byte('z'); x++ {
		mych <- string(x)
	}

}

func Bfunc(myChan chan string) {

	for {
		fmt.Println(<-myChan)
	}

}

func ABMain() {

	var ch1 = make(chan string)

	go Bfunc(ch1)
	go Afunc(ch1)

	fmt.Println("program tamamlandı")

	time.Sleep(100 * time.Millisecond)
}
