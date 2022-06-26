package main

import (
	"echo/database"
	"echo/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

func healthCheck(c echo.Context) error {
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()
	err := sqlDB.Ping()
	if err != nil {
	  return c.String(http.StatusInternalServerError, "データベース接続失敗")
	} else {
	  return c.String(http.StatusOK, "200, OK")
	}
  }
  
  func main() {
	e := echo.New()
	database.Connect()
	sqlDB, _ := database.DB.DB()

	// routes
	e.GET("/api", healthCheck)
	e.GET("/api/users", controllers.GetUsers)
	e.GET("/api/users/:id", controllers.GetUser)
	e.POST("/api/users", controllers.CreateUser)
	e.PATCH("/api/users/:id", controllers.UpdateUser)
	e.DELETE("/api/users/:id", controllers.DeleteUser)
	
	defer sqlDB.Close()
	e.Logger.Fatal(e.Start(":8080"))
  }