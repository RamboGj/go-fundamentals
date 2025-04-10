package loopsgo

import "fmt"

func Loops() {
	var res int
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, elem := range arr {
		fmt.Println(elem)
	}

	fmt.Println(res)
}
