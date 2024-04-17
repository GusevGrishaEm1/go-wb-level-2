package customerror

import "fmt"

type CustomBusinessError struct {
	Message string
	Code    int
}

func (e *CustomBusinessError) Error() string {
	return fmt.Sprintf("Error: %s (Code: %d)", e.Message, e.Code)
}
