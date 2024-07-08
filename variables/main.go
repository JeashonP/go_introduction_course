package main

import "fmt"

func main() {

	//VARIABLES
	var myIntVar int
	myIntVar = -12
	fmt.Println("My variable myIntVar is:", myIntVar)

	var myUintVar uint
	myUintVar = 12
	fmt.Println("My variable myUintVar is:", myUintVar)

	var myStringVar string
	myStringVar = "Json"
	fmt.Println("My variable myStringVar is:", myStringVar)

	var myBooleanVar bool
	myBooleanVar = true
	fmt.Println("My variable myBooleanVar is:", myBooleanVar)

	fmt.Println("My variable address is:", &myStringVar)

	//Nota: podemos setearvariables sin haberlas asignado
	myIntVar2 := 12
	fmt.Println("My variable myIntVar2 is:", myIntVar2)

	myStringVar2 := "Pinzón"
	fmt.Println("My variable myStringVar2 is:", myStringVar2)

	fmt.Println()

	//CONSTANTES
	const myFirstConst = "a12"
	fmt.Println("My const myFirstConst is:", myFirstConst)

	const myIntConst = 12
	fmt.Println("My const myIntConst is:", myIntConst)

	const myStringConst = "a"
	fmt.Println("My const myStringConst is:", myStringConst)

}
