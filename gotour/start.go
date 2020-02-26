package main

import "fmt"

// import (
// 	"sherin.dev/gotour/basic"
// 	"sherin.dev/gotour/intermediate"
// )

func main() {
	//basic.RunBasic()
	//intermediate.RunIntermediate()

	scrap()
}

func scrap() {
	arr := [5]int{1, 2, 3, 4, 5}
	process(&arr)
	fmt.Println(arr)

}
func process(a *[5]int) {
	a[4] = 20
}
