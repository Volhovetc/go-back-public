package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Создаём каналы для передачи данных между горутинами
	numbers := make(chan int, 10) // для передачи чисел от первой ко второй горутине
	squares := make(chan int)     // для передачи квадратов обратно в main

	// Запускаем первую горутину
	go generateNumbers(numbers)

	// Запускаем вторую горутину
	go squareNumbers(numbers, squares)

	// Собираем результаты в слайс
	var results []int
	for i := 0; i < 10; i++ {
		square := <-squares
		results = append(results, square)
	}

	// Закрываем каналы
	close(numbers)
	close(squares)

	// Выводим результаты
	fmt.Println("Квадраты чисел:", results)
}

// Первая горутина - генерирует случайные числа
func generateNumbers(ch chan<- int) {
	for i := 0; i < 10; i++ {
		num := rand.Intn(101) // случайное число от 0 до 100
		ch <- num
	}
}

// Вторая горутина - возводит числа в квадрат
func squareNumbers(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * num
	}
}
