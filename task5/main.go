/*
Приведи пример ситуации, в которой возникает дедлок при использовании каналов.

Задача:
Найди и исправь дедлок:
func main() {
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	<-ch
	<-ch // ← где ошибка?
}
*/

package task5

func main() {
    ch := make(chan int, 2)
    go func() {
        ch <- 42
        ch <- 100  
    }()
    <-ch  
    <-ch  
}