package main

import (
	"fmt"
	"math"
)

// 1. Функция вычисления площади треугольника
func triangleArea(base float64, height float64) float64 {
	return 0.5 * base * height
}

// 2. Сортировка массива по возрастанию
func sortArray(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		min_idx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min_idx] {
				min_idx = j
			}
		}
		arr[i], arr[min_idx] = arr[min_idx], arr[i]
	}
	return arr
}

// 3. Сумма квадратов чётных чисел
func sumOfSquares(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if i%2 == 0 {
			sum += i * i
		}
	}
	return sum
}

// 4. Проверка палиндрома
func isPalindrome(s string) bool {
	length := len([]rune(s))
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

// 5. Проверка числа на простоту
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

// 6. Генерация последовательности простых чисел
func generatePrimes(limit int) []int {
	var arr []int
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			arr = append(arr, i)
		}
	}
	return arr
}

// 7. Перевод числа в двоичную систему
func toBinary(n int) string {
	if n == 0 {
		return "0"
	}

	binary := ""
	for n > 0 {
		bit := n % 2
		binary = fmt.Sprintf("%d%s", bit, binary)
		n /= 2
	}
	return binary
}

// 8. Нахождение максимального элемента в массиве
func findMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := math.MinInt

	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	return max
}

// 9. Функция вычисления НОД (наибольший общий делитель)
func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// 10. Сумма элементов массива
func sumArray(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func main() {
	fmt.Println(triangleArea(2.5, 5))
	fmt.Println(sortArray([]int{6, 1, 4, 3, 2, 8, 9, 5, 7}))
	fmt.Println(sumOfSquares(15))
	fmt.Println(isPalindrome("шалаш"))
	fmt.Println(isPrime(25))
	fmt.Println(generatePrimes(52))
	fmt.Println(toBinary(25))
	fmt.Println(findMax([]int{64, 34, 25, 12, 22, 11, 90}))
	fmt.Println(gcd(12, 21))
	fmt.Println(sumArray([]int{64, 34, 25, 12, 22, 11, 90}))
}
