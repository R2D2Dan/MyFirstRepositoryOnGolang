package main

import (
	operation "Calculation/Operation"
	"bufio"
	"fmt"
	"os"
	"strconv"
	s "strings"

	textcolor "github.com/fatih/color"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	fmt.Printf("Знаки для действий:\n\t+ : Сложить\n\t- : Вычитание\n\t* : Умножение\n\t/ : Деление\n")
	fmt.Print("Введите 2 числа и знак между ними:")
	sc.Scan()
	test := s.ReplaceAll(sc.Text(), " ", "")
	a, err := strconv.ParseFloat(string(test[0]), 64)

	if err != nil {
		fmt.Println("Ошибка первого числа")
		return
	}
	b, err := strconv.ParseFloat(string(test[2]), 64)
	if err != nil {
		fmt.Println("Ошибка второго чилса")
		return
	}

	operatin := string(test[1])
	var res float64
	switch operatin {
	case "+":
		res = operation.Pluse(a, b)
	case "-":
		res = operation.Minus(a, b)
	case "*":
		res = operation.Multiplication(a, b)
	case "/":
		res, err = operation.Devision(a, b)
		if err != nil {
			fmt.Println(err)
		}

	}
	textcolor.Yellow(fmt.Sprintf("Результат: %.2f", res))
	textcolor.Red(fmt.Sprintf("Результат: %.2f", res))
	textcolor.Green(fmt.Sprintf("Результат: %.2f", res))
	textcolor.Blue(fmt.Sprintf("Результат: %.2f", res))
	textcolor.Cyan(fmt.Sprintf("Результат: %.2f", res))
	textcolor.Cyan(fmt.Sprintf("Результат: %.2f", res))
	textcolor.Cyan(fmt.Sprintf("Результат: %.2f", res))
	textcolor.Cyan(fmt.Sprintf("Результат: %.2f", res))
	textcolor.Cyan(fmt.Sprintf("Результат: %.2f", res))

}
