package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// 1. Проверка на простоту
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

// 2. Наибольший общий делитель (НОД)
func find_gcd(num1, num2 int) int {
	for num2 != 0 {
		num1, num2 = num2, num1%num2
	}
	return num1
}

// 3. Сортировка пузырьком
func bubble_sort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("Результат:", arr)
}

// 4. Таблица умножения в формате матрицы
func print_multiplication_table() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}

// 5. Фибоначчи с мемоизацией
func fibonacci_memo(n int) {
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

// 6. Обратные числа
func reverse_num(n int) int {
	reversed := 0
	for n != 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return reversed
}

// 7. Треугольник Паскаля
func print_pascal_triangle(levels int) {
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

// 8. Число палиндром
func is_num_palindrome(n int) bool {
	original := n
	reversed := 0

	for n != 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}

	if original == reversed {
		return true
	}
	return false
}

// 9. Нахождение максимума и минимума в массиве
func find_max_min(arr []int) {
	if len(arr) == 0 {
		fmt.Println("Результат: массив пустой")
		return
	}

	maximum := arr[0]
	minimum := arr[0]

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

// 10. Игра "Угадай число"
func guess_a_num() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	attempts := 10
	var guess int

	fmt.Println("Я загадал число от 1 до 100. У вас есть", attempts, "попыток, чтобы его угадать.")

	for attempts > 0 {
		fmt.Print("Введите вашу попытку: ")
		fmt.Scan(&guess)

		if guess < target {
			fmt.Println("Больше!")
		} else if guess > target {
			fmt.Println("Меньше!")
		} else {
			fmt.Println("Поздравляю! Вы угадали число:", target)
			return
		}

		attempts--
		fmt.Println("Осталось попыток:", attempts)
	}

	fmt.Println("Вы исчерпали все попытки. Загаданное число было:", target)
}

// 11. Числа Армстронга
func is_armstrong(n int) bool {
	original := n
	digits := 0
	sum := 0

	for temp := n; temp != 0; temp /= 10 {
		digits++
	}

	for n != 0 {
		digit := n % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		n /= 10
	}

	if original == sum {
		return true
	}
	return false
}

// 12. Подсчет слов в строке
func count_words(input string) int {
	wordCount := make(map[string]int)

	input = strings.ToLower(input)

	var cleanedInput strings.Builder
	for _, char := range input {
		if unicode.IsLetter(char) || unicode.IsSpace(char) {
			cleanedInput.WriteRune(char)
		}
	}

	words := strings.Fields(cleanedInput.String())

	for _, word := range words {
		wordCount[word]++
	}

	return len(wordCount)
}

// 13. Игра "Жизнь" (Conway's Game of Life)
func game_of_life(lifeTimeSeconds int) {
	board := [][]int{
		{0, 1, 0, 0, 0},
		{0, 0, 1, 1, 0},
		{1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
	}

	fmt.Println("Начальное состояние:")
	printBoard(board)

	for i := 0; i < lifeTimeSeconds; i++ {
		time.Sleep(1 * time.Second)
		board = updateBoard(board)
		fmt.Printf("Шаг %d:\n", i+1)
		printBoard(board)
	}
}

const (
	rows    = 5
	columns = 5
)

func printBoard(board [][]int) {
	fmt.Println("+---+---+---+---+---+")
	for _, row := range board {
		for _, cell := range row {
			if cell == 1 {
				fmt.Print("| X ") // Живая клетка
			} else {
				fmt.Print("|   ") // Мертвая клетка
			}
		}
		fmt.Println("|")
		fmt.Println("+---+---+---+---+---+")
	}
	fmt.Println()
}

func updateBoard(board [][]int) [][]int {
	newBoard := make([][]int, rows)
	for i := range newBoard {
		newBoard[i] = make([]int, columns)
		copy(newBoard[i], board[i])
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			liveNeighbors := countLiveNeighbors(board, i, j)

			if board[i][j] == 1 {
				if liveNeighbors < 2 || liveNeighbors > 3 {
					newBoard[i][j] = 0
				}
			} else {
				if liveNeighbors == 3 {
					newBoard[i][j] = 1
				}
			}
		}
	}

	return newBoard
}

func countLiveNeighbors(board [][]int, x, y int) int {
	liveCount := 0
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, direction := range directions {
		nx, ny := x+direction[0], y+direction[1]
		if nx >= 0 && nx < rows && ny >= 0 && ny < columns {
			liveCount += board[nx][ny]
		}
	}

	return liveCount
}

// 14. Цифровой корень числа
func find_digital_root(n int) int {
	if n < 10 {
		return n
	}
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return find_digital_root(sum)
}

// 15. Римские цифры
func int_to_roman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}
