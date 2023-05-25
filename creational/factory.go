package creational

import "log"

type Drinkable interface {
	Drink() string
}

type MilkShake struct{}

func NewMilkShake() *MilkShake {
	return &MilkShake{}
}

func (c *MilkShake) Drink() string {
	return "Strawberry Milkshake"
}

type Beer struct{}

func NewBeer() *Beer {
	return &Beer{}
}

func (d *Beer) Drink() string {
	return "Guinness"
}

func Bar(personType string) {
	var beverage Drinkable

	if personType == "alcoholic" {
		beverage = NewBeer()
	} else {
		beverage = NewMilkShake()
	}

	log.Panicln(beverage.Drink())
}
