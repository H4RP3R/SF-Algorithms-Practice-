// Реализуйте функцию для проверки того, что слайс отсортирован по возрастанию.
// Функция должна иметь сигнатуру:
// func checkSliceIsSorted(a []int) bool

package sortcheck

func checkSliceIsSorted(a []int) bool {
	if len(a) == 0 {
		return true
	}

	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}

	return true
}
