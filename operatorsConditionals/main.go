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

}