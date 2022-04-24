package models

import "time"

// RecipiesJSON holds the json string to try to use jsonb
type RecipesJSON struct {
	RecipeJsonString string
}

// Recipies defines the recipie data to be used from db and api
type Recipes struct {
	MealID                   string
	MealName                 string
	Drink                    string
	Category                 string
	Area                     string
	Instructions             string
	MealThumbnail            string
	Tags                     string
	Youtube                  string
	Ingredient1              string
	Ingredient2              string
	Ingredient3              string
	Ingredient4              string
	Ingredient5              string
	Ingredient6              string
	Ingredient7              string
	Ingredient8              string
	Ingredient9              string
	Ingredient10             string
	Ingredient11             string
	Ingredient12             string
	Ingredient13             string
	Ingredient14             string
	Ingredient15             string
	Ingredient16             string
	Ingredient17             string
	Ingredient18             string
	Ingredient19             string
	Ingredient20             string
	Measure1                 string
	Measure2                 string
	Measure3                 string
	Measure4                 string
	Measure5                 string
	Measure6                 string
	Measure7                 string
	Measure8                 string
	Measure9                 string
	Measure10                string
	Measure11                string
	Measure12                string
	Measure13                string
	Measure14                string
	Measure15                string
	Measure16                string
	Measure17                string
	Measure18                string
	Measure19                string
	Measure20                string
	Source                   string
	ImageSource              string
	CreativeCommonsConfirmed string
	Modified                 time.Time
}
