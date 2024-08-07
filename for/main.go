package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum++
	}

	fmt.Println(sum)

	for sum < 1000 {
		sum++
	}

	fmt.Println(sum)

	for {
		if sum > 1000 {
			break
		}
		sum++
	}
	fmt.Println(sum)

	//recorriendo arrays
	arr := []int{1, 2, 3, 1, 2, 3}

	fmt.Println()

	for i := range arr {
		fmt.Println("Index: ", i,  " - Value: ", arr[i])
	}

	for _, v := range arr {
		fmt.Println("Value: ", v)
	}

	//recorriendo mapas
	map2 := map[string]float64{
		"A": 12.3,
		"B": 23.1,
		"C": 34,
	}

	fmt.Print()

	for key, value := range map2 {
		fmt.Println("key: ", key, " Value: ", value)
	}

	fmt.Print()

	map3 := map[string][]int{
		"A": nil,
		"B": {2, 34, 1, 2, 4},
		"C": {4, 5, 3, 2, 1},
	}

	for key, value := range map3 {
		fmt.Println("key: ", key)
		for _, v := range value {
			fmt.Println("Value: ", v)
		}
		fmt.Println()
	}
}