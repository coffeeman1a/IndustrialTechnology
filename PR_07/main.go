package main

import (
	"fmt"
	"math"
)

// 1. Площадь треугольника
func triangleArea(base, height float64) float64 {
	return 0.5 * base * height
}
func task1(height, base float64) {
	fmt.Printf("Площадь треугольника: %f\n", triangleArea(base, height))
}

// 2. Сортировка пузырьком
func sortArray(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return arr
}
func task2(arr []int) {
	fmt.Println("Исходный массив:", arr)
	fmt.Println("Отсортированный массив:", sortArray(arr))
}

// 3. Сумма квадратов чётных чисел
func sumOfSquares(n int) int {
	sum := 0
	for i := 2; i <= n; i += 2 {
		sum += i * i
	}
	return sum
}
func task3(num int) {
	fmt.Printf("Сумма квадратов чётных чисел от 1 до %d: %d\n", num, sumOfSquares(num))
}

// 4. Проверка палиндрома
func isPalindrome(s string) bool {
	for i, j := 0, len([]rune(s))-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
func task4(input string) {
	if isPalindrome(input) {
		fmt.Printf("Строка '%s' является палиндромом\n", input)
	} else {
		fmt.Printf("Строка '%s' не является палиндромом\n", input)
	}
}

// 5. Проверка на простое число
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func task5(n int) {
	if isPrime(n) {
		fmt.Printf("Число %d является простым\n", n)
	} else {
		fmt.Printf("Число %d не является простым\n", n)
	}
}

// 6. Генерация простых чисел
func generatePrimes(limit int) []int {
	var primes []int
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}
func task6(n int) {
	fmt.Printf("Простые числа до %d: %v\n", n, generatePrimes(n))
}

// 7. Перевод в двоичную систему
func toBinary(n int) string {
	if n == 0 {
		return "0"
	}
	binary := ""
	for n > 0 {
		binary = fmt.Sprintf("%d%s", n%2, binary)
		n /= 2
	}
	return binary
}
func task7(n int) {
	fmt.Printf("Число %d в двоичной системе: %s\n", n, toBinary(n))
}

// 8. Нахождение максимального элемента
func findMax(arr []int) int {
	max := math.MinInt
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}
func task8(arr []int) {
	fmt.Println("Исходный массив:", arr)
	fmt.Println("Максимальный элемент:", findMax(arr))
}

// 9. НОД двух чисел
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func task9(a, b int) {
	fmt.Printf("НОД(%d, %d) = %d\n", a, b, gcd(a, b))
}

// 10. Сумма элементов массива
func sumArray(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}
func task10(arr []int) {
	fmt.Println("Исходный массив:", arr)
	fmt.Println("Сумма элементов массива:", sumArray(arr))
}

// Главная функция
func main() {
	task1(2.5, 5) // Площадь треугольника
	task2([]int{64, 34, 25, 12, 22, 11, 90}) // Сортировка массива
	task3(15) // Сумма квадратов чётных чисел
	task4("sas") // Проверка палиндрома
	task5(25) // Проверка на простое число
	task6(52) // Генерация простых чисел
	task7(25) // Перевод числа в двоичную систему
	task8([]int{64, 34, 25, 12, 22, 11, 90}) // Нахождение максимального элемента
	task9(12, 21) // НОД двух чисел
	task10([]int{64, 34, 25, 12, 22, 11, 90}) // Сумма элементов массива
}
