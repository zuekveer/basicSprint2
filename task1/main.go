/*
Что произойдёт, если вызвать `go someFunc()` внутри `main`, но не дождаться завершения этой горутины?

**Задача:**

Напиши функцию `RunNTimesConcurrently`, которая запускает `n` горутин, каждая из которых печатает свой номер, и программа корректно завершается после их завершения.

```go

	func RunNTimesConcurrently(n int) {
		// ваш код
	}

```
*/
package task1

import (
	"fmt"
	"sync"
)

func RunNTimesConcurrently(n int) {
	var wg sync.WaitGroup

	for i := range n {
		wg.Add(1)
		go func(num int) {
			defer wg.Done() 
			fmt.Println(num)
		}(i) 
	}

	wg.Wait()
}

func main() {
	RunNTimesConcurrently(5)
}
