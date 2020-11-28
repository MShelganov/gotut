package lesson2

import "fmt"

/**
 * Функция представляет блок операторов, которые все вместе выполняют какую-то определенную задачу.
 *
 * По умолчанию каждая программа на языке Go должна содержать как минимум одну функцию - функцию main,
 * которая является входной точкой в приложение
 *
 * func имя_функции (список_параметров) (типы_возвращаемых_значений, ..) {
 *     выполняемые_операторы
 * }
 */

func funcmain() {
	// Вызов функции
	hello()
	hello()
	hello()

	fmt.Println("x + y = ", add(4, 5)) // x + y = 9
	z := add(20, 6)                    // x + y = 26
	fmt.Println("x + y = ", z)
	z = add(2, 16) // x + y = 18
	fmt.Println("x + y = ", z)
	d := addFloat(3.4, 5.6, 1.2)
	fmt.Println("a + b + c = ", d)
	s := addNumbers(1, 2)    // 3
	s += addNumbers(3, 4, 5) // 15
	fmt.Println("sum = ", s)
	arr := []int{5, 6, 7, 2, 3}
	fmt.Println("slice = ", addNumbers(arr...))
	var age, name = addNameAndNumbers(4, 5, "Tom", "Simpson")
	fmt.Printf("Имя: %s, годиков: %d\n", name, age) // Имя: Tom Simpson, годиков: 9

	// Также функция может передаваться в качестве параметра в другую функцию.
	var f func(int, int) int = add
	fmt.Println("x + y = ", f(12, 6)) // 18
	// Анонимные функции - это функции, которым не назначено имя.
	f1 := func(x, y int) int { return x + y }
	fmt.Println(f1(3, 4)) // 7

	callPanic(false)
}

func hello() {
	fmt.Println("Hello World")
}

func add(x int, y int) int {
	return x + y
}

func addFloat(a, b, c float32) float32 {
	return a + b + c
}

func addNumbers(numbers ...int) int {
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// И также в даном случае можно было бы использовать именованные результаты: (z int, fullName string)
func addNameAndNumbers(x, y int, firstName, lastName string) (int, string) {
	var z int = x + y
	var fullName = firstName + " " + lastName
	return z, fullName
}

/**
 * Оператор panic позволяет сгенерировать ошибку и выйти из программы.
 */
func callPanic(b bool) {
	if b {
		panic("Better call Panic!")
	}
}
