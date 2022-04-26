package models

import (
	"time"
)

// Recipies defines the recipie data to be used from db and api
type Recipes struct {
	IdMeal                      int
	StrMeal                     string `json:"StrMeal"`
	StrDrinkAlternate           string
	StrCategory                 string
	StrArea                     string
	StrInstructions             string
	StrMealThumb                string
	StrTags                     string
	StrYoutube                  string
	Strngredient1               string
	Strngredient2               string
	Strngredient3               string
	Strngredient4               string
	Strngredient5               string
	Strngredient6               string
	Strngredient7               string
	Strngredient8               string
	Strngredient9               string
	Strngredient10              string
	Strngredient11              string
	Strngredient12              string
	Strngredient13              string
	Strngredient14              string
	Strngredient15              string
	Strngredient16              string
	Strngredient17              string
	Strngredient18              string
	Strngredient19              string
	Strngredient20              string
	StrMeasure1                 string
	StrMeasure2                 string
	StrMeasure3                 string
	StrMeasure4                 string
	StrMeasure5                 string
	StrMeasure6                 string
	StrMeasure7                 string
	StrMeasure8                 string
	StrMeasure9                 string
	StrMeasure10                string
	StrMeasure11                string
	StrMeasure12                string
	StrMeasure13                string
	StrMeasure14                string
	StrMeasure15                string
	StrMeasure16                string
	StrMeasure17                string
	StrMeasure18                string
	StrMeasure19                string
	StrMeasure20                string
	StrSource                   string
	StrImageSource              string
	StrCreativeCommonsConfirmed string
	DateCreated                 time.Time
	DateModified                time.Time
}

type Recipesslice struct {
	Meals []Recipes
}
