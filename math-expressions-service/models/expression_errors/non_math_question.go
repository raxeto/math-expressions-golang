package expression_errors

type NonMathQuestion struct{}

// Error implements the error interface for NonMathQuestion
func (e NonMathQuestion) Error() string {
	return "Non-math question"
}
