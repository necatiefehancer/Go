package objects

import (
	"fmt"
	"strconv"
)

// user

type User struct {
	ID        int
	Firstname string
	LastName  string
	UserName  string
	Age       int
	Pay       *Payment
}

type Admin struct {
	id       int
	userName string
	passWord string
	age      int
	ip       int
}

func (A *Admin) Setİd(newİd int) {
	A.id = newİd
}
func (A Admin) Getİd() int {
	return A.id
}

//payment

type Payment struct {
	Salary float64
	Bonus  float64
}

//Payment Constructor

func NewPayment() *Payment {
	var newPayment = new(Payment)
	return newPayment
}

//User Constructor

func NewUser() *User {
	var newUser = new(User)
	newUser.Pay = NewPayment()
	return newUser
}

//METHODS

func (u User) GetFullName() string {
	return u.Firstname + " " + u.LastName
}

func (u User) GetUserName() string {
	return u.UserName
}

func (u User) GetPayment() float64 {
	return u.Pay.Bonus + u.Pay.Salary
}

func Demo2() {

	// kullanıcı veri olusturma alanı 1
	fmt.Println("Kullanıcı oluşturma v1")

	user1 := &User{
		ID:        1,
		Firstname: "Tacettin Tufan",
		LastName:  "Hançer",
		Age:       3,
		UserName:  "tufan58",
		Pay: &Payment{
			Salary: 1000,
			Bonus:  850,
		},
	}

	fmt.Println(user1)

	fmt.Println(user1.GetFullName())
	fmt.Println("Maaş Tutarı " + strconv.FormatFloat(user1.GetPayment(), 'g', -1, 64))

	// kullanıcı oluşturma v2

	var user2 = NewUser()

	user2.Firstname = "Selçuk Poyraz"
	user2.LastName = "Hançer"
	user2.Age = 3
	user2.ID = 5834
	user2.UserName = "poyraz58"
	user2.Pay = NewPayment()
	user2.Pay.Bonus = 100
	user2.Pay.Salary = 900

	fmt.Println(user2)

}
