package logger

import (
	"sync"

	"test-domain.com/math-expressions-service/models/dto"
)

type mathExpressionLogger struct {
	errors sync.Map
}

var (
	instance *mathExpressionLogger
	once     sync.Once
)

func GetInstance() *mathExpressionLogger {
	once.Do(func() {
		instance = &mathExpressionLogger{}
	})
	return instance
}

func (logger *mathExpressionLogger) Error(expression string, endpoint string, expressionError error) {
	key := errorKey{Expression: expression, Endpoint: endpoint}
	actual, loaded := logger.errors.LoadOrStore(key, &errorValue{ExoressionError: expressionError, Frequency: 1})

	if loaded {
		errorValue := actual.(*errorValue)
		errorValue.mutex.Lock()
		errorValue.Frequency++
		errorValue.mutex.Unlock()
	}
}

func (logger *mathExpressionLogger) GetAllErrors() []dto.ErrorJson {
	errorsJson := []dto.ErrorJson{}

	logger.errors.Range(func(k, v interface{}) bool {
		key := k.(errorKey)
		value := v.(*errorValue)

		errorJson := dto.ErrorJson{
			Expression: key.Expression,
			Endpoint:   key.Endpoint,
			Frequency:  value.Frequency,
			Type:       value.ExoressionError.Error(),
		}

		errorsJson = append(errorsJson, errorJson)
		return true
	})

	return errorsJson
}
