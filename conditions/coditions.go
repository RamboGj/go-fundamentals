package conditions

import (
	"errors"
	"fmt"
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
}

func doError() error {
	return errors.New("error")
}