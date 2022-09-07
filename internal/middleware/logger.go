package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

var Logger = middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
	LogURI:     true,
	LogStatus:  true,
	LogLatency: true,
	LogMethod:  true,
	LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
		log.WithFields(log.Fields{
			"method":  values.Method,
			"uri":     values.URI,
			"status":  values.Status,
			"latency": values.Latency,
		}).Info("Request info")

		return nil
	},
})