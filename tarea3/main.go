package main
 
import "fmt"
 
func main() {

	var array []int

	println("Ingrese valores (ingrese 0 para finalizar):")
    
    for {
		var value int

		if _, err := fmt.Scanf("%d", &value); err != nil {
			continue
		}

		if value == 0 {
			break
		}

		array = append(array, value)
	}
 
    fmt.Println("Los valores del array son: ", array)
    
}