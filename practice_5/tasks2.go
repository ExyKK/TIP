package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// 1. Перевод чисел из одной системы счисления в другую
func num_sys_converter(num string, current_base int, target_base int) string {
	decimal, err := strconv.ParseInt(num, current_base, 64)
	if err != nil {
		fmt.Println("Ошибка при преобразовании числа:", err)
		return ""
	}

	result := strconv.FormatInt(decimal, target_base)
	return result
}

// 2. Решение квадратного уравнения
func solve_quadratic(a, b, c float64) {
	fmt.Printf("%.0fx2%+.0fx%+.0f = 0\n", a, b, c)
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		fmt.Println("Нет корней")
	} else if discriminant == 0 {
		fmt.Printf("Один корень: %.2f\n", -b/(2*a))
	} else {
		fmt.Printf("Два корня:\nx1 = %.2f\nx2 = %.2f\n", (-b-math.Sqrt(discriminant))/(2*a), (-b+math.Sqrt(discriminant))/(2*a))
	}
}

// 3. Сортировка чисел по модулю
func sort_abs(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return math.Abs(float64(arr[i])) < math.Abs(float64(arr[j]))
	})
	fmt.Printf("Отсортированный массив: %v\n", arr)
}

// 4. Слияние двух отсортированных массивов
func merge_sorted_arrays(arr1, arr2 []int) {
	merged := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged = append(merged, arr1[i])
			i++
		} else {
			merged = append(merged, arr2[j])
			j++
		}
	}

	fmt.Printf("Отсортированный массив: %v\n", merged)
}

// 5. Нахождение подстроки в строке без использования встроенных функций
func find_substring(text, pattern string) int {
	for i := 0; i <= len(text)-len(pattern); i++ {
		j := 0
		for j < len(pattern) && text[i+j] == pattern[j] {
			j++
		}

		if j == len(pattern) {
			return i
		}
	}
	return -1
}

// 1. Калькулятор с расширенными операциями
func calc(input string) {
	splitInput := strings.Split(input, " ")
	num1, _ := strconv.ParseFloat(splitInput[0], 64)
	num2, _ := strconv.ParseFloat(splitInput[2], 64)
	operator := splitInput[1]

	result := 0.0
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Деление на ноль недопустимо")
			return
		}
		result = num1 / num2
	case "^":
		result = math.Pow(num1, num2)
	case "%":
		if num2 == 0 {
			fmt.Println("Деление на ноль недопустимо")
			return
		}
		result = math.Mod(num1, num2)
	default:
		fmt.Println("Недопустимая операция")
		return
	}
	fmt.Printf("Результат: %f\n", result)
}

// 2. Проверка палиндрома
func is_palindrome(input string) bool {
	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, ",", "")
	input = strings.ReplaceAll(input, ".", "")

	runes := []rune(input)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		if runes[i] != runes[length-i-1] {
			return false
		}
	}
	return true
}

// 3. Нахождение пересечения трех отрезков
type Interval struct {
	first, last float64
}

func is_intersect(a, b, c Interval) bool {
	start := max(a.first, b.first, c.first)
	end := min(a.last, b.last, c.last)
	return start <= end
}

// 4. Выбор самого длинного слова в предложении
func get_longest_word(input string) string {
	input = strings.ReplaceAll(input, ",", "")
	input = strings.ReplaceAll(input, ".", "")
	inputArray := strings.Split(input, " ")

	longest := ""
	for i := 0; i < len(inputArray); i++ {
		if len(inputArray[i]) > len(longest) {
			longest = inputArray[i]
		}
	}
	return longest
}

// 5. Проверка высокосного года
func is_leap_year(year int) string {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return "Високосный"
	}
	return "Не високосный"
}

// 1. Числа Фибоначчи до определенного числа
func fibonacci(num int) {
	a, b := 0, 1
	for a <= num {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}

// 2. Определение простых чисел в диапазоне
func primes_interval(start, end int) {
	for i := start; i <= end; i++ {
		if is_prime(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func is_prime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// 3. Числа Армстронга в заданном диапазоне
func armstrong_interval(start, end int) {
	for i := start; i <= end; i++ {
		if is_armstrong(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func is_armstrong(num int) bool {
	numStr := string(num)
	n := len([]rune(numStr))
	sum := 0
	for i := 0; i < n; i++ {
		sum += int(numStr[i] - '0')
	}
	return math.Pow(float64(sum), float64(n)) == float64(num)
}

// 4. Реверс строки
func reverse_string(input string) string {
	runes := []rune(input)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		tmp := rune(input[i])
		runes[i] = runes[length-i-1]
		runes[length-i-1] = tmp
	}
	return string(runes)
}

// 5. Нахождение наибольшего общего делителя (НОД)
func find_gcd(num1, num2 int) int {
	for num2 != 0 {
		num1, num2 = num2, num1%num2
	}
	return num1
}

func main() {
	fmt.Println(num_sys_converter("5", 10, 2))
	solve_quadratic(1, 3, -70)
	sort_abs([]int{-5, 3, 4, 7, -8, -9})
	merge_sorted_arrays([]int{1, 3, 5, 12}, []int{8, 11, 13})
	fmt.Println(find_substring("samplestring", "str"))
	calc("21 / 3")
	fmt.Println(is_palindrome("А роза упала на лапу Азора."))
	fmt.Println(is_intersect(Interval{1, 10}, Interval{8, 11}, Interval{-1, 9}))
	fmt.Println(get_longest_word("Hello, my beautiful world"))
	fmt.Println(is_leap_year(2024))
	fibonacci(13)
	primes_interval(1, 20)
	armstrong_interval(1, 500)
	fmt.Println(reverse_string("ABCD"))
	fmt.Println(find_gcd(48, 36))
}
