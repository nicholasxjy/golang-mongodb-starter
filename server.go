package main

import (
	"fmt"

	"edenedu/conf"
	"edenedu/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	e := echo.New()

	e.SetDebug(true)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.SecureWithConfig(middleware.DefaultSecureConfig))
	e.Use(middleware.Static("/static"))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:9000", "http://localhost:8080"},
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath: "/",
		// CookieSecure:   true,
		CookieHTTPOnly: true,
	}))

	// connect mongodb
	session, err := mgo.Dial(conf.MongodbURL)
	if err != nil {
		fmt.Println("mongodb connecting error")
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	defer session.Close()

	// init controllers inject the session

	testController := &controllers.TestController{DBSession: session}

	// api route group
	g := e.Group("/api/v1")
	g.GET("/tests", testController.TestGet)
	g.GET("/tests/insert", testController.TestInsert)
	e.Run(standard.New(":9000"))

}
