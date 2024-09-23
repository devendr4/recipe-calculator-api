package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Recipe struct {
	Title       string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
	Image       string       `json:image`
}

type Ingredient struct {
	Name string `json:"name"`

	Price float64 `json:"price"`
	Unit  string  `json:"unit"`
}

func getRecipes(c echo.Context) error {
	recipes := GetRecipes()
	return c.JSON(http.StatusOK, recipes)
}

func createRecipe(c echo.Context) error {
	title := c.FormValue("title")
	image, err := c.FormFile("image")
	if err != nil {
		return err
	}
	src, err := image.Open()
	defer src.Close()
	println(image.Filename)
	// uploadImg(src, image.Filename, image.Header.Get("Content-Type"))

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
