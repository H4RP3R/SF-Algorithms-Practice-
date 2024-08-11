// Реализуйте на GO любой из алгоритмов сортировки. Проверьте корректность работы алгоритма с помощью функции
// из Задания №1. В реализованной функции не должно возникать ошибки выхода за границы слайса (index out of range).

// Набор тестовых значений:

// []int{0, 1, 2, 3, 4, 5}
// []int{9, 7, 4, 1, 3, 5}
// []int{0}
// []int{}
// []int{1, 1}
// []int{3, 2, 1}
// []int{5, 15, 2, 13, 7, 16, 10, 2}
// []int{1, 9, 7, 4, 6, 2, 1, 13, 22, -3, 12, 76}

// Придумайте дополнительные тестовые значения самостоятельно.
// Тестовые значения можно генерировать с помощью пакета math/rand. В этом случае не забудьте вызвать в начале программы:
// rand.Seed( time.Now().Unix() )

package main

import (
	"fmt"
	"math/rand"
)

func randIntArray(size, min, max int) []int {
	max++
	arr := make([]int, size)

	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(max-min) + min
	}

	return arr
}

func qSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	pivot := arr[0]
	var less, greater []int
	for _, n := range arr[1:] {
		if n < pivot {
			less = append(less, n)
		} else {
			greater = append(greater, n)
		}
	}

	return append(append(qSort(less), pivot), qSort(greater)...)
}

func main() {
	arr1 := []int{5, 15, 2, 13, 7, 16, 10, 2}
	fmt.Println(arr1)
	fmt.Println(qSort(arr1))

	fmt.Println()
	arr2 := randIntArray(12, 0, 10)
	fmt.Println(arr2)
	fmt.Println(qSort(arr2))
}
