package main

import (
	"os"
	"fmt"
	tempcov "github.com/quii/go-book/2.1TempConv"
	"strconv"
)

func main(){
	for _, arg := range os.Args[1:]{
		t, err := strconv.ParseFloat(arg, 64)

		if err != nil{
			fmt.Println("IUHFIUEHFiuh")
			os.Exit(1)
		}

		c := tempcov.Celsius(t)
		fmt.Println(c)
		fmt.Println(tempcov.CToF(c))
	}



}
