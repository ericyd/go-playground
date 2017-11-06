package main

import "fmt"

func main() {
	fmt.Println("this is a test")

	// in the same pkg, does not need explicit import
	extraFunc()
}
