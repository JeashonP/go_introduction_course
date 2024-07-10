package main

import "fmt"

func main() {

	//creamos e instanciamos nuestro map
	m1 := make(map[int]string)

	//agregamos valores
	m1[1] = "a"
	m1[2] = "b"
	m1[3] = "c"

	fmt.Println(m1)
	fmt.Println(m1[1])

	//para eliminar un dato del map
	delete(m1,2)
	fmt.Println(m1)

	//modificar un valor
	m1[1] = "A2"
	fmt.Println(m1)

	m1[7] = ""
	fmt.Println(m1)
	fmt.Println(m1[7])

	v, ok := m1[7]
	fmt.Println("Thw value: ", v, "Present?", ok)

	_, ok = m1[99]
	fmt.Println("Thw value: ", v, "Present?", ok)

	//otra forma de instanciar los mapas 
	m2 := map[int]string{
		4: "A",
		5: "B",
		6: "C",
	}
	fmt.Println(m2)
	
}