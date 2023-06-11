package operations

import (
	"test-domain.com/math-expressions-service/models/expression_errors"
)

// Subtraction structure
type division struct{}

// Eval method for division
func (d division) Eval(op1, op2 float64) (float64, error) {
	if op2 == 0 {
		return 0, expression_errors.DivisionByZero{}
	}
	return op1 / op2, nil
}
