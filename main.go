package main

import "fmt"

var global1 int
var global2 int = 10
var global3 = "this is ok"

// bad_var := "this isn't ok"

func main() {
	local1 := 9

	var local2 = 5
	var local3 = 7

	fmt.Println(global1, global2, global3, local1, local2, local3)

	loops := 1

	// while oops
	for loops < 10 {
		fmt.Println(loops)
		loops++
	}

	// for loops
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Infinite loop
	// for {
	// 	// run forever
	// }

	// C-style for loop
	// j := 0
	// for j < 10 {
	// 	fmt.Println(string(10))
	// }

	// C-style infinite loop
	// for {
	// 	// run forever
	// }

}
