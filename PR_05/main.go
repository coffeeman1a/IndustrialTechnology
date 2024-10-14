package main

import (
	"errors"
	"fmt"
	"math"
	"math/cmplx"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// res, err := convertInt("10403", 10, 16)
	// if err == nil {
	// 	fmt.Print(res)
	// }

	// x1, x2 := quadraticEquation(1, 14, 45)
	// fmt.Printf("Корни: x1 = %.2f, x2 = %.2f\n", x1, x2)

	// fmt.Print(sortByAbs([]int{1, -10, 4, 5, -34, 6}))

	// fmt.Print(sortedArraytMerge([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}))

	// fmt.Print(subStringSearch("something", "some"))

	// fmt.Print(calculator(10, 0, "/"))

	// fmt.Print(palindrome("tE neT"))

	// fmt.Print(interseptionPoint([]int{1, 5}, []int{3, 6}, []int{4, 8}))

	// fmt.Print(longestWord("Ну, типо оч большое предлождение, какое же слово тут будет самое большое?"))

	// fmt.Print(leapYearCheck(2025))

	// fmt.Print(fibonacciNumbers(77))

	// fmt.Print(primeNumbers(2, 11))

	// fmt.Print(armstrongNumbers(1, 500))

	// fmt.Print(reverseString("ПОСМОТРИТЕ! Я НАОБРОТ! ХИ-ХИ"))

	fmt.Print(gcd(48, 18))

}

// *** конвертация числа в СИ ***

func convertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}

// *** квадратное уравнение ***

func quadraticEquation(a, b, c float64) (complex128, complex128) {
	D := b*b - 4*a*c
	sqrtD := cmplx.Sqrt(complex(D, 0))
	x1 := (-complex(b, 0) + sqrtD) / (2 * complex(a, 0))
	x2 := (-complex(b, 0) - sqrtD) / (2 * complex(a, 0))
	return x1, x2
}

// *** сортировка чисел по модулю  ***

func sortByAbs(a []int) []int {
	sort.Slice(a, func(i, j int) bool {
		return math.Abs(float64(a[i])) < math.Abs(float64(a[j]))
	})

	return a
}

// *** слияние двух отсортированных массивов ***

func sortedArraytMerge(a1, a2 []int) []int {
	var r []int
	return append(append(r, a1[:]...), a2[:]...)
}

// *** нахождение подстроки в строке без использования встроенных функций ***

func subStringSearch(s, subs string) bool {
	return strings.Contains(s, subs)
}

// *** калькулятор с расширенными операциями ***

func calculator(v1, v2 float64, opr string) (float64, error) {
	var res float64
	var validOpr bool = true

	switch opr {
	case "+":
		res = v1 + v2
	case "-":
		res = v1 - v2
	case "*":
		res = v1 * v2
	case "/":
		if v2 != 0 {
			res = v1 / v2
		} else {
			validOpr = false
			return 0, errors.New("Ошибка: деление на ноль")
		}
	case "^":
		res = math.Pow(v1, v2)
	case "%":
		if int(v2) != 0 {
			res = float64(int(v1) % int(v2))
		} else {
			validOpr = false
			return 0, errors.New("Ошибка: деление на ноль в операции %")
		}
	default:
		validOpr = false
		return 0, errors.New("Ошибка: неверный оператор")
	}

	if validOpr {
		return res, nil
	}

	return 0, errors.New("Неизветсаня ошибка")
}

// *** проверка палиндрома ***

func palindrome(s string) bool {
	s = strings.ToLower(strings.Replace(s, " ", "", -1))
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

// *** нахождение пересечения трех отрезков ***

func interseptionPoint(x []int, y []int, z []int) ([]int, error) {
	maxStart := math.Max(float64(x[0]), math.Max(float64(y[0]), float64(z[0])))
	minEnd := math.Min(float64(x[1]), math.Max(float64(y[1]), float64(z[1])))

	if maxStart <= minEnd {
		return []int{int(maxStart), int(minEnd)}, nil
	}

	return []int{0, 0}, errors.New("Точки пересечения не найдено!")
}

// *** выбор самого длинного слова в предложении ***

func longestWord(s string) string {
	maxLen := 0
	idxLenMap := map[int]string{
		0: "",
	}
	re := regexp.MustCompile(`[[:punct:] ]+`)
	cleanedString := re.ReplaceAllString(s, " ")
	r := regexp.MustCompile(`\s+`).Split(cleanedString, -1)
	i := 0
	for i < len(r) {
		idxLenMap[len(r[i])] = r[i]
		if maxLen < len(r[i]) {
			maxLen = len(r[i])
		}
		i++
	}

	return idxLenMap[maxLen]
}

// *** Проверка високосного года ***

func leapYearCheck(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
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

// ** Поиск простых чисел ***

func primeNumbers(lb, rb int) []int {
	var res []int
	for i := lb; i <= rb; i++ {
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

// ** числа армстронга ***

func armstrongNumbers(lb, rb int) []int {
	var res []int
	for i := lb; i <= rb; i++ {
		if isArmstorng(i) {
			res = append(res, i)
		}
	}
	return res
}

func isArmstorng(n int) bool {
	nStr := strconv.Itoa(n)
	nDigits := len(nStr)
	sumOfPowers := 0

	for i := 0; i < nDigits; i++ {
		digit, _ := strconv.Atoi(string(nStr[i]))
		sumOfPowers += int(math.Pow(float64(digit), float64(nDigits)))
	}

	return sumOfPowers == n
}

// *** реверс строки ***

func reverseString(s string) string {
	runes := []rune(s)
	l := len(runes)

	var res []rune
	for i := l - 1; i >= 0; i-- {
		res = append(res, runes[i])
	}

	return string(res)
}

// *** нахождение наибольшего общего делителя (НОД) ***

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
