package main

import (
	"net/http"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		// for StackdriverLogging
		logger.Debug("DEBUG LEVEL with severity", zap.String("severity", "DEBUG"))
		logger.Info("INFO LEVEL with severity", zap.String("severity", "INFO"))
		return c.String(http.StatusOK, "Hello, Go World!")
	})
	err = e.Start(":80")
	if err != nil {
		logger.Error("ERROR LEVEL with severity", zap.String("severity", "ERROR"), zap.Error(err))
		e.Logger.Fatal(err)
	}
}
