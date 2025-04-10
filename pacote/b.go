package pacote

import (
	"fmt"
	"myFirstGoProject/pacote/internal/foo"
)

var Bar string = "Hello, Bar!"

func PrintMine() {
	fmt.Println(foo.Mine)
}