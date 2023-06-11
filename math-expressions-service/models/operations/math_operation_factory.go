package operations

import "strings"

type MathOperationFactory struct{}

func (mof MathOperationFactory) AllowedOperations() []string {
	return []string{"plus", "minus", "multiplied by", "divided by"}
}

func (mof MathOperationFactory) Create(operation string) IMathOperation {
	operation = strings.ToLower(operation)

	switch operation {
	case "plus":
		return addition{}
	case "minus":
		return subtraction{}
	case "multiplied by":
		return multiplication{}
	case "divided by":
		return division{}
	default:
		return nil
	}
}
