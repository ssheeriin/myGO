package main

import (
	"fmt"
	"runtime"
)

var flag1, flag2 bool
var i, j int = 1, 2

func main() {
	fmt.Printf("Hello %s \n", "sherin")

	sum, dif := testNamedReturn(10, 5)
	fmt.Printf("Sum of %d and %d is %d and diff is %d \n", 10, 5, sum, dif)

	testVariables()
	testShortVarDeclarations()
	testLoops()
	testIfAndElse()
	testSwitch()
}

func testNamedReturn(x, y int) (sum, diff int) {
	sum = x + y
	diff = x - y
	return
}

func testVariables() {
	fmt.Println(flag1, flag2)
	fmt.Println(i, j)

}
func testShortVarDeclarations() {
	k := 120
	fmt.Println(k)
}

func testLoops() {

	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	i := 0
	for true {
		if i == 10 {
			break
		}
		fmt.Print(i)
		i++
	}

	fmt.Println()

	/* for {

	} */ // inifite loop
}

func testIfAndElse() {
	if i := 999; i == 999 {
		fmt.Println(i)
	} else {
		fmt.Println("Hello")
	}

	fmt.Println(i)
}

func testSwitch() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
