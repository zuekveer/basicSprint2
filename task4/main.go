/*
Можно ли закрыть канал, из которого никто не читает? Что произойдёт?

**Задача:**

Реализуй функцию `FanIn`, объединяющую данные из двух каналов в один:

```go
func FanIn(a, b <-chan int) <-chan int {
	// ваш код
}
```
*/

package task4

import (
	"fmt"
	"time"
)

func FanIn(a, b <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out) 

		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					continue
				}
				out <- v
			case v, ok := <-b:
				if !ok {
					b = nil 
					continue
				}
				out <- v
			}
		}
	}()

	return out
}
func main() {
	a := make(chan int)
	b := make(chan int)

	go func() {
		defer close(a)
		a <- 1
		a <- 2
	}()

	go func() {
		defer close(b)
		b <- 3
		time.Sleep(100 * time.Millisecond)
		b <- 4
	}()

	for v := range FanIn(a, b) {
		fmt.Println(v) 
	}
}
