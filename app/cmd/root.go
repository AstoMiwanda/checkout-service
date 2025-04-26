package cmd

import (
	"checkout-service/internal/repository"
	"checkout-service/internal/usecase"
	"context"
	"fmt"
	"gorm.io/gorm"
	"os"
	"time"

	"checkout-service/app/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	EnvFilePath string
	rootCmd     = &cobra.Command{
		Use:   "checkout-service",
		Short: "Start Checkout Service",
	}
	databaseConn *gorm.DB

	rootConfig     config.Root
	productService *usecase.ProductService
	orderService   *usecase.OrderService
)

func init() {
	cobra.OnInitialize(func() {
		initConfigReader()
		initPostgresDB(rootConfig.Postgres)

		initApp()
	})
}
func initConfigReader() {
	fmt.Println(EnvFilePath)
	rootConfig = config.Load(EnvFilePath)
	logrus.Info("config loaded app: ", rootConfig.App.ServiceName)
	logrus.Info("Postgres loaded :", rootConfig.Postgres.Host)
}

func initPostgresDB(conf config.Postgres) {
	var err error
	secured := false
	if rootConfig.App.Env != "local" {
		secured = true
	}

	databaseConn, err = config.OpenPostgresConnection(conf, secured)
	if err != nil {
		logrus.Error("Could not establish connection to postgres", err)
		os.Exit(1)
	}
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = config.Ping(ctxTimeout, databaseConn)
	if err != nil {
		logrus.Error("Could not ping postgres", err)
		os.Exit(1)
	}
	logrus.Infof("postgres: connected")
}

func Execute() {

	fmt.Println(EnvFilePath)
	rootCmd.PersistentFlags().StringVarP(&EnvFilePath, "env", "e", ".env", ".env file to read from")
	if err := rootCmd.Execute(); err != nil {
		logrus.Error("can't start the CLI", err)
		os.Exit(1)
	}
	fmt.Println(EnvFilePath)
}

func initApp() {

	// repository
	productRepository := repository.NewProductRepository(databaseConn)
	orderRepository := repository.NewOrderRepository(databaseConn)
	discountRepository := repository.NewDiscountRepository(databaseConn)

	// service
	productService = usecase.NewProductService(productRepository)
	orderService = usecase.NewOrderService(orderRepository, discountRepository, productRepository)

}
