package main
package test

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

func test() {
	fmt.Println(add(55, 13))
}