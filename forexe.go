package main

import (
	"fmt"
)

func main() {
	var sum int
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			sum = sum + i
		} else if i%2 != 0 {
			sum = sum - i
		}

	}
	fmt.Println(sum)
}
