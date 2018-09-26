package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	// ctx := context.Background()
	// traceCli, err := trace.NewClient(ctx, "【プロジェクトID】")
	// if err != nil {
	// 	logger.Fatal("Fatal", zap.Error(err))
	// }

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	e.GET("/", func(c echo.Context) error {
		logger.Info("RealIP: " + c.RealIP())
		// span := traceCli.SpanFromRequest(c.Request())
		// defer span.Finish()
		// for StackdriverLogging
		logger.Debug("DEBUG LEVEL with severity", zap.String("severity", "DEBUG"))
		logger.Info("INFO LEVEL with severity", zap.String("severity", "INFO"))
		logger.Warn("WARN LEVEL with severity", zap.String("severity", "WARN"))
		logger.Error("ERROR LEVEL with severity", zap.String("severity", "ERROR"))
		return c.JSON(http.StatusOK, "{\"pg\":\"golang\",\"word\":\"Hello\"}")
	})
	err = e.Start(":80")
	if err != nil {
		logger.Error("ERROR LEVEL with severity", zap.String("severity", "ERROR"), zap.Error(err))
		e.Logger.Fatal(err)
	}
}
