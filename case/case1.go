package main

import "fmt"

func case1(a []int, b []int) []int {

	result := []int{}

	for i := 0; i < len(a); i++ {
		k := i
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				j = len(b)
			} else if a[i] != b[k] && a[i] != b[j] && j+1 == len(b) {
				result = append(result, a[i])
			}
		}
	}

	return result

}

func main() {

	/*
		1.jika array lama terdapat di array baru maka tidak masuk result
		2.lalu jika array lama tidak terdapat di array baru maka masuk result
	*/
	arrayLama := []int{3, 4, 3, 9, 8, 2}
	arrayBaru := []int{7, 5, 2, 3, 5, 0}

	result := case1(arrayLama, arrayBaru)

	fmt.Println(result)
}
