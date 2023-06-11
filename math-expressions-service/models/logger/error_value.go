package logger

import (
	"sync"
)

type errorValue struct {
	ExoressionError error
	Frequency       uint
	mutex           sync.Mutex
}
