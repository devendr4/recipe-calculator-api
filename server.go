package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Recipe struct {
	Title       string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Unit  string  `json:"unit"`
}

func main() {
	recipe := Recipe{Title: "torta", Ingredients: []Ingredient{{
		Name: "sugar", Price: 2, Unit: "kg",
	}}}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, recipe)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
