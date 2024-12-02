package main

import (
	"fmt"
	"math"
)

// Задачи на линейное программирование

// 1. Сумма цифр числа
func sum_of_digits(num int) int {
	return num/1000 + num/100%10 + num/10%10 + num%10
}

// 2. Преобразование температуры
func temperature_converter(num int) float64 {
	return float64(num)*float64(1.8) + 32
}

// 3. Удвоение каждого элемента массива
func double_an_array(arr []int) {
	arr[0] *= 2
	arr[1] *= 2
	arr[2] *= 2
	arr[3] *= 2
}

// 4. Объединение строк
func concat_strings(str1 string, str2 string) string {
	return str1 + " " + str2
}

// 5. Расчет расстояния между двумя точками
type Point struct {
	x int
	y int
}

func count_distance(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p2.x)-float64(p1.x), 2) + math.Pow(float64(p2.y)-float64(p1.y), 2))
}

// Задачи с условным оператором

// 1. Проверка на четность/нечетность
func is_even(num int) string {
	if num%2 == 0 {
		return "Четное"
	}
	return "Нечетное"
}

// 2. Проверка высокосного года
func is_leap_year(year int) string {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return "Високосный"
	}
	return "Не високосный"
}

// 3. Определение наибольшего из трех чисел
func biggest_number(num1 int, num2 int, num3 int) int {
	if num1 >= num2 && num1 >= num3 {
		return num1
	} else if num2 >= num1 && num2 >= num3 {
		return num2
	}
	return num3
}

// 4. Категория возраста
// [0, 11] - ребёнок
// [12, 17] - подросток
// [18, 64] - взрослый
// [65, ...] - пожилой
func age_category(age int) string {
	if 0 <= age && age < 12 {
		return "Ребенок"
	} else if 12 <= age && age < 18 {
		return "Подросток"
	} else if 18 <= age && age < 65 {
		return "Взрослый"
	}
	return "Пожилой"
}

// 5. Проверка делимости на 3 и 5
func is_divisible_by_3_5(num int) string {
	if num%3 == 0 && num%5 == 0 {
		return "Делится"
	}
	return "Не делится"
}

// Задачи на циклы

// 1. Факториал числа
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

// 2. Числа Фибоначчи
func fibonacci(count int) {
	var current int
	prev := 1
	pre_prev := 0
	if count > 1 {
		fmt.Print("0 1 ")
	} else if count > 0 {
		fmt.Print("0 ")
	}
	for i := 0; i < count-2; i++ {
		current = pre_prev + prev
		fmt.Printf("%d ", current)
		pre_prev = prev
		prev = current
	}
	fmt.Println()
}

// 3. Реверс массива
func reverse_array(arr []int) {
	var tmp int
	length := len(arr)
	for i := 0; i < length/2; i++ {
		tmp = arr[i]
		arr[i] = arr[length-i-1]
		arr[length-i-1] = tmp
	}
}

// 4. Поиск простых чисел
func print_prime_numbers(num int) {
	for n := 2; n <= num; n++ {
		if is_prime(n) {
			fmt.Printf("%d ", n)
		}
	}
	fmt.Println()
}

func is_prime(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// 5. Сумма чисел в массиве
func array_sum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println(sum_of_digits(1234))

	fmt.Printf("%.1f (Fahrenheit)\n", temperature_converter(23))

	array := []int{1, 2, 3, 4}
	double_an_array(array[:])
	fmt.Println(array)

	fmt.Println(concat_strings("hello", "world"))

	fmt.Println(count_distance(Point{1, 2}, Point{3, 4}))

	fmt.Println(is_even(12))

	fmt.Println(is_leap_year(2024))

	fmt.Println(biggest_number(1, 2, 3))

	fmt.Println(age_category(20))

	fmt.Println(is_divisible_by_3_5(75))

	fibonacci(8)

	reverse_array(array)
	fmt.Println(array)

	print_prime_numbers(10)

	fmt.Println(array_sum(array))
}
