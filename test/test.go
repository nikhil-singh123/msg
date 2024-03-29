package main

import "fmt"

func main() {
	var a []int
	var b []int
	a = []int{1, 2, 3}
	b = []int{4, 5, 6}

	a = append(a, b...)

	fmt.Println(a)
}
