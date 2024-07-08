package main

import (
	"fmt" 
	"unsafe"
)

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

	/*
		int8    8 bits    -128 to 127
		int16   16 bits   -2^15 to 2^15 -1
		int32   32 bits   -2^31 to 2^31 -1
		int64   64 bits   -2^63 to 2^63 -1
		int Platform dependent
	*/

	fmt.Println()

	//imprimir con Printf para dar un formato especifico 
	var my8BitsIntVar int8
	fmt.Printf("Int default value: %d\n", my8BitsIntVar)

	//validamos cuanta memoria tiene instanciada la variable
	fmt.Printf("Type: %T, bytes %d, bits %d\n", my8BitsIntVar, unsafe.Sizeof(my8BitsIntVar), unsafe.Sizeof(my8BitsIntVar)*8)

	var my16BitsIntVar int16
	fmt.Printf("Type: %T, bytes %d, bits %d\n", my16BitsIntVar, unsafe.Sizeof(my16BitsIntVar), unsafe.Sizeof(my16BitsIntVar)*8)

	var my32BitsIntVar int32
	fmt.Printf("Type: %T, bytes %d, bits %d\n", my32BitsIntVar, unsafe.Sizeof(my32BitsIntVar), unsafe.Sizeof(my32BitsIntVar)*8)

	var my64BitsIntVar int64
	fmt.Printf("Type: %T, bytes %d, bits %d\n", my64BitsIntVar, unsafe.Sizeof(my64BitsIntVar), unsafe.Sizeof(my64BitsIntVar)*8)

}
