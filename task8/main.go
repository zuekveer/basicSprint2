/*
Почему следующая структура занимает 24 байта, а не 17?
type MyStruct struct {
	A int8
	B int64
	C int8
}
​
Задача:
Оптимизируй структуру MyStruct2 по памяти, чтобы занимала минимально возможное число байт:
type MyStruct2 struct {
	A int8
	B int64
	C int8
}
*/

package task8

type MyStruct2 struct {
	B int64
	A int8
	C int8
}
