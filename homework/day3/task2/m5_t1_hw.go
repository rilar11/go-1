/*
	Задача №2
	
	Вход:
	Пользователь должен ввести "правильный пароль", состоящий из:
	цифр, букв латинского алфавита(строчные и прописные) и 
	специальных символов  special = "_!@#$%^&"

	Всего 4 набора различных символов.
	В пароле обязательно должен быть хотя бы один символ из каждого набора.
	Длина пароля от 8(мин) до 15(макс) символов.
	Максимальное количество попыток ввода неправильного пароля - 5.
	Каждый раз выводим номер попытки.
	*Желательно выводить пояснение, почему пароль не принят и что нужно исправить.

	digits = "0123456789"
	lowercase = "abcdefghiklmnopqrstvxyz"
	uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
	special = "_!@#$%^&"

	Выход:
	Написать, что пароль правильный и он принят.

	Пример: 
	хороший пароль -> o58anuahaunH!
	хороший пароль -> aaaAAA111!!!
	плохой пароль -> saucacAusacu8 

	Реализацию оформить через функцию:
	1. checkPassword(pass string) (bool, errors <- на усмотрение)
	2. main() // для интерактива

*/
package main

import (
	"fmt"
	"os"
	"strings"
)

const digits = "0123456789"
const lowercase = "abcdefghiklmnopqrstvxyz"
const uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
const special = "_!@#$%^&"
//функция проверки пароля
func checkPassword (pass string) (bool, error) {
	// создаём слайс для сбора всех проблем с паролем
	var problems []string
	if len(pass) <= 8 || len(pass) >= 15 {
		problems = append(problems,fmt.Sprintf("пароль должен иметь от 8 до 15 симолов; ")) //заполняем слайс
	}
	passwordContainsDigits := false
	passwordContainsLowercase := false
	passwordContainsUppercase := false
	passwordContainsSpecial := false
    //проходимся циклом по каждому символу и ищем его в одной из групп
	for _,symb := range pass {
	  withDigits := strings.ContainsRune(digits, symb)
	  withLowercase := strings.ContainsRune(lowercase, symb)
	  withUppercase := strings.ContainsRune(uppercase, symb)
	  withSpecial := strings.ContainsRune(special, symb)	

	  if withDigits {
		passwordContainsDigits = true
	  }
	  if withLowercase {
		passwordContainsLowercase = true
	  }
	  if withUppercase {
		passwordContainsUppercase = true
	  }
	  if withSpecial {
		passwordContainsSpecial = true
	  }
	}

	if !passwordContainsDigits {
		problems = append(problems,fmt.Sprintf("введенный пароль должен содержать хотя бы одну цифру; "))
	}
	if !passwordContainsLowercase {
		problems = append(problems,fmt.Sprintf("введенный пароль должен содержать хотя бы одну букву нижнего регистра; "))
	}
	if !passwordContainsUppercase {
		problems = append(problems,fmt.Sprintf("введенный пароль должен содержать хотя бы одну букву верхнего регистра; "))
	}
	if !passwordContainsSpecial {
		problems = append(problems,fmt.Sprintf("введенный пароль должен содержать хотя бы один спецсимвол; "))
	}
	//Если в слайсе что-то есть, выводим это как ошибку
    if len(problems) > 0 {
		return false, fmt.Errorf("%s",problems)
	}

	return true, nil
}

func main() {
	var pass string
	for i :=1; i <= 5; i++ {
		fmt.Printf ("Попытка %d\n", i)
		fmt.Print("Введите пароль: ")
		_, err := fmt.Scan(&pass)
		if err != nil {
			fmt.Println("При вводе пароля возникла проблема: " , err)
			continue
		}
		if check, errmsg := checkPassword(pass) ; check {
			fmt.Println("Пароль принят")
			os.Exit(0)
		} else {
			fmt.Println("Пароль не соответствует требованиям: ", errmsg)
		}
	}

	fmt.Println("Превышено максимально доступное число попыток ввода пароля.")
}
