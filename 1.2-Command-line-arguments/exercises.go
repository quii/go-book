package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	exercise1()
	exercise2()
}

func exercise1() {
	fmt.Println("Exercise 1.1, Modify the echo program to also print os.Args[0]")
	fmt.Println(strings.Join(os.Args, " "))
}

func exercise2() {
	fmt.Println("Exercise 1.2 Modify the echo program to print the index and value of each of its arguments, one per line")
	for index, value := range os.Args {
		fmt.Println(index, value)
	}
}
