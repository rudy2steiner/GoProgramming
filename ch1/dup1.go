package main

import (
	"fmt"
	"os"
	"bufio"
	)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// ctrl+D 结束信号
	for input.Scan(){
		counts[input.Text()]++
	}
	fmt.Println("Dup states:")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t %s \n", n, line)
		}
	}
}
