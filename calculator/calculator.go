package calculator

import (
	"Practica3/parser"
	"fmt"
	"strconv"
)

func Resolve(tokens []parser.Token) (string, error) {
	var res string

	// the token array has a postfix notation, we have to resolve it
	// we use a stack to do this
	stack := make([]string, 0)

	for _, token := range tokens {
		switch token.Type {
		case "number":
			fmt.Println("Number: ", token.Info)
			stack = append(stack, token.Info)
		case "operator":
			fmt.Println("Operator: ", token.Info)
			// pop the last two elements from the stack
			// and apply the operator
			// push the result back to the stack
			right := stack[len(stack)-1]
			fmt.Println("Right: ", right)
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			fmt.Println("Left: ", left)
			stack = stack[:len(stack)-1]
			fmt.Println("Stack: ", stack)
			result := applyOperator(left, right, token.Info)
			stack = append(stack, result)
		default:
		}
	}
	// the stack should have only one element
	if len(stack) == 1 {
		res = stack[0]
	} else {
		return "", fmt.Errorf("error on resolve module")
	}
	return res, nil

}

func applyOperator(left string, right string, operator string) string {
	var res string
	switch operator {
	case "+":
		// we have to convert the strings to ints
		// and add them
		intLeft, _ := strconv.Atoi(left)
		intRight, _ := strconv.Atoi(right)
		res = strconv.Itoa(intLeft + intRight)
	case "-":
		intLeft, _ := strconv.Atoi(left)
		intRight, _ := strconv.Atoi(right)
		res = strconv.Itoa(intLeft - intRight)
	case "*":
		intLeft, _ := strconv.Atoi(left)
		intRight, _ := strconv.Atoi(right)
		res = strconv.Itoa(intLeft * intRight)
	case "/":
		intLeft, _ := strconv.Atoi(left)
		intRight, _ := strconv.Atoi(right)
		res = strconv.Itoa(intLeft / intRight)
	}
	return res
}
