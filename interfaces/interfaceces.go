package interfaces

import "strconv"

func Demo1() {

}

//interfaceces

type Icar interface {
	Run() bool
	Stop() bool
	İnformation() string
}

// Base Struct`s

type Car struct {
	brand string
	model string
	color string
	speed int
	price float64
}

// Car get set Methods

func (c *Car) GetCarBrand() string {
	return c.brand
}

func (c *Car) SetCarBrand(newBrand string) {
	c.brand = newBrand
}

func (c *Car) GetCarModel() string {
	return c.model
}

func (c *Car) SetCarModel(newModel string) {
	c.model = newModel
}

func (c *Car) GetCarColor() string {
	return c.color
}

func (c *Car) SetCarColor(newColor string) {
	c.color = newColor
}

func (c *Car) GetCarSpeed() int {
	return c.speed
}

func (c *Car) SetCarSpeed(newSpeed int) {
	c.speed = newSpeed
}

func (c *Car) GetCarPrice() float64 {
	return c.price
}

func (c *Car) SetCarPrice(newPrice float64) {
	c.price = newPrice
}

type SpecialProduction struct {
	special bool
}

//special get set methods

func (s *SpecialProduction) GetSpecial() bool {
	return s.special
}

func (s *SpecialProduction) SetSpecial(newSpecial bool) {
	s.special = newSpecial
}

//Ferrari

type Ferrari struct {
	Car //composition ferrari üretildiği zaman car dışarı açık olan field erişimi sağlar
	SpecialProduction
}

func (_ Ferrari) Run() bool {
	return true
}

func (_ Ferrari) Stop() bool {
	return true
}

func (f *Ferrari) Info() string {
	var info string = f.GetCarBrand() + f.GetCarModel() + f.GetCarColor() + strconv.Itoa(f.GetCarSpeed()) + strconv.FormatFloat(f.GetCarPrice(), 'g', -1, 64)

	if f.GetSpecial() {
		info += "special production"
	} else {
		info += "unspecial production"
	}
	return info
}

//Lamborghini

type Lamborghini struct {
	Car //composition Lamborghini üretildiği zaman car dışarı açık olan field erişimi sağlar
	SpecialProduction
}

func (_ Lamborghini) Run() bool {
	return true
}

func (_ Lamborghini) Stop() bool {
	return true
}

func (f *Lamborghini) Info() string {
	var info string = f.GetCarBrand() + f.GetCarModel() + f.GetCarColor() + strconv.Itoa(f.GetCarSpeed()) + strconv.FormatFloat(f.GetCarPrice(), 'g', -1, 64)

	if f.GetSpecial() {
		info += "special production"
	} else {
		info += "unspecial production"
	}
	return info
}

//Mercedes

type Mercedes struct {
	Car //composition Mercedes üretildiği zaman car dışarı açık olan field erişimi sağlar
	SpecialProduction
}

func (_ Mercedes) Run() bool {
	return true
}

func (_ Mercedes) Stop() bool {
	return true
}

func (f *Mercedes) Info() string {
	var info string = f.GetCarBrand() + f.GetCarModel() + f.GetCarColor() + strconv.Itoa(f.GetCarSpeed()) + strconv.FormatFloat(f.GetCarPrice(), 'g', -1, 64)

	if f.GetSpecial() {
		info += "special production"
	} else {
		info += "unspecial production"
	}
	return info
}
