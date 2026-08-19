/*
Сформировать данные для отправки заказа из
магазина по накладной и вывести на экран:
1) Наименование товара (минимум 1, максимум 100)
2) Количество (только числа)
3) ФИО покупателя (только буквы)
4) Контактный телефон (10 цифр)
5) Адрес: индекс(ровно 6 цифр), город, улица, дом, квартира

Эти данные не могут быть пустыми.
Проверить правильность заполнения полей.

Реализовать конструктор и несколько методов у типа "Накладная"

Пример:
invoice = NewInvoice()

или 

order = NewOrder()

*/

package main

import (
	"errors"
	"strings"
	"fmt"
	"bufio"
	"os"

)
 const chisla = "0123456789"
 const engBukvi = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
 const ruBukvi = "АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдежзийклмнопрстуфхцчшщъыьэюя"

// проверка числовой последовательности на только числа и длину
 func validateNumberString(str string, dlina int) bool {
	if len(str) != dlina {
		return false
	}

	for _, num := range str {
		if !strings.ContainsRune(chisla, num) {
			return false
		}
	}
	return true
  }

  // проверка фио
  func validateFio(fio string) bool {
	if len(fio) == 0 {
		return false
	}
	for _, symb := range fio {
		if symb == ' ' { //не забываем про пробел в фио
			continue
		}
		if !strings.ContainsRune(engBukvi, symb) && !strings.ContainsRune(ruBukvi, symb){
			return false
		}
	}
	return true
  }

  type Address struct {
	index string
	city  string
	street string
	house string
	apartment string
  }

  type Invoice struct {
	productname string
	quantity int
	fio string
	phone string
	Address
  }

  func NewInvoice() *Invoice {
	return &Invoice{}
  }

  // функция проверки вводимого имени продукта и добавления в структуру при успехе
  func (inv *Invoice) setProductName(productname string) error {
	if len(productname) < 1 || len(productname) >= 100 {
		return errors.New("Имя товара должно содержать не менее одного и не более 100 символов")
	}
	inv.productname = productname
	return nil
  }
  // функция проверки колиичества товара и добавления в структуру при успехе
  func (inv *Invoice) setQuantity(quantity int) error {
	if quantity <= 0 {
		return errors.New("количество товара должно быть больше 0")
	}
	inv.quantity = quantity
	return nil
  }
  //функция проверки фио и добавления в стурктуру при успехе
  func (inv *Invoice) setFio(fio string) error {
	if !validateFio(fio) {
		return errors.New("ФИО может содержать только русские/ланинские буквы и пробелы")
	}
	inv.fio = fio
	return nil
  }
  //функция проверки телефонного номера и добавления в структуру при успехе
  func (inv *Invoice) setPhoneNumber(phone string) error {
	if !validateNumberString(phone, 10) {
		return errors.New("телефон должен состоять из 10 цифр")
	}
	inv.phone = phone
	return nil
  }
  //функция проверки адресса и встраиванияв структуру при успехе
  func (inv *Invoice) setAddress(index, city, street, house, apartment string) error {
	if !validateNumberString(index, 6) {
		return errors.New("индекс должен содержать 6 цифр")
	}
	if city == "" || street == "" || house == "" || apartment == "" {
		return errors.New("поля адресса не должны быть пустыми")
	}
	inv.Address = Address{
		index: index,
		city: city,
		street: street,
		house: house,
		apartment: apartment,
	}
	return nil
  }
  //вывод информации из структуры
  func (inv *Invoice) printer() {
	fmt.Printf("1) Наименование товара: %s\n", inv.productname)
	fmt.Printf("2) Количество: %d\n", inv.quantity)
	fmt.Printf("3) ФИО покупателя: %s\n", inv.fio)
	fmt.Printf("4) Контактный телефон: %s\n", inv.phone)
	fmt.Printf("5) Адрес:\nИндекс: %s\nГород: %s\nУлица: %s\nДом: %s\nКвартира: %s\n", inv.Address.index, inv.Address.city, inv.Address.street, inv.Address.house, inv.Address.apartment)
	}
	//пишем ридер строк с bufio так как нам надо как-то считывать всю строку включая пробелы
	func stringReader(str string) string{
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(str)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("При чтении строки произошла ошибка", err)
		}
		return strings.TrimSpace(text)
	}

	func main() {
		var quant int
		invoice := NewInvoice()

		name := stringReader("Введите наименование товара: ")
		if err := invoice.setProductName(name); err != nil {
			fmt.Println("Ошибка: ", err)
			return
		}

		fmt.Print("Введите колличество товара: ")
		_, err := fmt.Scan(&quant)
		if err != nil {
			fmt.Println("При вводе колличества товара произошла ошибка: ", err)
		}
		if err := invoice.setQuantity(quant); err != nil {
			fmt.Println("Ошибка: ", err)
			return
		}

		fio := stringReader("Введите ФИО: ")
		if err:= invoice.setFio(fio); err != nil {
			fmt.Println("Ошибка: ", err)
			return
		}

		phone := stringReader("Введите контактый телефон: ")
		if err := invoice.setPhoneNumber(phone); err !=nil {
			fmt.Println("Ошибка: ", err)
			return
		}

		index := stringReader("Введите индекс: ")
		city := stringReader("Введите город: ")
		street := stringReader("Введите улицу:")
		house := stringReader("введите номер дома: ")
		apartment := stringReader("Введите номер квартиры: ")
		if err := invoice.setAddress(index, city, street, house, apartment); err != nil {
			fmt.Println("Ошибка: ")
			return
		}

		invoice.printer()

	}

// Ваш код