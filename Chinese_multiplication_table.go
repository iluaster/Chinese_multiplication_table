package main

import "fmt"

func main() {
	for i := 1; i <= 9; i = i + 3 {
		for j := 1; j <= 9; j++ {
			fmt.Printf("%v * %v = %v \t", i, j, i*j)
			fmt.Printf("%v * %v = %v \t", i+1, j, (i+1)*j)
			fmt.Printf("%v * %v = %v\n", i+2, j, (i+2)*j)
		}
		fmt.Printf("\n")
	}
}
