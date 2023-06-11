package expression_errors

type DivisionByZero struct{}

// Error implements the error interface for DivisionByZero
func (e DivisionByZero) Error() string {
	return "Division by zero"
}
