package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ziflex/lecho/v3"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

func main() {
	// Echo instance
	e := echo.New()

	// define zerolog with console output
	// log := zerolog.New(os.Stdout)
	// define zero log with console and file output
	log := initConsoleLoggerWithRollingFile()
	// uses a wrapper to set zerolog into echo
	e.Logger = lecho.From(log)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Logger =

	// generate request id
	e.Use(middleware.RequestID())

	// Routes
	e.GET("/", rootGetPathHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func rootGetPathHandler(c echo.Context) error {
	// validate
	validateRequest(c)

	// process
	processRequest(c)

	// assemble result
	assemleResult(c)

	return c.String(http.StatusOK, "Hello, World!")
}

func validateRequest(c echo.Context) {
	c.Echo().Logger.Info("Validate request")
}

func processRequest(c echo.Context) {
	c.Logger().Info("process request")
}

func assemleResult(c echo.Context) {
	c.Logger().Info("assemble processing result")
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func initConsoleLoggerWithRollingFile() zerolog.Logger {
	var lumberJackLogger = &lumberjack.Logger{
		Filename:   "./zerolog.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}

	multi := zerolog.MultiLevelWriter(lumberJackLogger, os.Stdout)
	logger := zerolog.New(multi).With().Timestamp().Logger()
	logger.Info().Msg("Zerolog and Lumberjack setup complete")

	return logger
}
