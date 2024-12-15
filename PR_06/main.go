package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func task1(n int) {
	if n <= 1 {
		fmt.Println("Число не является простым")
		return
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			fmt.Printf("Число не является простым, найден делитель: %d\n", i)
			return
		}
	}

	fmt.Println("Число является простым")
}

func task2(a, b int) {
	for b != 0 {
		a, b = b, a%b
	}
	fmt.Println("Результат: НОД =", a)
}

func task3(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Println("Шаг", i+1, ":", arr)
	}
	fmt.Println("Результат:", arr)
}

func task4() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}

func task5(n int) {
	var memo = make(map[int]int)
	fib(n, memo)
	for key, value := range memo {
		fmt.Printf("Для n = %v число фибоначи = %v\n", key, value)
	}
}

func fib(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = fib(n-1, memo) + fib(n-2, memo)
	return memo[n]
}

func task6(n int) {
	reversed := 0
	for n != 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	fmt.Println("Результат:", reversed)
}

func task7(levels int) {
	triangle := make([][]int, levels)
	for i := 0; i < levels; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][i] = 1, 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
		fmt.Println(triangle[i])
	}
}

func task8(n int) {
	original, reversed := n, 0
	for n != 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	if original == reversed {
		fmt.Println("Результат: число является палиндромом")
	} else {
		fmt.Println("Результат: число не является палиндромом")
	}
}

func task9(arr []int) {
	maximum, minimum := arr[0], arr[0]
	for _, value := range arr {
		if value > maximum {
			maximum = value
		}
		if value < minimum {
			minimum = value
		}
	}
	fmt.Printf("Результат: максимум = %d, минимум = %d\n", maximum, minimum)
}

func task10() {
	rand.Seed(time.Now().UnixNano())
	target, attempts := rand.Intn(100)+1, 10
	fmt.Println("Я загадал число от 1 до 100. Попробуй угадать!")
	for attempts > 0 {
		attempts--
	}
	fmt.Println("Вы исчерпали попытки!")
}

func task11(n int) {
	original, sum, digits := n, 0, 0
	for temp := n; temp != 0; temp /= 10 {
		digits++
	}
	for n != 0 {
		digit := n % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		n /= 10
	}
	if original == sum {
		fmt.Println("Число является числом Армстронга")
	}
}

func main() {
	// Закомментированные вызовы всех задач
	// task1(17)
	// task2(48, 18)
	// task3([]int{5, 2, 4, 3, 1})
	// task4()
	// task5(10)
	// task6(12345)
	// task7(5)
	// task8(12321)
	// task9([]int{3, 5, 1, 8, 2})
	// task10()
	// task11(153)
	// Добавляйте вызовы по необходимости!
}
