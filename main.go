package main

import (
	"github.com/chadiso/golang_tryout/app/handlers"
	"github.com/chadiso/golang_tryout/db"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())

	// Setting up Database connection
	options := db.ConnectionOptions{
		DatabaseName:   "golang_tryout",
		User:           "root",
		Password:       "",
		Port:           "3306",
		Host:           "box",
		DatabaseParams: "charset=utf8&parseTime=True&loc=Local",
	}

	requestHandler := handlers.Context{DB: db.GetConnection(options)}

	// Router
	e.GET("/status.json", func(c echo.Context) error {
		hash := map[string]interface{}{
			"status": "success",
		}
		return c.JSON(http.StatusOK, hash)
	})
	e.GET("/transactions.json", requestHandler.ListTransactions)

	// Ensure DB connection is closed
	defer requestHandler.DB.Close()

	e.Logger.Fatal(e.Start(":8080"))
}
