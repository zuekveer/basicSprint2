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
	"sync"
)

type SafeCounter struct {
	value int
	mu    sync.RWMutex
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *SafeCounter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

func main() {
	counter := SafeCounter{}

	var wg sync.WaitGroup
	for range 1000 {
		wg.Add(1)
		go func() {
			counter.Inc()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter.Value())
}
