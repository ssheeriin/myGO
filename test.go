package main

import "fmt"

var flag1, flag2 bool
var i, j int = 1, 2

func main() {
	fmt.Printf("Hello %s \n", "sherin")

	sum, dif := testNamedReturn(10, 5)
	fmt.Printf("Sum of %d and %d is %d and diff is %d \n", 10, 5, sum, dif)

	testVariables()

	testShortVarDeclarations()
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
