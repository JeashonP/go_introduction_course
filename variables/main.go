package main

import (
	"fmt"
	"unsafe"
	"strconv"
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
	    Type    Sizq      Range
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

	//UINT
	/*
	    Type    Sizq      Range
		uint8    8 bits    0 to 255
		uint16   16 bits   0 to 2^16 -1
		uint32   32 bits   0 to 2^32 -1
		uint64   64 bits   0 to 2^64 -1
		uint Platform dependent
	*/

	fmt.Println()

	var my8BitsUintVar uint8
	fmt.Printf("Type: %T, bytes %d, bits %d\n", my8BitsUintVar, unsafe.Sizeof(my8BitsUintVar), unsafe.Sizeof(my8BitsUintVar)*8)

	var my16BitsUintVar uint16
	fmt.Printf("Type: %T, bytes %d, bits %d\n", my16BitsUintVar, unsafe.Sizeof(my16BitsUintVar), unsafe.Sizeof(my16BitsUintVar)*8)

	var my32BitsUintVar uint32
	fmt.Printf("Type: %T, bytes %d, bits %d\n", my32BitsUintVar, unsafe.Sizeof(my32BitsUintVar), unsafe.Sizeof(my32BitsUintVar)*8)

	var my64BitsUintVar uint64
	fmt.Printf("Type: %T, bytes %d, bits %d\n", my64BitsUintVar, unsafe.Sizeof(my64BitsUintVar), unsafe.Sizeof(my64BitsUintVar)*8)

	fmt.Println()

	//TIPOS DE DATOS FLOTANTES
	var myFloat32Var float32
	fmt.Printf("Float default value: %f\n", myFloat32Var)
	fmt.Printf("Type: %T, bytes %d, bits %d\n", myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)

	var myFloat64Var float64
	fmt.Printf("Float default value: %f\n", myFloat64Var)
	fmt.Printf("Type: %T, bytes %d, bits %d\n", myFloat64Var, unsafe.Sizeof(myFloat64Var), unsafe.Sizeof(myFloat64Var)*8)

	fmt.Println()

	var myStringVar3 string
	fmt.Printf("String default value: %s\n", myStringVar3)

	myStringVar4 := `My string variable in golang
	whit multiple
	line` 	
	fmt.Printf("The variable myStringVar4 value is: %s\n", myStringVar4)

	//scope
	{
		fmt.Println()

		floatVar := 33.11
		fmt.Printf("Type: %T, value %f\n", floatVar, floatVar)

		//convertir a string
		floatStrVar := fmt.Sprintf("%.2f", floatVar)
		fmt.Printf("Type: %T, value %s\n", floatStrVar, floatStrVar)

		intVar := 22
		fmt.Printf("Type: %T, value %d\n", intVar, intVar)

		//convertir a string
		intStrVar := fmt.Sprintf("%d", intVar)
		fmt.Printf("Type: %T, value %s\n", intStrVar, intStrVar)

		intVal1, err := strconv.ParseInt("1333", 0, 64)
		fmt.Println(err)
		fmt.Printf("Type: %T, value %d\n", intVal1, intVal1)

		intVal2, err := strconv.ParseInt("aa122", 0, 64)
		fmt.Println(err)
		fmt.Printf("Type: %T, value %d\n", intVal2, intVal2)

		//tomamos el error pero no hacemos nada con el
		floatVal1, err := strconv.ParseFloat("-11.2", 64)
		fmt.Printf("Type: %T, value %f\n", floatVal1, floatVal1)
	}

}
