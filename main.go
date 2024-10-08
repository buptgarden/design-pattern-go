package main

import (
	"fmt"

	"github.com/buptgarden/design-pattern-go/abstractfactory"
	"github.com/buptgarden/design-pattern-go/adapter"
	"github.com/buptgarden/design-pattern-go/bridge.go"
	"github.com/buptgarden/design-pattern-go/builder"
	"github.com/buptgarden/design-pattern-go/factorymethod"
	"github.com/buptgarden/design-pattern-go/prototype"
	"github.com/buptgarden/design-pattern-go/singleton"
)

func main() {
	factoryMethod()

	abstractFactory()

	builderChapter()

	prototypeChapter()

	singletonChapter()

	adapterChapter()

	bridgeChapter()

}

// chapt worker

func factoryMethod() {
	ak47, _ := factorymethod.GetGun("ak47")
	musket, _ := factorymethod.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func abstractFactory() {

	adidasFactory, _ := abstractfactory.GetSportsFactory("adidas")
	nikeFactor, _ := abstractfactory.GetSportsFactory("nike")

	printShoeDetails(adidasFactory.MakeShoe())
	printShirtDetails(adidasFactory.MakeShirt())

	printShoeDetails(nikeFactor.MakeShoe())
	printShirtDetails(nikeFactor.MakeShirt())
}

func builderChapter() {
	normalBuilder := builder.GetBuilder("normal")
	iglooBuilder := builder.GetBuilder("igloo")

	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	printHouseDetails(normalHouse)

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()

	printHouseDetails(iglooHouse)
}

func prototypeChapter() {
	file1 := &prototype.File{Name: "file1"}
	file2 := &prototype.File{Name: "file2"}
	file3 := &prototype.File{Name: "file3"}

	folder1 := &prototype.Folder{
		Children: []prototype.Inode{file1},
		Name:     "Folder1",
	}

	folder2 := &prototype.Folder{
		Children: []prototype.Inode{folder1, file2, file3},
		Name:     "Folder2",
	}

	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print("  ")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")
}

func singletonChapter() {
	for i := 0; i < 30; i++ {
		// go singleton.GetInstance()
		go singleton.GetInstance_Once()
	}

	fmt.Scanln()
}

func adapterChapter() {
	client := &adapter.Client{}

	mac := &adapter.Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windows := &adapter.Window{}

	windowsAdapter := &adapter.WindowsAdaper{
		Windows: windows,
	}

	client.InsertLightningConnectorIntoComputer(windowsAdapter)
}

func bridgeChapter() {
	hpPrinter := &bridge.HP{}
	epsonPrinter := &bridge.Epson{}

	macComquter := &bridge.Mac{}
	macComquter.SetPrinter(hpPrinter)
	macComquter.Print()
	fmt.Println()

	macComquter.SetPrinter(epsonPrinter)
	macComquter.Print()
	fmt.Println()

	windowComputer := &bridge.Windows{}

	windowComputer.SetPrinter(hpPrinter)
	windowComputer.Print()
	fmt.Println()

	windowComputer.SetPrinter(epsonPrinter)
	windowComputer.Print()
	fmt.Println()
}

// tool function

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
