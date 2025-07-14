/*
то произойдёт, если select не имеет готовых кейсов и нет default?

Задача:
Напиши функцию WaitAny, которая принимает 2 канала и возвращает значение из того, который первым прислал данные:
func WaitAny(a, b <-chan string) string {
	// ваш код
}
*/

package task3

import (
	"fmt"
	"time"
)

func WaitAny(a, b <-chan string) string {
	select {
	case v := <-a:
		return v
	case v := <-b:
		return v
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "from ch2"
	}()

	result := WaitAny(ch1, ch2)
	fmt.Println(result) 
}
