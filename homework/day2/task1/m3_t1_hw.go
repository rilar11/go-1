/*
Задача №1.
Программа получает на вход последовательность из 5 целых чисел.

Вам нужно определить вид последовательности:
 - возрастающая
 - убывающая
 - случайная
 - постоянная

В качестве ответа следуют выдать прописными латинскими буквами тип последовательности:
1. ASCENDING (строго возрастающая)
2. WEAKLY ASCENDING (нестрого возрастающая, то есть неубывающая)
3. DESCENDING (строго убывающая)
4. WEAKLY DESCENDING (нестрого убывающая, то есть невозрастающая)
5. CONSTANT (постоянная)
7. RANDOM (случайная)

Примеры входных и выходных данных:
In: 11 9 4 2 -1  Out: DESCENDING
In: 3 8 8 11 12  Out: WEAKLY ASCENDING
In: 2 -1 7 21 1  Out: RANDOM
In: 5 5 5 5 5     Out: CONSTANT

Подсказка: используем метод строки strings.Split()
*/

package main

import "fmt"
func main() {
	//Запрашиваем последовательность
	var a, b, c, d, e int
	fmt.Print("Введите 5 чисел: ")
	fmt.Scan (&a, &b, &c, &d, &e)

	chislo := []int{a, b, c, d, e}

	// По умолчанию определяем все типы последовательности как true
	constant := true
	ascending := true
	weaklyascending := true
	descending := true
	weaklydescending :=true

	//Проверяем элементы для определения типа последовательности
	for i := 0; i < len(chislo)-1; i++ {
		if chislo[i] != chislo[i+1] {
			constant = false
		}
		if chislo[i] >= chislo[i+1] {
			ascending = false
		}
		if chislo[i] > chislo[i+1] {
			weaklyascending = false
		}
		if chislo[i] <= chislo[i+1] {
			descending = false
		}
		if chislo[i] < chislo[i+1] {
			weaklydescending = false
		}
	}
	// выводим тип последовательности
	switch {
	case constant:
		fmt.Println("CONSTANT")
	case ascending:
		fmt.Println("ASCENDING")
	case weaklyascending:
		fmt.Println("WEAKLY ASCENDING")
	case descending:
		fmt.Println("DESCENDING")
	case weaklydescending:
		fmt.Println("WEAKLY DESCENDING")
	default:
		fmt.Println("RANDOM")
	}
}

// Ваш код
