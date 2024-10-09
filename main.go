package main

import (
	"fmt"

	"github.com/buptgarden/design-pattern-go/abstractfactory"
	"github.com/buptgarden/design-pattern-go/builder"
	"github.com/buptgarden/design-pattern-go/factorymethod"
)

func main() {

	ak47, _ := factorymethod.GetGun("ak47")
	musket, _ := factorymethod.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)

	adidasFactory, _ := abstractfactory.GetSportsFactory("adidas")
	nikeFactor, _ := abstractfactory.GetSportsFactory("nike")

	printShoeDetails(adidasFactory.MakeShoe())
	printShirtDetails(adidasFactory.MakeShirt())

	printShoeDetails(nikeFactor.MakeShoe())
	printShirtDetails(nikeFactor.MakeShirt())

	normalBuilder := builder.GetBuilder("normal")
	iglooBuilder := builder.GetBuilder("igloo")

	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	printHouseDetails(normalHouse)

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()

	printHouseDetails(iglooHouse)

}

func printDetails(g factorymethod.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Gun Power: %d", g.GetPower())
	fmt.Println()
}

func printShoeDetails(s abstractfactory.IShoe) {
	fmt.Printf("Shoe Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Shoe size: %d", s.GetSize())
	fmt.Println()
}

func printShirtDetails(s abstractfactory.IShirt) {
	fmt.Printf("Shirt Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Shirt size: %d", s.GetSize())
	fmt.Println()
}

func printHouseDetails(h builder.House) {
	fmt.Printf("House window type %s \n", h.WindowType)
	fmt.Printf("House door type %s \n", h.DoorType)
	fmt.Printf("House floor size %d \n", h.Floor)
}
