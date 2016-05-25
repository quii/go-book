package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const(
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g degrees C", c)
}

func (c Fahrenheit) String() string {
	return fmt.Sprintf("%g degrees F", c)
}

// In go when you declare a type alias on X on Y you get a function X(Y)

// CTof converts Celsius into Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32)}

// FToC converts Fahrenheit into Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9)}