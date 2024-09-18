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

func getRecipes(c echo.Context) error {
	// dynamoClient := getClient()
	recipe := Recipe{Title: "torta", Ingredients: []Ingredient{{
		Name: "apple", Price: 2, Unit: "kg",
	}, {
		Name: "butter", Price: 2, Unit: "kg",
	}}}
	// rec := GetRecipes(dynamoClient)

	return c.JSON(http.StatusOK, recipe)
}

func createRecipe(c echo.Context) error {
	title := c.FormValue("title")
	recipe := Recipe{Title: title, Ingredients: []Ingredient{{
		Name: "apple", Price: 2, Unit: "kg",
	}}}
	return c.JSON(http.StatusOK, recipe)
}

func main() {
	e := echo.New()
	e.GET("/recipes", getRecipes)
	e.POST("/recipe", createRecipe)
	e.Logger.Fatal(e.Start(":8080"))
}
