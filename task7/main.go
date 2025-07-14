/*
Можно ли в Go создать функцию, работающую с любым типом, но ограничить только числовыми типами?

Задача:
Напиши обобщённую функцию SumSlice, которая суммирует элементы слайса чисел (int, float64, ...):
func SumSlice[T Number](items []T) T {
	//
}
*/

package task7

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func SumSlice[T constraints.Integer | constraints.Float](items []T) T {
	var sum T
	for _, v := range items {
		sum += v
	}
	return sum
}
func main() {
	ints := []int{1, 2, 3, 4, 5}
	fmt.Println(SumSlice(ints)) 

	floats := []float64{1.1, 2.2, 3.3}
	fmt.Println(SumSlice(floats)) 
}
