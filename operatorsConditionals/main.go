package main

import (
	"fmt"
)

func main() {
	//los operadores siempre nos devuelven tru or false
	yearsOld := 32
	fmt.Println("Operators")
	fmt.Println(yearsOld > 30)
	fmt.Println(yearsOld < 32)
	fmt.Println(yearsOld <= 32)
	fmt.Println(yearsOld >= 40)
	fmt.Println(yearsOld == 32)

	fmt.Println()
	//operador logico OR
	/**
	A       B       A or B
	false   false   false
	true    false   true
	false   true    true
	true    true    true
	*/
	fmt.Println(yearsOld < 32 || yearsOld == 32)
	fmt.Println(yearsOld < 32 || yearsOld == 33)
	fmt.Println(yearsOld < 40 || yearsOld == 33)

	fmt.Println()
	//operador logico AND
	/**
	A       B       A and B
	false   false   false
	true    false   false
	false   true    false
	true    true    true
	*/
	fmt.Println(yearsOld < 32 && yearsOld == 32)
	fmt.Println(yearsOld < 32 && yearsOld == 33)
	fmt.Println(yearsOld < 40 && yearsOld == 32)

	fmt.Println()
	//operador logico no
	fmt.Println(true)
	fmt.Println(!true)

	fmt.Println(yearsOld < 40)
	fmt.Println(!(yearsOld < 40))

	//agrupar los diferentes operadores
	fmt.Println(yearsOld < 25 && yearsOld == 32 || yearsOld < 40)
	fmt.Println(yearsOld < 25 && (yearsOld == 32 || yearsOld < 40))

	//condicional if
	fmt.Println()
	yearsOldNew := 20

	if yearsOldNew > 18 {
		fmt.Printf("%d is higher than 18\n", yearsOldNew)
	}

	boolVal := true

	if boolVal {
		fmt.Println("Is true")
	} else {
		fmt.Println("Is false")
	}

	if value := true; value {
		fmt.Println("values is: ", value)
	}

	number := 3

	if number == 1 {
		fmt.Println("one")
	} else if number == 2 {
		fmt.Println("two")
	} else if number == 3 {
		fmt.Println("three")
	}

	switch number {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("undefined number")
	}

	switch numberNew := 1; numberNew {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("undefined number")
	}

	switch {
	case number > 0:
		fmt.Println("positive")
	case number < 0:
		fmt.Println("negative")
	case number == 0:
		fmt.Println("Zero")
	}
}