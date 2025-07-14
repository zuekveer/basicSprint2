/*
Как Go-планировщик решает, на каком потоке запускать горутину? Как связаны M:N модель, GOMAXPROCS и планирование?

Задача:
Запусти 100_000 горутин, каждая из которых спит 1 секунду, и измерь общее время выполнения.
func main() {
	// запуск горутин
}
*/

package task6

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(100_000)

	for range 100_000 {
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Second)
		}()
	}

	wg.Wait()
	fmt.Printf("Total time: %v\n", time.Since(start))
}