package utils

import (
	"checkout-service/app/config"
	"checkout-service/pkg/constant"
	"github.com/labstack/echo/v4"
	"strings"
)

func JsonOK(c echo.Context, data interface{}) error {
	return JsonCors(c, 200, data)
}

func JsonErr(c echo.Context, code int, data interface{}) error {
	return JsonCors(c, code, data)
}

func JsonCors(c echo.Context, code int, data interface{}) error {
	c.Response().Header().Set(constant.ACAM, strings.Join(config.GetEnvCors("CORS_METHOD_ALLOWED"), ", "))
	c.Response().Header().Set(constant.ACAH, constant.ACAH_VALUE)
	c.Response().Header().Set(constant.ACAC, constant.ACAC_VALUE)
	c.Response().Header().Set(constant.CC, constant.CC_VALUE)
	c.Response().Header().Set(constant.XCTO, constant.XCTO_VALUE)
	c.Response().Header().Set("Content-Security-Policy", "default-src 'self'")
	return c.JSON(code, data)
}
