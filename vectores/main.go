package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var myIntVar int
	fmt.Println(myIntVar)
	fmt.Printf("Type: %T, bytes: %d, bits: %d\n", myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar) * 8)

	var myArrayVar [6] int 
	fmt.Println(myArrayVar)
 
	myArrayVar1 := [3] string {"val1", "val2", "val3"}
	fmt.Println(myArrayVar1)

	//vectores
	myArrayVar[1] = 2
	myArrayVar[2] = 5
	myArrayVar[3] = 9
	fmt.Println(myArrayVar)
	fmt.Println("size: ", len(myArrayVar))
	fmt.Printf("Type: %T, bytes: %d, bits: %d\n", myArrayVar, unsafe.Sizeof(myArrayVar), unsafe.Sizeof(myArrayVar) * 8)

	//slice
	fmt.Println()
	var slice1 [] int 
	fmt.Printf("size:  %d, value: %d\v", len(slice1), slice1)

	//alimentamos el slice
	slice1 = append(slice1, 10,20,30,40,50)
	fmt.Printf("size:  %d, value: %d\v", len(slice1), slice1)

	//alimentar un slice extrayecto los valores de un array
	slice := myArrayVar[2:4]
	fmt.Println(slice)
	fmt.Println("size: ", len(slice))

	//conocer la direccion de memoria
	fmt.Println(&myArrayVar[2])
	fmt.Println(&slice[0])

	fmt.Println(myArrayVar)

	slice[0] = 80
	slice[1] = 100

	fmt.Println(myArrayVar)

	//mostras desde el inicio hasta donde se escoja
	fmt.Println(myArrayVar[:4])

	//mostrar desde donde se escoja hasta el final
	fmt.Println(myArrayVar[2:])

	//instanciar un slice
	slice2 := make([]int,3)
	fmt.Println(slice2)
	fmt.Printf("size:  %d, value: %d\v", len(slice2), slice2)
}