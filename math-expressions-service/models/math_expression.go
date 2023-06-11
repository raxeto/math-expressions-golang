package models

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"test-domain.com/math-expressions-service/models/expression_errors"
	"test-domain.com/math-expressions-service/models/operations"
)

type MathExpression struct {
	expression string
}

// SetExpression trims leading and trailing whitespaces, and replaces multiple whitespaces with a single space
func (me *MathExpression) SetExpression(expression string) {
	expression = strings.TrimSpace(expression)

	// Regular expression to match consecutive whitespaces
	re := regexp.MustCompile(`\s+`)
	expression = re.ReplaceAllString(expression, " ")

	me.expression = expression
}

func (me *MathExpression) GetExpression() string {
	return me.expression
}

func (m *MathExpression) Evaluate() (float64, error) {
	validateRes, err := m.Validate()

	if !validateRes {
		if err == nil {
			err = errors.New("math expression error")
		}
		return 0, err
	}

	parts := regexp.MustCompile(`\s+`).Split(strings.TrimRight(m.expression, "?"), -1)

	result, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, err
	}

	factory := operations.MathOperationFactory{}
	for i := 3; i < len(parts)-1; i += 2 {
		operation := parts[i]

		operand, err := strconv.ParseFloat(parts[i+1], 64)
		for err != nil {
			operation += " " + parts[i+1]
			i++
			operand, err = strconv.ParseFloat(parts[i+1], 64)
		}

		mathOperation := factory.Create(operation)
		if mathOperation != nil {
			evalRes, err := mathOperation.Eval(result, operand)
			if err != nil {
				return 0, err
			}
			result = evalRes
		}
	}

	return result, nil

}

func (m *MathExpression) Validate() (bool, error) {
	factory := operations.MathOperationFactory{}
	allowedOperations := factory.AllowedOperations()

	operations := `\s+` + strings.Join(allowedOperations, `\s+|\s+`) + `\s+`
	validPattern := fmt.Sprintf(`(?i)^What is ([-+]?\d+)((%s)([-+]?\d+))*\s*\?$`, operations)

	if regexp.MustCompile(validPattern).MatchString(m.expression) {
		return true, nil
	}

	emptyQuestionPattern := `(?i)^What is\s*\?*$`

	if len(strings.TrimSpace(m.expression)) == 0 || regexp.MustCompile(emptyQuestionPattern).MatchString(m.expression) {
		return false, expression_errors.EmptyQuestion{}
	}

	// Invalid syntax - What is [(number)0-n (plus|minus|multiplied by|divided by)0-n (number)0-n]1-n?
	operations = strings.Join(allowedOperations, "|")
	invalidSyntaxPattern := fmt.Sprintf(`(?i)^What is\s+(([-+]?\d+)*\s*(%s)*\s*([-+]?\d+)*)+\s*\?$`, operations)

	if regexp.MustCompile(invalidSyntaxPattern).MatchString(m.expression) {
		return false, expression_errors.InvalidSyntax{}
	}

	// Unsupported operation - What is [(word)0-n number (word)0-n]1-n?
	// (\b\w+\b\s*)* - matches 0 or more words
	unsupportedOperPattern := `(?i)^What is ((\b\w+\b\s*)*[-+]?\d+\s*(\b\w+\b\s*)*)+\?$`

	if regexp.MustCompile(unsupportedOperPattern).MatchString(m.expression) {
		return false, expression_errors.UnsupportedOperation{}
	}

	return false, expression_errors.NonMathQuestion{}
}
