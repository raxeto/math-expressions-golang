package expression_errors

type InvalidSyntax struct{}

// Error implements the error interface for InvalidSyntax
func (e InvalidSyntax) Error() string {
	return "Invalid syntax"
}
