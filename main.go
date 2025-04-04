package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/simple-bookStore-server/config"
	"github.com/simple-bookStore-server/controllers"
)

//middlewares here

func middlewareOne(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("This is from the middleware one ")
		return next(ctx)
	}
}

func SomeHandler(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		map[string]any{"Hello": "Hello, 世界 !"},
	)
}

func main() {

	e := echo.New()

	//connecting to database
	config.DatbaseInit()
	gorm := config.DB()

	dbgorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbgorm.Ping()

	//routes

	bookr := e.Group("/book")

	bookr.GET("/:id", controllers.GetBook)
	// bookr.POST("/:id", controllers.AddBook)
	bookr.PUT("/:id", controllers.UpdateBook)
	bookr.DELETE("/:id", controllers.DeleteBook)

	//starting the server
	fmt.Println("the server is running on port 3000")
	e.Logger.Fatal(e.Start(":3000"))

}
