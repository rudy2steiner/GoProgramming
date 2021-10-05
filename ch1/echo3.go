package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Program:" + os.Args[0])
	fmt.Println(strings.Join(os.Args[1:]," "))
}
