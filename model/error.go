package model

import "fmt"

type DatabaseError struct {
	Detail string
}

func (dbe *DatabaseError) Error() string {
	return dbe.Detail
}

type ValidationError struct {
	Field   string
	Message string
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("%s:%s", ve.Field, ve.Message)
}

func NewValidationError(field, msg string) error {
	return &ValidationError{field, msg}
}
