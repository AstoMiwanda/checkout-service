package cmd

import (
	"os"

	"checkout-service/app/config"
	restHandler "checkout-service/internal/rest"
	"checkout-service/internal/rest/middleware"
	"github.com/labstack/echo/v4"
	middleware2 "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var restCommand = &cobra.Command{
	Use:   "rest",
	Short: "Start REST server",
	Run:   restServer,
}

func init() {
	rootCmd.AddCommand(restCommand)
}

func restServer(cmd *cobra.Command, args []string) {
	// Start REST server
	e := echo.New()
	e.Use(middleware.LoggerMiddleware())
	e.Validator = middleware.NewCustomValidator()
	e.Use(middleware.RequestLogger())
	e.Use(
		middleware2.CORSWithConfig(middleware2.CORSConfig{
			AllowOrigins:     config.GetEnvCors("CORS_ORIGIN_ALLOWED"),
			AllowMethods:     config.GetEnvCors("CORS_METHOD_ALLOWED"),
			AllowHeaders:     config.GetEnvCors("CORS_HEADER_ALLOWED"),
			AllowCredentials: false,
		}),
	)

	apiGroup := e.Group("/api")
	apiGroupV1 := apiGroup.Group("/v1")
	//apiGroupV1.Use(middleware.ValidateClientSecret())

	restHandler.NewProductHandler(apiGroupV1, productService)
	restHandler.NewOrderHandler(apiGroupV1, orderService)

	err := e.Start(":9090")
	if err != nil {
		logrus.Error("Could not start server", err)
		os.Exit(1)
	}
}
