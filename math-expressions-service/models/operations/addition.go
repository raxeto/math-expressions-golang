package operations

// Addition structure
type addition struct{}

// Eval method for Addition
func (a addition) Eval(op1, op2 float64) (float64, error) {
	return op1 + op2, nil
}
