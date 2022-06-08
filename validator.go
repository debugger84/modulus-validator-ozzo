package validator

import (
	application "github.com/debugger84/modulus-application"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func AsAppValidationErrors(
	err error,
) []application.ValidationError {
	if err == nil {
		return nil
	}

	if errorSet, ok := err.(validation.Errors); ok {
		errors := make([]application.ValidationError, 0, len(errorSet))
		for key, val := range errorSet {
			if errorSet, ok := val.(validation.ErrorObject); ok {
				errors = append(
					errors, application.ValidationError{
						Field:      key,
						Err:        errorSet.Error(),
						Identifier: application.ErrorIdentifier(errorSet.Code()),
					},
				)
			}
		}

		return errors
	}
	return []application.ValidationError{
		{
			Field:      "",
			Identifier: application.UnknownError,
			Err:        err.Error(),
		},
	}
}
