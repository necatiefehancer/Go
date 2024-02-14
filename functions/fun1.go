package functions

import (
	"fmt"
	"myModules/data"
)

func Add(x int, y int) (int, error) {

	return x + y, nil

}

func SayHello1() {
	fmt.Println("HELLO")
}

func SayHello2(message string) {
	fmt.Printf("HELLO %v\n", message)
}

func Change1(x int) {
	fmt.Println(x)
	x = 96
}

func Change1ByReference(x *int) {
	fmt.Println(*x)
	*x = 96
}

func Swap(m1, m2 string) (string, string) {
	return m2, m1
}

func VariadicSum(terms ...int) (int, int) {

	///slice alırlar

	var sum int = 0

	for _, term := range terms {
		sum += term
	}

	return len(terms), sum

}

func NamedVariableFunctions(sum int) (x, y int) {
	x = sum * sum
	y = sum * sum * sum
	return
}

func NamedVariableFunctionsVariadic(terms ...int) (sum int, lenterms int) {
	for _, v := range terms {
		sum += v
	}
	lenterms = len(terms)
	return
}

func VariadicMessageFunctions(messages ...string) {
	for _, message := range messages {
		fmt.Println(message)
	}
}

func VariadicSumFunctions(terms ...int) (sum int) {
	for _, num := range terms {
		sum += num
	}
	return
}

func AnonimFunctions() {
	var anonim = func(terms ...int) (numTerms int, sum int) {
		for _, v := range terms {
			sum += v
		}
		numTerms = len(terms)
		return
	}

	fmt.Println(anonim(9, 2, 2))
}

func DeferStatement() {
	data.Start()
}
