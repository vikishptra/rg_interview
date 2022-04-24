package main

import "fmt"

func downServer(a []int) int {
	var result = 0
	c := []int{}
	b := []int{}
	for i := 0; i < len(a); i++ {
		if a[i] == 0 {
			b = append(b, a[i])
		}
	}
	for i := 0; i < len(b); i++ {
		if i < len(b) || len(b)%5 == 0 {
			c = append(c, b[i])
			if len(c)%5 == 0 {
				result = result + 1
				if len(c) == len(a) {
					result = result - 1
				}
			}
		}
	}
	return result
}

func main() {

	a := []int{1, 0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0} //2
	result := downServer(a)
	fmt.Println(result)
	a = []int{1, 0, 1, 1, 0, 0, 0, 0, 0, 1} //1
	result = downServer(a)
	fmt.Println(result)
	a = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0} //1
	result = downServer(a)
	fmt.Println(result)
	a = []int{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0} //2
	result = downServer(a)
	fmt.Println(result)
}
