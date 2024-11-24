package main

import (
	"fmt"
	"strings"
)

func parseBoolExpr(expression string) bool {
	for strings.Contains(expression, "(") {
		var builder strings.Builder
		openingParenthesis := strings.LastIndex(expression, "(")
		subExpression := expression[openingParenthesis:]
		closingParenthesis := strings.Index(subExpression, ")") + openingParenthesis

		operator := expression[openingParenthesis-1 : openingParenthesis]
		innerExpression := expression[openingParenthesis+1 : closingParenthesis]

		evaluatedExpression := calculate(operator, innerExpression)

		firstHalfExp := expression[:openingParenthesis-1]
		var evaluatedHalfExp string

		if evaluatedExpression {
			evaluatedHalfExp = "t"
		} else {
			evaluatedHalfExp = "f"
		}

		tailHalfExp := expression[closingParenthesis+1:]

		// expression = firstHalfExp + evaluatedHalfExp + tailHalfExp

		// Is more efficient using builder.
		builder.WriteString(firstHalfExp)
		builder.WriteString(evaluatedHalfExp)
		builder.WriteString(tailHalfExp)
		expression = builder.String()
		fmt.Println(expression)
	}

	return expression == "t"
}

func isTrue(s string) bool {
	return s == "t"
}

func calculate(operator, exp string) bool {
	switch operator {
	case "&":
		return every(strings.Split(exp, ","), isTrue)
	case "|":
		return any(strings.Split(exp, ","), isTrue)
	case "!":
		return exp == "f"
	}

	return false
}

func every(slice []string, predicate func(string) bool) bool {
	for _, v := range slice {
		if !predicate(v) {
			return false
		}
	}

	return true
}

func any(slice []string, predicate func(string) bool) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}

	return false
}

func main() {
	// expression := "|(f,f,f,t)"
	expression := "&(|(f))"
	fmt.Printf("Hello, Go! %v\n", parseBoolExpr(expression))

}
