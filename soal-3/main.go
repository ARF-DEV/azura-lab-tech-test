package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func kalkulator(input string) int {
	splittedVal := strings.Split(input, " ")

	if len(splittedVal) != 3 {
		panic("Input is invalid")
	}

	num1, err := strconv.Atoi(splittedVal[0])

	if err != nil {
		panic(err.Error())
	}

	num2, err := strconv.Atoi(splittedVal[2])
	op := splittedVal[1]

	if err != nil {
		panic(err.Error())
	}

	if num1 > 1000000 || num2 > 1000000 {
		panic("the number that's used cannot exceds 1000000")
	}

	result := 0
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "/":
		result = num1 / num2
	case "*":
		result = num1 * num2
	default:
		panic("can only except these four operation (+, -, *, /)")
	}

	return result
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		panic("error on input")
	}

	input = strings.TrimSuffix(input, "\n")

	fmt.Println(kalkulator(input))
}
