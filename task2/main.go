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