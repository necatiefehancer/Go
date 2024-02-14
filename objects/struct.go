package objects

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

///go da def constructor yoktur geliştiriciler gereksiz görüp elle yazılabilir idda etmiştir

func NewPerson() *Person {
	x := new(Person)
	return x
}

///parameter alan constructor oluşturumu

func NewPerson2(Firstname, Lastname string, Age int) *Person {
	x := new(Person)
	x.Firstname = Firstname
	x.Lastname = Lastname
	x.Age = Age
	return x
}
func Demo1() {

	//	var Efe = Person{Firstname: "necatiEfe"}

	//	fmt.Println(Efe.Firstname)

	// var Efe = new(Person)
	// Efe.Age = 23
	// fmt.Println(Efe.Age)

	x := NewPerson()
	x.Firstname = "BoşConstruc"
	fmt.Println(x.Firstname)

	var x2 Person = *NewPerson2("Adnan", "Hançer", 59)
	fmt.Println(x2)

	var struct1 = Person{Firstname: "Necati Efe"}
	fmt.Println(struct1.Firstname)
	pointer1 := &struct1
	pointer1.Firstname = "Tacettin Tufan"
	fmt.Println(struct1.Firstname)

}
