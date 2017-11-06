package main

import "fmt"

func extra() {
	fmt.Println("extra print")
}

func add(x, y int) (z int) {
	z = x + y
	return
}

// Rules:
// 1. it kind of uses a functional-eque function defition
// e.g. a -> b -> c
// curryAdd returns a function with parameters (b int) which itself returns an int
// 2. The returned function cannot be named
func curryAdd(a int) func(b int) int {
	return func(b int) int {
		return add(a, b)
	}
}

// huh, this is definitely different
// if you want to pass a function to another function, you need to
// specify it's signature in the parameter
func firstClass(f func(b int) int) int {
	f(2)
	return 1
}

// func curry(f func(vars ...interface{}) func(vars ...interface{}) interface{}) {
// 	return
// }
