package models

import (
	"errors"
	"testing"

	"test-domain.com/math-expressions-service/models/expression_errors"
)

func TestMathExpressionValidate(t *testing.T) {
	tests := []struct {
		expression string
		valid      bool
		err        error
	}{
		{
			expression: "What is 5?",
			valid:      true,
			err:        nil,
		},
		{
			expression: "What is +5?",
			valid:      true,
			err:        nil,
		},
		{
			expression: "What is -5?",
			valid:      true,
			err:        nil,
		},
		{
			expression: "WhAt iS 5?",
			valid:      true,
			err:        nil,
		},
		{
			expression: " What   is   5 ? ",
			valid:      true,
			err:        nil,
		},
		{
			expression: "What is 11 plus 5?",
			valid:      true,
			err:        nil,
		},
		{
			expression: "What is 11 minus 50?",
			valid:      true,
			err:        nil,
		},
		{
			expression: "What is 2 multiplied by 50?",
			valid:      true,
			err:        nil,
		},
		{
			expression: "What is 62 divided by 2?",
			valid:      true,
			err:        nil,
		},
		{
			expression: "What is -11 plus -5?",
			valid:      true,
			err:        nil,
		},
		{
			expression: "What is -11 minus -50?",
			valid:      true,
			err:        nil,
		},
		{
			expression: "What is -2 multiplied by -50?",
			valid:      true,
			err:        nil,
		},
		{
			expression: "What is 2 multiplied by -50?",
			valid:      true,
			err:        nil,
		},
		{
			expression: "What is -62 divided by -2?",
			valid:      true,
			err:        nil,
		},
		{
			expression: "What is 62 divided by -2?",
			valid:      true,
			err:        nil,
		},
		{
			expression: "What is 5 plus 3 minus 1?",
			valid:      true,
			err:        nil,
		},
		{
			expression: "What is 3 plus 2 multiplied by 3?",
			valid:      true,
			err:        nil,
		},
		{
			expression: "What is 3 plus 7 divided by 5?",
			valid:      true,
			err:        nil,
		},
		{
			expression: "What is 3 plus 7 divided by 5 minus 1?",
			valid:      true,
			err:        nil,
		},
		{
			expression: "What is 30 minus 7 plus 10 divided by 11?",
			valid:      true,
			err:        nil,
		},
		{
			expression: "What is 2 multiplied by 7 plus 6 divided by 10?",
			valid:      true,
			err:        nil,
		},
		{
			expression: "What is 21 divided by -7 plus 6 plus 2?",
			valid:      true,
			err:        nil,
		},
		{
			expression: "",
			valid:      false,
			err:        expression_errors.EmptyQuestion{},
		},
		{
			expression: "  ",
			valid:      false,
			err:        expression_errors.EmptyQuestion{},
		},
		{
			expression: "What is?",
			valid:      false,
			err:        expression_errors.EmptyQuestion{},
		},
		{
			expression: "What is 45 plus?",
			valid:      false,
			err:        expression_errors.InvalidSyntax{},
		},
		{
			expression: "What is plus 5?",
			valid:      false,
			err:        expression_errors.InvalidSyntax{},
		},
		{
			expression: "What is plus?",
			valid:      false,
			err:        expression_errors.InvalidSyntax{},
		},
		{
			expression: "What is 5 plus minus 3?",
			valid:      false,
			err:        expression_errors.InvalidSyntax{},
		},
		{
			expression: "What is  plus -3 minus 4 multiplied by 100 ?",
			valid:      false,
			err:        expression_errors.InvalidSyntax{},
		},
		{
			expression: "What is 3 cubed?",
			valid:      false,
			err:        expression_errors.UnsupportedOperation{},
		},
		{
			expression: "What is sine value of 3?",
			valid:      false,
			err:        expression_errors.UnsupportedOperation{},
		},
		{
			expression: "What is logarithm of 16 with a base 4?",
			valid:      false,
			err:        expression_errors.UnsupportedOperation{},
		},
		{
			expression: "What is cat?",
			valid:      false,
			err:        expression_errors.NonMathQuestion{},
		},
		{
			expression: "Who is the President of the United States?",
			valid:      false,
			err:        expression_errors.NonMathQuestion{},
		},
		{
			expression: "Is this a dog?",
			valid:      false,
			err:        expression_errors.NonMathQuestion{},
		},
		// ... add more test cases
	}

	for _, test := range tests {
		me := MathExpression{}
		me.SetExpression(test.expression)
		valid, err := me.Validate()

		if valid != test.valid || !errors.Is(err, test.err) {
			t.Errorf("Validate(%q) = %v, %v; want %v, %v", test.expression, valid, err, test.valid, test.err)
		}
	}
}
