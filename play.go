package main

import "fmt"

func main() {
	fmt.Println("this is a test")

	// in the same pkg, does not need explicit import
	z := add(2, 3)
	fmt.Printf("z = %d\n", z)

	// first class functions?
	add2 := curryAdd(2)
	_ = curryAdd(3)

	fmt.Println(add2(7))

	one := firstClass(add2)
	fmt.Println(one)
}
