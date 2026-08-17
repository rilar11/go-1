/*
Задача № 4. Проверить, является ли четырехзначное число палиндромом
Пример:
Вход: 1221  Выход: 1221 - палиндром
Вход: 1234  Выход: 1234 - не палиндром
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	var chislo int

	// Запрашиваем число
fmt.Print("Введите целое 4х значное число в диапозоне от 1000 до 9999: ")
_, err := fmt.Scan(&chislo)
if err != nil {
  fmt.Println("При вводе числа произошла ошибка:", err)
os.Exit(1)
}

// проверка на соответствие диапозона
if chislo < 1000 || chislo > 9999 {
  fmt.Println("Число должно быть в диаозоне от 1000 до 9999")
  os.Exit(1)
}

// выявляем цифры из числа в качестве переменных
cifra1 := chislo / 1000
cifra2 := (chislo / 100) % 10
cifra3 := (chislo / 10) % 10
cifra4 := chislo % 10

//проверяем число на палиндром
if cifra1 == cifra4 && cifra2 == cifra3 {
	fmt.Printf("%d - палиндром\n", chislo)
} else {
	fmt.Printf("%d - не палиндром\n", chislo)
}

}