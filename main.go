package main

import (
	"fmt"
)

func main() {
	if calculateAge() {
		var band = true
		for band {
			band = menu()
		}
	} else {
		fmt.Println("You don't have the age required")
	}
}

type Wine struct {
	name  string
	price int
}

var inventory = []Wine{
	{name: "Apple Wine", price: 50000},
	{name: "Blackberry Wine", price: 100000},
}

func calculateAge() bool {
	fmt.Println("Enter your age")
	var age int
	fmt.Scanln(&age)
	if age >= 18 {
		return true
	}
	return false
}

func menu() bool {
	fmt.Println("Welcom to vineria sofka\n" +
		"1.Show inventory\n" + "2.Make an order\n" + "3.Exit")
	fmt.Println("Enter an option (just number)")
	var option int
	fmt.Scanln(&option)
	switch option {
	case 1:
		showInventory()
		break
	case 2:
		order()
		break
	case 3:
		return false
		break
	default:
		fmt.Println("Incorrect option")
	}
	return true
}

func showInventory() {
	for _, wine := range inventory {
		fmt.Println(wine.name, ", price:", wine.price)
	}
}

func order() {
	var unidad int
	var total int
	for _, wine := range inventory {
		fmt.Println(wine.name, ", price:", wine.price)
		fmt.Println("¿Cuantos va llevar?")
		fmt.Scanln(&unidad)
		total = unidad * wine.price
	}
	cobrar(total)
}

func cobrar(total int) {
	var pago int
	fmt.Println("Total to pay", total)
	fmt.Println("Enter your pay")
	fmt.Scanln(&pago)
	if pago > total {
		fmt.Println("Thanks for the tip")
	} else if pago == total {
		fmt.Println("Come back soom")
	} else {
		fmt.Println("lacks money:", (pago - total))
	}

}
