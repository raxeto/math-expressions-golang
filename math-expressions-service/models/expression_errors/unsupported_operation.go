package expression_errors

type UnsupportedOperation struct{}

// Error implements the error interface for UnsupportedOperation
func (e UnsupportedOperation) Error() string {
	return "Unsupported operation"
}
