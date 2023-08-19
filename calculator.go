package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func precedence(operator rune) int {
	switch operator {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}

func shuntingYard(tokens []string) []string {
	var output []string
	var operators []rune

	for _, token := range tokens {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			output = append(output, strconv.FormatFloat(num, 'f', -1, 64))
		} else if token == "(" {
			operators = append(operators, '(')
		} else if token == ")" {
			for len(operators) > 0 && operators[len(operators)-1] != '(' {
				output = append(output, string(operators[len(operators)-1]))
				operators = operators[:len(operators)-1]
			}
			operators = operators[:len(operators)-1] // Pop '('
		} else {
			for len(operators) > 0 && precedence(rune(token[0])) <= precedence(operators[len(operators)-1]) {
				output = append(output, string(operators[len(operators)-1]))
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, rune(token[0]))
		}
	}

	for len(operators) > 0 {
		output = append(output, string(operators[len(operators)-1]))
		operators = operators[:len(operators)-1]
	}

	return output
}

func evaluateRPN(tokens []string) float64 {
	stack := []float64{}

	for _, token := range tokens {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, num)
		} else {
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2] // Pop operands
			var result float64

			switch token {
			case "+":
				result = operand1 + operand2
			case "-":
				result = operand1 - operand2
			case "*":
				result = operand1 * operand2
			case "/":
				result = operand1 / operand2
			}

			stack = append(stack, result)
		}
	}

	return stack[0]
}

func main() {
	fmt.Println("Simple Calculator in Go")
	fmt.Println("Enter an arithmetic expression (e.g., 2 + 3 * ( 4 - 1 ) ):")

	reader := bufio.NewReader(os.Stdin)
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	tokens := strings.Fields(expression)
	rpnTokens := shuntingYard(tokens)
	result := evaluateRPN(rpnTokens)

	fmt.Printf("Result: %.2f\n", result)
}
