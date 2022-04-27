package semantics

import (
	"fmt"
	"strconv"
	"strings"
)

func (s *SemanticAnalyzer) parseString(str string) string {
	// turn string into array
	arr := strings.Split(str, " ")
	chomped_arr := []string{}
	// print line by line to check how spaces are handled
	for _, v := range arr {
		// if the string is blank, skip it
		if v == "" {
			continue
		} else {
			chomped_arr = append(chomped_arr, v)
		}

	}
	// print chomped array
	// check index 1 for operator
	op1, _ := strconv.Atoi(chomped_arr[0])
	op2, _ := strconv.Atoi(chomped_arr[2])
	switch chomped_arr[1] {
	case "+":
		res := op1 + op2
		r := fmt.Sprintf("%d", res)
		// eliminate the 3 first elements of the array
		chomped_arr = chomped_arr[3:]
		// add the result to the array at the beginning
		chomped_arr = append([]string{r}, chomped_arr...)

		if len(chomped_arr) == 1 {
			return chomped_arr[0]
		}
		return s.parseString(strings.Join(chomped_arr, " "))
	case "-":
		res := op1 - op2
		r := fmt.Sprintf("%d", res)
		chomped_arr = chomped_arr[3:]
		// add the result to the array at the beginning
		chomped_arr = append([]string{r}, chomped_arr...)

		if len(chomped_arr) == 1 {
			return chomped_arr[0]
		}
		return s.parseString(strings.Join(chomped_arr, " "))
	case "*":
		res := op1 * op2
		r := fmt.Sprintf("%d", res)
		chomped_arr = chomped_arr[3:]
		// add the result to the array at the beginning
		chomped_arr = append([]string{r}, chomped_arr...)

		if len(chomped_arr) == 1 {
			return chomped_arr[0]
		}
		return s.parseString(strings.Join(chomped_arr, " "))
	case "/":
		res := op1 / op2
		r := fmt.Sprintf("%d", res)
		chomped_arr = chomped_arr[3:]
		// add the result to the array at the beginning
		chomped_arr = append([]string{r}, chomped_arr...)

		if len(chomped_arr) == 1 {
			return chomped_arr[0]
		}
		return s.parseString(strings.Join(chomped_arr, " "))
	}

	return ""
}
