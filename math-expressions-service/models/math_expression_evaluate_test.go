package models

import (
	"errors"
	"testing"

	"test-domain.com/math-expressions-service/models/expression_errors"
)

func TestMathExpressionEvaluate(t *testing.T) {
	tests := []struct {
		expression string
		result     float64
		err        error
	}{
		{
			expression: "What is 5?",
			result:     5,
			err:        nil,
		},
		{
			expression: "What is +5?",
			result:     5,
			err:        nil,
		},
		{
			expression: "What is -5?",
			result:     -5,
			err:        nil,
		},
		{
			expression: "WhAt iS 5?",
			result:     5,
			err:        nil,
		},
		{
			expression: " What   is   5 ? ",
			result:     5,
			err:        nil,
		},
		{
			expression: "What is 11 plus 5?",
			result:     16,
			err:        nil,
		},
		{
			expression: "What is 11 minus 50?",
			result:     -39,
			err:        nil,
		},
		{
			expression: "What is 2 multiplied by 50?",
			result:     100,
			err:        nil,
		},
		{
			expression: "What is 62 divided by 2?",
			result:     31,
			err:        nil,
		},
		{
			expression: "What is 61 divided by 2?",
			result:     30.5,
			err:        nil,
		},
		{
			expression: "What is -11 plus -5?",
			result:     -16,
			err:        nil,
		},
		{
			expression: "What is -11 minus -50?",
			result:     39,
			err:        nil,
		},
		{
			expression: "What is -2 multiplied by -50?",
			result:     100,
			err:        nil,
		},
		{
			expression: "What is 2 multiplied by -50?",
			result:     -100,
			err:        nil,
		},
		{
			expression: "What is -62 divided by -2?",
			result:     31,
			err:        nil,
		},
		{
			expression: "What is 5 plus 3 minus 1?",
			result:     7,
			err:        nil,
		},
		{
			expression: "What is 3 plus 2 multiplied by 3?",
			result:     15,
			err:        nil,
		},
		{
			expression: "What is 3 plus 7 divided by 5?",
			result:     2,
			err:        nil,
		},
		{
			expression: "What is 3 plus 7 divided by 5 minus 1?",
			result:     1,
			err:        nil,
		},
		{
			expression: "What is 30 minus 7 plus 10 divided by 11?",
			result:     3,
			err:        nil,
		},
		{
			expression: "What is 2 multiplied by 7 plus 6 divided by 10?",
			result:     2,
			err:        nil,
		},
		{
			expression: "What is 21 divided by -7 plus 6 plus 2?",
			result:     5,
			err:        nil,
		},
		{
			expression: "What is 61 divided by 0?",
			result:     0,
			err:        expression_errors.DivisionByZero{},
		},
		{
			expression: "",
			result:     0,
			err:        expression_errors.EmptyQuestion{},
		},
		{
			expression: "  ",
			result:     0,
			err:        expression_errors.EmptyQuestion{},
		},
		{
			expression: "What is?",
			result:     0,
			err:        expression_errors.EmptyQuestion{},
		},
		{
			expression: "What is 45 plus?",
			result:     0,
			err:        expression_errors.InvalidSyntax{},
		},
		{
			expression: "What is plus 5?",
			result:     0,
			err:        expression_errors.InvalidSyntax{},
		},
		{
			expression: "What is plus?",
			result:     0,
			err:        expression_errors.InvalidSyntax{},
		},
		{
			expression: "What is 5 plus minus 3?",
			result:     0,
			err:        expression_errors.InvalidSyntax{},
		},
		{
			expression: "What is  plus -3 minus 4 multiplied by 100 ?",
			result:     0,
			err:        expression_errors.InvalidSyntax{},
		},
		{
			expression: "What is 3 cubed?",
			result:     0,
			err:        expression_errors.UnsupportedOperation{},
		},
		{
			expression: "What is sine value of 3?",
			result:     0,
			err:        expression_errors.UnsupportedOperation{},
		},
		{
			expression: "What is logarithm of 16 with a base 4?",
			result:     0,
			err:        expression_errors.UnsupportedOperation{},
		},
		{
			expression: "What is cat?",
			result:     0,
			err:        expression_errors.NonMathQuestion{},
		},
		{
			expression: "Who is the President of the United States?",
			result:     0,
			err:        expression_errors.NonMathQuestion{},
		},
		{
			expression: "Is this a dog",
			result:     0,
			err:        expression_errors.NonMathQuestion{},
		},
		// ... add more test cases
	}

	for _, test := range tests {
		me := MathExpression{}
		me.SetExpression(test.expression)
		result, err := me.Evaluate()

		if result != test.result || !errors.Is(err, test.err) {
			t.Errorf("Evaluate(%q) = %v, %v; want %v, %v", test.expression, result, err, test.result, test.err)
		}
	}
}
