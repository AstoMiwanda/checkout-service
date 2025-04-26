package middleware

import "github.com/labstack/echo/v4"

type customValidator struct{}

func NewCustomValidator() echo.Validator {
	return &customValidator{}
}

func (c customValidator) Validate(i interface{}) error {
	return nil
}
