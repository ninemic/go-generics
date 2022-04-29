package main

import "fmt"

func main() {
	fmt.Println(Sum([]int64{1, 2, 3, 4}...))
	fmt.Println(Sum(1, 2, 3, 4))
}

func Sum[V int64](n ...V) V {
	var s V

	for _, v := range n {
		s += v
	}
	return s
}
