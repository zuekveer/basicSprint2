/*
Чем `sync.Mutex` отличается от `sync.RWMutex`?

**Задача:**

Реализуй потокобезопасный счетчик `SafeCounter` со следующими методами:

```go
type SafeCounter struct {
	// ваш код
}

func (c *SafeCounter) Inc()
func (c *SafeCounter) Value() int
```
*/

package task2

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