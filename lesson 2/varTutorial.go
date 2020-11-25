package main

import "fmt"

func main() {

	sliceExample()
	numbersExample()
	strinngExample()
}

/**
 * Срез имеет длину и емкость.
 * Длина среза - это количество содержащихся в нем элементов.
 * Емкость среза - это количество элементов в базовом массиве, считая от первого элемента в срезе.
 * Длину и емкость среза s можно получить с помощью выражений len (s) и cap (s).
 * Вы можете увеличить длину среза, повторно разрезав его, если он имеет достаточную емкость.
 * Попробуйте изменить одну из операций среза в программе-примере, чтобы расширить ее возможности, и посмотрите, что произойдет.
 */
func sliceExample() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func numbersExample() {
	// платформозависимый тип 32, 64
	var num0 int
	// Через ","
	var num1, num2, num3 int = 1, 2, 3
	fmt.Println("Куча Int: ", num0, num1, num2, num3)

	// Автоматически выбранный int
	var autoInt = -10

	// int8, int16, int32, int64
	var (
		smallInt  int8  = -128
		mediumInt int16 = -32768
		bigInt    int64 = 1<<32 - 1
	)
	//++num нет
	num0++
	fmt.Println("Инкримент: ", num0, "Int: ", autoInt, "Small: ", smallInt, "Medium: ", mediumInt, "Big", bigInt)

	// Так же и с uint
	// uint только положительные
	// byte: синоним типа uint8, представляет целое число от 0 до 255 и занимает 1 байт
	var byteNum byte = 255
	fmt.Println("byte: ", byteNum)
}

func strinngExample() {
	var str string
	fmt.Printf(str)

	str1 := "string test 1"
	fmt.Printf(str1)
}
