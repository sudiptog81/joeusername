package joeusername

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var bound int
	if len(os.Args) > 1 {
		bound, _ = strconv.Atoi(os.Args[1])
	}
	fmt.Println(Generate(bound))
}
