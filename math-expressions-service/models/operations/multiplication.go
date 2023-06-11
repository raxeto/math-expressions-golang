package operations

// Subtraction structure
type multiplication struct{}

// Eval method for multiplication
func (d multiplication) Eval(op1, op2 float64) (float64, error) {
	return op1 * op2, nil
}
