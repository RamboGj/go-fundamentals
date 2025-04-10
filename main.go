package main

import (
	"fmt"
	"time"
)

func main() {
	do(1)
	do(2)
	do(3)

	fmt.Println(isWeekend(time.Now()))
}

func do(x int) {
	switch x {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default") 
	}
}

	
func isWeekend(x time.Time) bool {
	switch {
	case x.Weekday() > 0 && x.Weekday() < 6:
		return false
	default: 
		return true
	}
}