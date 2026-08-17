/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример: 
вход: 346, выход: 643
вход: 120, выход: 021
вход: 100, выход: 001
*/
package main

import (
  "fmt"
  "os"
  "strconv"
)

func main() {
  var chislo int
    
  // Запрашиваем число
  fmt.Print("Введите целое 3х значное число в диапозоне от 100 до 999: ")
  _, err := fmt.Scan(&chislo)
  if err != nil {
  fmt.Println("При вводе числа произошла ошибка:", err)
  os.Exit(1)
  }

  // проверка на соответствие диапозона
if chislo < 100 || chislo > 999 {
  fmt.Println("Число должно быть в диаозоне от 100 до 999")
  os.Exit(1)
}

// преобразуем в строку
chislo_string := strconv.Itoa(chislo)
inv_chislo := ""
// Вертаем число
for i := len(chislo_string) - 1; i >= 0; i-- {
	inv_chislo += string(chislo_string[i])
} 

fmt.Printf ("Ревёрс числа : %v\n", inv_chislo)
}