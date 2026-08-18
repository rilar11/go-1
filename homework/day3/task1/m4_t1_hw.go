/* 
	Задача №1
	Написать функцию, которая расшифрует строку. 
	code = "220411112603141304" Out: well done
	Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
	Отчет с "00" -> 'a' и до "25" -> 'z', "26" -> ' '(пробел).
	Вход: строка из цифр. Выход: Текст.
	Проверка работы функции выполняется через другую строку.
	Рекомендую использовать мапы, будет лучше, если вы их создадите с помощью цикла

	Задача №1.1
	Реализовать и функцию зашифровки

	codeToString(code) -> "???????' 

	stringToCode("hello") -> "??????"
*/
package main

import (
	"fmt"
	"os"
	"strings"
)
//функция декодирования числовой последовательности
func decodeString(codeStr string, symbolMap map[string]rune) (string, error) {

	var sb strings.Builder
	for i := 0; i < len(codeStr); i += 2 { // +2 так как каждый код 2 цифры
		codePair := codeStr[i : i+2]
		symbol, pass := symbolMap[codePair]
		if !pass {
			return "", fmt.Errorf("не найден символ для кода: %s", codePair)
		}
		sb.WriteRune(symbol)
	}
	return sb.String(), nil
}
//функция кодирования строки
func codeString(text string, codeMap map[rune]string) (string, error) {
	var sb strings.Builder
	for _, char := range text {
		code, pass := codeMap[char]
		if !pass {
			return "", fmt.Errorf("не найден код для символа: %c", char)
		}
		sb.WriteString(code)
	}
	return sb.String(), nil
}

//генерим мапки
func generateMaps() (map[string]rune, map[rune]string) {
	symbolMap := make(map[string]rune)
	codeMap := make(map[rune]string)

	// проходимся по символам a-z
	for i := 0; i < 26; i++ {
		code := fmt.Sprintf("%02d", i) // i = 0 code = 00, i = 1, code = 01
		symbol := rune('a' + i) // i = 0 symbol = 'a' + 0 = 'a', i = 1 symbol = 'a' + 1 = 'b'
		// ну и наполняем мапки
		codeMap[symbol] = code
		symbolMap[code] = symbol
	}

	// добавляем в мапы пробел
	codeMap[' '] = "26"
	symbolMap["26"] = ' '

	return symbolMap, codeMap
}

func main() {
	var selector, code, text string
	symbolMap, codeMap := generateMaps()

	fmt.Println("Выберите действие:")
	fmt.Println("1 - Декодировать последовательность")
	fmt.Println("2 - Кодировать строку")
	fmt.Print("Введите 1 или 2: ")

	_, err := fmt.Scan(&selector)
	if err != nil {
		fmt.Println("Ошибка при выборе действия:", err)
		os.Exit(1)
	}

	switch selector {
	case "1":
		fmt.Print("Введите последовательность для декодирования: ")
		_, err = fmt.Scan(&code)
		if err != nil {
			fmt.Println("Ошибка ввода последовательности: ", err)
			os.Exit(1)
		}

		decodedString, err := decodeString(code, symbolMap)
		if err != nil {
			fmt.Println("Ошибка декодирования. ", err)
			os.Exit(1)
		}
		fmt.Println("Декодированная строка: ", decodedString)

	case "2":
		fmt.Print("Введите текст (строчные буквы a-z и пробелы): ")
		_, err = fmt.Scanln(&text)
		if err != nil {
			fmt.Println("Ошибка ввода текста: ", err)
			os.Exit(1)
		}

		encodedString, err := codeString(text, codeMap)
		if err != nil {
			fmt.Println("Ошибка кодирования.", err)
			os.Exit(1)
		}
		fmt.Println("Кодированная последовательность: ", encodedString)

	default:
		fmt.Println("Неверный выбор. В следующий раз введите 1 или 2.")
		os.Exit(1)
		
	}
}
