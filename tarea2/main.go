package main
 
import "fmt"
 
func main() {

	array1 := [5]int{4, 2, 5, 6, 7}
    
    for i, v := range array1 {
		array1[i] = v + 20
	}
 
    fmt.Println("Los valores del array son: ", array1)
    
}