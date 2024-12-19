// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: © 2017 LabStack and Echo contributors

package echopprof

import (
	"net/http"
	"net/http/pprof"

	"github.com/labstack/echo/v4"
)

func RegisterPprofRoutes(e *echo.Echo) {
	// Create a subrouter for pprof
	pprofGroup := e.Group("/debug/pprof")

	// Bind the standard pprof handlers
	pprofGroup.GET("/", echo.WrapHandler(http.HandlerFunc(pprof.Index)))
	pprofGroup.GET("/cmdline", echo.WrapHandler(http.HandlerFunc(pprof.Cmdline)))
	pprofGroup.GET("/profile", echo.WrapHandler(http.HandlerFunc(pprof.Profile)))
	pprofGroup.POST("/symbol", echo.WrapHandler(http.HandlerFunc(pprof.Symbol)))
	pprofGroup.GET("/symbol", echo.WrapHandler(http.HandlerFunc(pprof.Symbol)))
	pprofGroup.GET("/trace", echo.WrapHandler(http.HandlerFunc(pprof.Trace)))
	pprofGroup.GET("/allocs", echo.WrapHandler(http.HandlerFunc(pprof.Handler("allocs").ServeHTTP)))
	pprofGroup.GET("/block", echo.WrapHandler(http.HandlerFunc(pprof.Handler("block").ServeHTTP)))
	pprofGroup.GET("/goroutine", echo.WrapHandler(http.HandlerFunc(pprof.Handler("goroutine").ServeHTTP)))
	pprofGroup.GET("/heap", echo.WrapHandler(http.HandlerFunc(pprof.Handler("heap").ServeHTTP)))
	pprofGroup.GET("/mutex", echo.WrapHandler(http.HandlerFunc(pprof.Handler("mutex").ServeHTTP)))
	pprofGroup.GET("/threadcreate", echo.WrapHandler(http.HandlerFunc(pprof.Handler("threadcreate").ServeHTTP)))
}
