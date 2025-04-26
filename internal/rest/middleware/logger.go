package middleware

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"io"
)

func LoggerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				c.Logger().Error(err)
				return err
			}
			return nil
		}
	}
}

func RequestLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqBody, err := io.ReadAll(c.Request().Body)
			if err != nil {
				return err
			}

			logrus.WithFields(logrus.Fields{
				"method":  c.Request().Method,
				"uri":     c.Request().RequestURI,
				"headers": c.Request().Header,
				"body":    string(reqBody),
			}).Info("request :")

			c.Request().Body = io.NopCloser(bytes.NewBuffer(reqBody))

			err = next(c)
			if err != nil {
				c.Error(err)
			}

			resp := c.Response()
			logrus.WithFields(logrus.Fields{
				"status":  resp.Status,
				"headers": resp.Header(),
			}).Info("response : ")
			return nil
		}
	}
}
