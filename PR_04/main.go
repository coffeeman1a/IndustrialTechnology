package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Print(sumOfDigitalNumber())
	// tempConverter()
	// fmt.Print(doubleArray([4]int{1, 2, 3, 4}))
	// fmt.Print(stringMerge([]string{"Hello", "world"}))
	// distBetweenPoints(1, 1, 4, 5)
	// fmt.Print(oddOrEven(5))
	// fmt.Print(leapYearCheck(2021))
	// fmt.Print(maxOfThree(4, 9, 7))
	// fmt.Print(ageCategory())
	// fmt.Print(divisionCheck(16))
	// fmt.Print(factorial(5))
	// fmt.Print(fibonacciNumbers(7))
	// fmt.Print(recursiveArray([]int{1, 2, 3, 4, 5}))
	// fmt.Print(primeNumbers(20))
	fmt.Print(arraySum([]int{1, 2, 3, 4, 5}))
}

// *** Сумма цифр числа ***

func sumOfDigitalNumber() int {
	var n int
	var sum int

	fmt.Print("Введите целое четырёхзначное число: ")
	fmt.Scan(&n)

	sum += numberToDigit(n)
	n = n / 10
	sum += numberToDigit(n)
	n = n / 10
	sum += numberToDigit(n)
	n = n / 10
	sum += numberToDigit(n)
	n = n / 10

	return sum
}

func numberToDigit(n int) int {
	return int(math.Mod(float64(n), 10))
}

// *** Преобразование температуры ***

func tempConverter() {
	converters := map[string]func(float64) float64{
		"Celsius":    celsiusToFahrenheit,
		"Fahrenheit": fahrenheitToCelsius,
	}

	var inputDegrees float64
	var inputScale string

	fmt.Print("Введите градусы: ")
	fmt.Scan(&inputDegrees)
	fmt.Print("Введите шкалу (Celsius/Fahrenheit): ")
	fmt.Scan(&inputScale)

	// Применяем конвертер из мапы без условий
	result := converters[inputScale](inputDegrees)
	fmt.Print(result)
}

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

// *** Удвоение каждого элемента массива ***

func doubleArray(n [4]int) []int {
	res := make([]int, len(n))
	res[0] = n[0] * 2
	res[1] = n[1] * 2
	res[2] = n[2] * 2
	res[3] = n[3] * 2
	return res
}

// *** Объединение строк ***

func stringMerge(s []string) string {
	return s[0] + " " + s[1]
}

// *** Расчет расстояния между двумя точками ***

func distBetweenPoints(x1, y1, x2, y2 int) {
	fmt.Printf("Координаты точки 1: x = %d, y = %d\n", x1, y1)
	fmt.Printf("Координаты точки 2: x = %d, y = %d\n", x2, y2)

	dist := math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))

	fmt.Println(dist)
}

// *** Проверка на четность/нечетность ***

func oddOrEven(n int) string {
	if n%2 == 0 {
		return "Четное"
	} else {
		return "Нечетное"
	}
}

// *** Проверка високосного года ***

func leapYearCheck(year int) string {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return "Високосный"
	} else {
		return "Невисокосный"
	}
}

// *** Определение наибольшего из трех чисел ***

func maxOfThree(v1, v2, v3 int) int {
	max := 0
	if v1 > max {
		max = v1
	}
	if v2 > max {
		max = v2
	}
	if v3 > max {
		max = v3
	}

	return max
}

// *** Категория возраста ***

func ageCategory() string {
	// 0-10 child
	// 11 - 18 teenager
	// 19-59 adult
	// >60 old

	var age int
	fmt.Printf("Введите ваш возраст")
	fmt.Scan(&age)

	if age <= 10 {
		return "ребёнок"
	} else if age > 10 && age < 19 {
		return "подросток"
	} else if age >= 19 && age < 60 {
		return "взрослый"
	} else if age >= 60 {
		return "пожилой"
	}

	return "Необходимо ввести ваш возраст!"
}

// *** Проверка деление на 3 и на 5 ***

func divisionCheck(n int) string {
	if n%5 == 0 && n%3 == 0 {
		return "Делится"
	} else {
		return "Не делится"
	}
}

// *** Факториал числа ***

func factorial(n int) int {
	res := 1

	for n > 1 {
		res *= n
		n--
	}

	return res
}

// *** числа Фибоначчи ***

func fibonacciNumbers(c int) []int {
	var res []int
	for i := 0; i < c; i++ {
		res = append(res, fib(i))
	}
	return res
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// *** Обратный массив ***

func recursiveArray(a []int) []int {
	l := len(a) - 1
	var res []int
	for i := l; i >= 0; i-- {
		res = append(res, a[i])
	}
	return res
}

// ** Поиск простых чисел ***

func primeNumbers(n int) []int {
	var res []int
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			res = append(res, i)
		}
	}
	return res
}

func isPrime(n int) bool {
	d := 2
	for n%d != 0 {
		d++
	}

	return d == n
}

// *** сумма чисел в массиве ***

func arraySum(n []int) int {
	l := len(n)
	sum := 0
	for i := 0; i <= l-1; i++ {
		sum += n[i]
	}
	return sum
}
