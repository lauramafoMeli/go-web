package tools

import "fmt"

type FieldError struct {
	Field string
	Msg   string
}

func (e *FieldError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

func ValidateProductFields(fields map[string]any, requariedFields ...string) (err error) {
	for _, field := range requariedFields {
		if _, ok := fields[field]; !ok || fields[field] == "" {
			return &FieldError{Field: field, Msg: "is required"}
		}
	}
	return nil
}
