package intermediate

import "fmt"

/*
RunIntermediate : entry for intermediate
*/
func RunIntermediate() {
	fmt.Println("***** INTERMEDIATE ******")

	fmt.Println("--- testPointers ---")
	testPointers()

	fmt.Println("--- testStruct ---")
	testStruct()

	fmt.Println("--- testArray ---")
	testArray()

	fmt.Println("--- testSlices ---")
	testSlices()
}

func testPointers() {
	i, j := 1, 2
	iPtr := &i
	jPtr := &j
	fmt.Println(&i, &j)
	fmt.Println(*iPtr)
	*jPtr = 10
	fmt.Println(j)

}

type car struct {
	name string
	year int
}

func testStruct() {
	c := car{"bmw", 2020}
	fmt.Println(c.name, c.year)

	// pointer to struct
	c1 := &c
	c1.year = 2025 // same as using *c1.year
	fmt.Println(*c1)

	c2 := car{name: "mercedez"}
	fmt.Println(c2)
}

func testArray() {
	arr := [2]int{1, 2}
	fmt.Println(arr)

	var arr2 [2]string
	arr2[0], arr2[1] = "elt1", "elt2"

	fmt.Println(arr2)

}

func testSlices() {
	// slices
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4] // 3,5,7
	fmt.Println(s)
	s[0] = 10 // 3 -> 10  // slices are references backed by original array
	fmt.Println(primes)

	// slice literals
	points := []struct {
		x int
		y int
	}{
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
		{9, 10},
	}
	points = points[3:]
	fmt.Println(points)
	fmt.Println("len=", len(points), "capacity=", cap(points))

	// nil slices
	var ns []int
	fmt.Println(ns, len(ns), cap(ns))
	if ns == nil {
		fmt.Println("nil!")
	}

	// slice with make
	dSlice := make([]int, 5) // cap = 5, len = 5
	fmt.Println(dSlice, cap(dSlice), len(dSlice))
	dSlice2 := make([]string, 0, 3) // cap = 3, len = 0
	fmt.Println(dSlice2, cap(dSlice2), len(dSlice2))

}
