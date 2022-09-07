package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

var Logger = middleware.BodyDump(func(c echo.Context, request []byte, response []byte) {
	log.WithFields(log.Fields{
		"message": c.Request().Method,
		"uri":     c.Request().RequestURI,
		"status":  c.Response().Status,
	}).Info(string(response[:]))
})
