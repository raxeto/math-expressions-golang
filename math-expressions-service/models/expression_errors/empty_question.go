package expression_errors

type EmptyQuestion struct{}

// Error implements the error interface for EmptyQuestion
func (e EmptyQuestion) Error() string {
	return "Empty question"
}
