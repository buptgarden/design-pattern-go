package main

import (
	"fmt"
	"log"

	"github.com/buptgarden/design-pattern-go/abstractfactory"
	"github.com/buptgarden/design-pattern-go/adapter"
	"github.com/buptgarden/design-pattern-go/bridge.go"
	"github.com/buptgarden/design-pattern-go/builder"
	"github.com/buptgarden/design-pattern-go/chain"
	"github.com/buptgarden/design-pattern-go/command"
	"github.com/buptgarden/design-pattern-go/composite"
	"github.com/buptgarden/design-pattern-go/decorator"
	"github.com/buptgarden/design-pattern-go/facade"
	"github.com/buptgarden/design-pattern-go/factorymethod"
	"github.com/buptgarden/design-pattern-go/flyweight"
	"github.com/buptgarden/design-pattern-go/iterator"
	"github.com/buptgarden/design-pattern-go/prototype"
	"github.com/buptgarden/design-pattern-go/proxy"
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

	compositeChapter()

	decoratorChapter()

	facadeChapter()

	flywightAdapter()

	proxyChapter()

	chainChapter()

	commandChapter()

	iteratorChapter()

}

// chapt worker

func iteratorChapter() {
	user1 := &iterator.User{
		Name: "a",
		Age:  12,
	}

	user2 := &iterator.User{
		Name: "b",
		Age:  13,
	}

	userCollection := &iterator.UserCollection{
		Users: []*iterator.User{user1, user2},
	}

	iterator := userCollection.CreateIterator()

	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}

func commandChapter() {
	tv := &command.Tv{}

	onCommand := &command.OnCommand{
		Device: tv,
	}

	offCommand := &command.OffCommand{
		Device: tv,
	}

	onButton := &command.Button{
		Command: onCommand,
	}

	onButton.Press()

	offButtion := &command.Button{
		Command: offCommand,
	}

	offButtion.Press()
}

func chainChapter() {
	cashier := &chain.Cashier{}

	medical := &chain.Medical{}

	medical.SetNext(cashier)

	doctor := &chain.Doctor{}
	doctor.SetNext(medical)

	reception := &chain.Reception{}
	reception.SetNext(doctor)

	patient := &chain.Patient{
		Name: "abc",
	}

	reception.Execute(patient)
}

func proxyChapter() {
	nginxServer := proxy.NewNginxServer()

	appStatusURL := "/app/status"
	createruserURL := "/create/user"

	httpCode, body := nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createruserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createruserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
}
func flywightAdapter() {
	game := flyweight.NewGame()

	game.AddTerrorist(flyweight.TerroristDressType)
	game.AddTerrorist(flyweight.TerroristDressType)
	game.AddTerrorist(flyweight.TerroristDressType)
	game.AddTerrorist(flyweight.TerroristDressType)
	game.AddTerrorist(flyweight.TerroristDressType)

	game.AddCounterTerrorist(flyweight.ConunterDressType)
	game.AddCounterTerrorist(flyweight.ConunterDressType)
	game.AddCounterTerrorist(flyweight.ConunterDressType)
	game.AddCounterTerrorist(flyweight.ConunterDressType)
	game.AddCounterTerrorist(flyweight.ConunterDressType)

	dressFacotryInstanc := flyweight.GetDressFactorSingleInstance()

	for dreeType, dress := range dressFacotryInstanc.DressMap {
		fmt.Printf("DressColorType: %s\n DressColor: %s\n", dreeType, dress.GetColor())
	}

}

func facadeChapter() {
	fmt.Println()
	walletFacade := facade.NewWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.AddMoneyToWallet("abc", 1234, 10)

	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()

	err = walletFacade.DeductMoneyFromWallet("abc", 1234, 5)

	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

}

func decoratorChapter() {
	piaaz := &decorator.VeggieMania{}

	piaazWithCheese := &decorator.CheeseTopping{
		Pizza: piaaz,
	}

	pizzaWithTomato := &decorator.TomatoTopping{
		Pizza: piaazWithCheese,
	}

	fmt.Printf("price of veggeMania with tomato and  cheese topping is %d\n", pizzaWithTomato.GetPrice())

}

func compositeChapter() {
	file1 := &composite.File{
		Name: "file1",
	}
	file2 := &composite.File{Name: "file2"}
	file3 := &composite.File{Name: "file3"}

	folder1 := &composite.Folder{
		Name: "folder1",
	}

	folder1.Add(file1)

	folder2 := &composite.Folder{
		Name: "folder2",
	}

	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("rose")
}

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
