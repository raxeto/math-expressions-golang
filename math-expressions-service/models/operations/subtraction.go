package operations

// Subtraction structure
type subtraction struct{}

// Eval method for Subtraction
func (s subtraction) Eval(op1, op2 float64) (float64, error) {
	return op1 - op2, nil
}
