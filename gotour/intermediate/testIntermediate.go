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

/**
Notes:
1. array is passed by value
2. to pass by refernce , pass a pointer to array
3. arr := [2]int -> pointer to array is *arr [2]int
4. as shown above, pointer type is exactly [2]int , so have to be explicitly specified.
*/
func testArray() {

	arr := [2]int{1, 2}
	fmt.Println(arr)

	var arr2 [2]string
	arr2[0], arr2[1] = "elt1", "elt2"
	fmt.Println(arr2) // pass by value

	var arr3 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	// pass by reference
	arr4 := [1]string{"hello"}
	_processArray(&arr2)
	//_processArray(&arr4) // error

}
func _processArray(ptr *[2]string) {
	fmt.Println(ptr)
}

/**
Notes:
- built on arrays
- slices are by default pass by reference

https://blog.golang.org/go-slices-usage-and-internals
*/
func testSlices() {
	// A slice literal is declared just like an array literal, except you leave out the element count
	primes := []int{2, 3, 5, 7, 11, 13}
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

	// slice with make([]T, len, cap)
	dSlice := make([]int, 5) // cap = 5, len = 5
	fmt.Println(dSlice, cap(dSlice), len(dSlice))

	dSlice2 := make([]string, 0, 3) // cap = 3, len = 0
	fmt.Println(dSlice2, cap(dSlice2), len(dSlice2))

	// slice from array
	x := [3]string{"a", "b", "c"}
	sl := x[:] // a slice referencing the storage of x
	fmt.Println(sl)

	// append slice
	var sApnd []int
	fmt.Println(sApnd, len(sApnd), cap(sApnd))

	sApnd = append(sApnd, 1, 2, 3, 4, 5)
	fmt.Println(sApnd, len(sApnd), cap(sApnd))

}
