package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Карта с нашими римскими числами
var letters = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}
var lettersMore = map[string]int{
	"X":    1,
	"XX":   2,
	"XXX":  3,
	"XL":   4,
	"L":    5,
	"LX":   6,
	"LXX":  7,
	"LXXX": 8,
	"XC":   9,
}

// В данной функции мы определям наши дальнешие действия и получаем срез строковый
func sign(s string) ([]string, string) {
	slice := strings.Split(s, "+")
	action := ""
	if len(slice) == 1 {
		slice = strings.Split(s, "-")
		if len(slice) == 1 {
			slice = strings.Split(s, "/")
			if len(slice) == 1 {
				slice = strings.Split(s, "*")
				if len(slice) == 1 {
					err := fmt.Errorf("Неверное действие!")
					panic(err)
				} else {
					action = "*"
				}
			} else {
				action = "/"
			}
		} else {
			action = "-"
		}

	} else {
		action = "+"
	}
	//fmt.Println(slice, action)
	return slice, action
}

// Функция для чтения с клавиатуры
func keyboard() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal()
	}
	input = strings.TrimSpace(input)
	return input
}

// Функция для перевода римских чисел в арабские
func toArabic(number string) int {
	var arabic int
	for key, value := range letters {
		if number == key {
			arabic = value
		}
	}
	if arabic != 0 {
		return arabic
	} else {
		panic(fmt.Errorf("Что-то не то с цифрами!"))
	}
}

// функция для перевода араских чисел в римские для ответа в действиях с римскими числами
func toRom(i int) string {
	var rom string
	if i <= 10 {
		for key, value := range letters {
			if i == value {
				rom = key
			}
		}
	} else if i > 10 && i < 100 {
		sl := strings.Split(strconv.Itoa(i), "")
		for key, value := range lettersMore {
			if sl[0] == strconv.Itoa(value) {
				rom = key
				for key, value = range letters {
					if sl[1] == strconv.Itoa(value) {
						rom = rom + key
					}
				}
			}
		}
	} else if i == 100 {
		rom = "C"
	}
	return rom
}

// Проверка является ли число целым
func isInt(i float64) bool {
	if sub := math.Round(i) - i; sub == 0 {
		return true
	} else {
		return false
	}
}
func main() {
	fmt.Println("Введите выражение:")
	line := keyboard()
	slice, action := sign(line)
	if len(slice) > 2 { //Проверяем, что будем обрабатывать всего дно действие
		err := fmt.Errorf("Калькулятор может выполнять только 1 действие!")
		panic(err)
	}
	arabic := []bool{} //Создаем срез со значениями ложь/истина, чтобы проверить, что оба числа одинаковые
	expression := make([]int, 0)
	for _, value := range slice { //В цикле создаем срез со значениями, которые далее будем использовать при расчетах
		if n, err := strconv.ParseFloat(value, 64); err == nil {
			if n >= 1 && n <= 10 && isInt(n) {
				num := int(n)
				//fmt.Println(value)
				expression = append(expression, num)
				arabic = append(arabic, true)
			} else {
				err = fmt.Errorf("Числа не входят в диапозон!")
				panic(err)
			}
		} else {
			num := toArabic(value)

			expression = append(expression, num)
			arabic = append(arabic, false)
		}
	}
	if arabic[0] != arabic[1] { //Проверяем, что оба числа совпадают в знаках
		err := fmt.Errorf("Числа должны быть либо арабскими, либо римскими и только!")
		panic(err)
	}
	result := 0
	switch action { //Производим расчеты в зависимости от действия
	case "+":
		result = add(expression)
	case "-":
		result = deg(expression)
	case "/":
		result = div(expression)
	case "*":
		result = mult(expression)
	}

	if arabic[0] != true {

		resultRom := toRom(result) //Переводим ответ в римские числа

		if resultRom == "" {
			fmt.Println("Ответ с римскими цифрами не должен быть отрицательным и = 0, прости!")
		} else {
			fmt.Println("Ответ:", resultRom)
		}
	} else {
		fmt.Println("Ответ:", result)
	}
}

// Функции для расчетов
func add(i []int) int {

	return i[0] + i[1]
}
func deg(i []int) int {
	return i[0] - i[1]
}
func div(i []int) int {
	return i[0] / i[1]
}
func mult(i []int) int {
	return i[0] * i[1]
}
