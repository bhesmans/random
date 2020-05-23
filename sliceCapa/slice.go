package main

import "fmt"

func setEl(sl []int, i, a int) {
	if i < cap(sl) {
		sl[0:cap(sl)][i] = a
	}
}

func main() {
	a := make([]int, 2, 3)
	a[0] = 1
	a[1] = 2
	// a[2] = 666 I swear I have the capacity to do it !

	fmt.Println(a)

	setEl(a, 0, 0)
	setEl(a, 1, 1)
	setEl(a, 2, 666)

	fmt.Println(a)
	fmt.Println(len(a))

	fmt.Println(a[0:3])
}
