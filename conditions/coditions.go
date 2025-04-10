package conditions

import (
	"errors"
	"fmt"
	"time"
)

func Conditions() {
	if x := 2; 1 < 2 {
		fmt.Println("True")
	} else if x > 0 {
		fmt.Println("Greater than 0")
	}

	if error := doError(); error != nil {
		fmt.Println("Different than nil")
	}

	do(1)
	do(2)
	do(3)

	fmt.Println(isWeekend(time.Now()))
}

func doError() error {
	return errors.New("error")
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