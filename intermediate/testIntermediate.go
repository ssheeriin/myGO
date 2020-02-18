package intermediate

import "fmt"

/*
Intermediate : entry for intermediate
*/
func Intermediate() {

	testPointers()
}

func testPointers() {
	i, j := 1, 2

	fmt.Println(&i, &j)

}
