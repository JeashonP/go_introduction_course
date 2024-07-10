package main
 
import "fmt"
 
func main() {
 
    license := true
    age := 15

	if license == true && age == 19 {
		fmt.Println("Puede seguir avanzando")
	} else if license == false && age == 19 {
		fmt.Println("No puede seguir circulando")
	} else if license == true && age == 15 { 
		fmt.Println("No puede seguir circulando 2")
	}
    
}