package functions

import "fmt"

func Functions() {
	fmt.Println("Hello World")

	sayHi()
	var sum int = somar(10, 20)
	fmt.Println("Sum is:", sum)
	a, b := swap(10, 20)
	fmt.Println(a, b)

	res, rem := divide(13, 5)
	fmt.Println(res, rem)
	
	x := somarClosure(2)(1)
	fmt.Println(x)

	sumVariety := somarVariety(10, 50, 20)
	fmt.Println(sumVariety)

	// Arrays
	arr := [3]int{0: 32, 2: 99}
	fmt.Println(arr)
}

func sayHi() {
	fmt.Println("hello")
}

func somar(a int, b int) int {
	return a + b
}

func swap(a int, b int) (int, int) {
	return b, a
}

func divide(a int, b int) (res int, rem int) {
	res = a / b
	rem = a % b
	return res, rem
}

func somarClosure(a int) func(int) int {
	return func(b int) int {
		return a + b;
	}
}

func somarVariety(nums ...int) int {
	var out int

	for _, n := range nums {
		out += n
	}

	return out
}