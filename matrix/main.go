package main

import "fmt"
import mx "./matrix"

func main() {
	m := mx.NewEmptyMatrix(3, 3)
	fmt.Println(m)
	m.InitData([][]int{
		{2, 0, 4},
		{3, 1, 0},
		{4, 2, 2},
	})
	fmt.Println(m)
	fmt.Println(m.Det())
}
