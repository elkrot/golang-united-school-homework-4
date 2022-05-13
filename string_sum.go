package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	if strings.TrimSpace(input) == "" {
		return "", fmt.Errorf("Input parameter error: %w", errorEmptyInput)
	}

	sum := 0
	sign := "+"
	operand_count := 0
	strv := ""
	fire := false
	for i, v := range input {
		if !(string(v) == "+" || string(v) == "-" || string(v) == " " || i == len(input)-1) {
			strv += string(v)
		} else {

			if strv != "" || i == len(input)-1 {
				if i == len(input)-1 {
					strv += string(v)
				}
				n, err := strconv.Atoi(strings.TrimSpace(strv))

				if err == nil {
					if sign == "-" {
						sum -= n
					} else {
						sum += n
					}
					strv = ""
					fire = true
					operand_count++
				} else {
					return "", fmt.Errorf("conversion process failed with error : %w", err)
				}

			}
			if string(v) == "-" {
				sign = "-"
				fire = false
			} else {
				if (fire == true && string(v) == " ") || string(v) == "+" {
					sign = "+"
				}
			}
		}
	}

	if operand_count != 2 {
		return "", fmt.Errorf("Input parameter error: %w", errorNotTwoOperands)
	}

	return strconv.Itoa(sum), nil
}
