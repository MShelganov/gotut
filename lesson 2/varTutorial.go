package main

import "fmt"

func main() {
	fmt.Println("===========================================")
	fmt.Println("Numbers Example: ")
	numbersExample()
	fmt.Println("===========================================")
	fmt.Println("String Example: ")
	stringExample()
	fmt.Println("===========================================")
	fmt.Println("Boolean Example: ")
	boolExample()
	fmt.Println("===========================================")
	fmt.Println("Const Example: ")
	constExample()
	fmt.Println("===========================================")
	fmt.Println("Array Example: ")
	arrayExample()
	fmt.Println("===========================================")
	fmt.Println("Slice Example: ")
	sliceExample()
	fmt.Println("===========================================")
}

func numbersExample() {
	// Платформозависимый тип 32, 64
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
	// Возвращает остаток от деления (в этой операции могут принимать участие только целочисленные операнды):
	num1 = 35 % 3 // 35 - 33 = 2
	fmt.Println("Инкримент: ", num0, "35 % 3 = ", num1, "Int: ", autoInt, "Small: ", smallInt, "Medium: ", mediumInt, "Big", bigInt)

	// Так же и с uint
	// uint только положительные
	// byte: синоним типа uint8, представляет целое число от 0 до 255 и занимает 1 байт
	var byteNum byte = 255
	fmt.Println("byte: ", byteNum)

	// Представляет число с плавающей точкой от 1.4*10^(-45) до 3.4*10^(38) (для положительных). Занимает в памяти 4 байта (32 бита)
	var sf float32 = 18.4
	// Представляет число с плавающей точкой от 4.9*10^(-324) до 1.8*10^(308) (для положительных) и занимает 8 байт.
	var bf float64 = 3.14
	fmt.Println("float 32 и 64: ", sf, bf)
	// Деление будет без остатка потому что 10 и 4 целочисленные
	var div1 float32 = 10 / 4     // 2
	var div2 float32 = 10.0 / 4.0 // 2.5
	fmt.Println("Деление 10 на 4: ", div1, "Деление 10.0 на 4.0 ", div2)
}

func boolExample() {
	// Default
	var boolFalse bool
	var boolTrue bool = true
	fmt.Println("Значение по умолчания: ", boolFalse, "True: ", boolTrue)
}

/**
 * Строковый литерал - последовательность символов, заключенная в двойные кавычки.
 * 		- \n: переход на новую строку
 * 		- \r: возврат каретки
 * 		- \t: табуляция
 * 		- \": двойная кавычка внутри строк
 * 		- \\: обратный слеш
 */
func stringExample() {
	// Пустая строка.
	var str string
	fmt.Printf(str)

	// Неявная типизация
	str1 := "Явно не указали тип String."
	fmt.Printf(str1)
}

func constExample() {
	// float
	const pi = 3.1415
	fmt.Println("Pi = ", pi)
	// Int
	const (
		a = 1
		b = 2
		c
		_ // Пропуск переменной
		d = 3
		f
	)
	fmt.Println("Набор констант: ", a, b, c, d, f) // 1, 2, 2, 3, 3
}

/**
 * Массивы представляют последовательность элементов определенного типа.
 */
func arrayExample() {
	// Массив размером 5
	// При таком определении все элементы массива инициализируются значениями по умолчанию.
	var arr [5]int
	fmt.Println(arr)
	// Инициализировать элементы массива другими значениями
	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)
	// Длина массива 5
	var arr2 = [...]int{5, 4, 3, 2, 1}
	fmt.Println(arr2)
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
