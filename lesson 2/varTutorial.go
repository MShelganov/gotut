package main

import "fmt"

func main() {
	// платформозависимый тип 32, 64
	var num0 int
	var num1 int = 1

	// Автоматически выбранный int
	var autoInt = -10

	// int8, int16, int32, int64
	var bigInt int64 = 1<<32 - 1

	//++num нет
	num1++

	// came

	// пустая строка
	var str string

	fmt.Println(num0, num1, autoInt, bigInt, str)
}
