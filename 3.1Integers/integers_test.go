package integers

import (
	"fmt"
)

func ExampleBinaryOperators() {
	x := 2
	x += 2

	fmt.Print(x)

	//Output: 4
}

func ExampleOverflow() {
	var u uint8 = 255
	fmt.Print(u, u+1, u*u)

	//Output: 255 0 1
}
