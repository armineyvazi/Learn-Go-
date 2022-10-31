package main

import (
	"fmt"
)

const (
	Small = iota
	Medium
	Large
)
const (
	Ground = iota
	Air
)

type BeltSize int
type Shipping int

func (b BeltSize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}

func (s Shipping) String() string {
	return []string{"Ground", "Air"}[s]

}

type Conveyor interface {
	Convey() BeltSize
}

type shipper interface {
	Ship() Shipping
}

type WarehouseAutomator interface {
	Conveyor
	shipper
}

type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return "Spam mail"
}

func (s *SpamMail) Ship() Shipping {
	return Air
}

func (s *SpamMail) Convey() BeltSize {
	return Small
}

func automate(item WarehouseAutomator) {
	fmt.Printf("Convey %v on %v Conveyor\n", item, item.Convey())
	fmt.Printf("ship %v via %v\n", item, item.Ship())

}

type ToxicWaste struct {
	weight int
}

func (t *ToxicWaste) ship() Shipping {
	return Ground
}

func (t *ToxicWaste) Convey() BeltSize {
	return Air
}
func main() {
	mail := SpamMail{4000}
	automate(&mail)

}
