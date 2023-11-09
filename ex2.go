package main

import "fmt"

func main() {
	var mark int = 100
	if mark >= 90 && mark <= 100 {
		fmt.Println("A")
	} else if mark >= 80 && mark <= 89 {
		fmt.Println("B")

	} else if mark >= 70 && mark <= 79 {
		fmt.Println("C")

	} else if mark >= 60 && mark <= 69 {
		fmt.Println("D")

	} else if mark >= 0 && mark <= 59 {
		fmt.Println("F")

	}
}
