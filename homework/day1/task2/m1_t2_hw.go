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
//Дробим число на цифры
cifra1 := chislo / 100
cifra2 := (chislo /10) % 10
cifra3 := chislo % 10 

//Собираем ревёрс число
inv_chislo := (cifra3 * 100) + (cifra2*10) + cifra1

//Рисуем вывод
fmt.Printf ("Ревёрс числа : %03d\n", inv_chislo)
}