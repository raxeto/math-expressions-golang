package operations

// IMathOperation defines the interface for math operations
type IMathOperation interface {
	Eval(a, b float64) (float64, error)
}
