package main

import "fmt"

func main() {
	x_train := [][]int{
		{1, 0, 0, 1},
		{1, 0, 0, 1},
		{1, 0, 0, 1},
		{1, 0, 0, 1},
	}
	y_train := []int{1, 0, 0, 1}
	x_test := []int{1, 0, 1, 0}

	predicWithNaiveBasian(x_train, y_train, x_test)
}

func predicWithNaiveBasian(x_train [][]int, y_train []int, x_test []int) int {
	fmt.Println(x_train)
	fmt.Println(y_train)
	fmt.Println(x_test)

	return 0
}
