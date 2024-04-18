// Package customerror предоставляет пользовательский тип ошибки для обработки бизнес-ошибок.
package customerror

import "fmt"

// CustomBusinessError представляет пользовательский тип ошибки для бизнес-ошибок.
type CustomBusinessError struct {
	// Message - сообщение об ошибке.
	Message string
	// Code - код ошибки.
	Code int
}

// Error возвращает сообщение и код ошибки, отформатированные как строка.
func (e *CustomBusinessError) Error() string {
	return fmt.Sprintf("Ошибка: %s (Код: %d)", e.Message, e.Code)
}
