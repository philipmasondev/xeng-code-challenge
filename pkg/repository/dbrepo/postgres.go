package dbrepo

import (
	"encoding/json"
	"fmt"
	"time"
	"web-application-template/pkg/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

/**********************************************************************************************
	TODO:
	www.themealdb.com/api/json/v1/1/filter.php?i=chicken_breast

	This API response does not fit into the current format for the table. Need to expand the
	ability for updating table with multiple objectes in JSON string.

**********************************************************************************************/

// InsertRecipe will update recipesNonJson table
func (m *postgresDBRepo) InsertRecipe(rec string) error {

	r := models.MealsJSON{}

	json.Unmarshal([]byte(rec), &r)

	sqlStatment := `insert into recipesnonjson (meal, drinkalternate, category, area, instructions, mealthumd, tags, youtube, ingredient1, ingredient2, ingredient3, ingredient4,
		ingredient5, ingredient6, ingredient7, ingredient8, ingredient9, ingredient10, ingredient11, ingredient12, ingredient13, ingredient14, ingredient15, ingredient16,
		ingredient17, ingredient18, ingredient19, ingredient20, measure1, measure2, measure3, measure4, measure5, measure6, measure7, measure8, measure9, measure10, measure11,
		measure12, measure13, measure14, measure15, measure16, measure17, measure18, measure19, measure20, source, imageSource, creativecommonsconfirmed, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35,
				$36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53)`

	_, err := m.DB.Exec(sqlStatment,
		r.Meals[0].IDMeal,
		r.Meals[0].StrDrinkAlternate,
		r.Meals[0].StrCategory,
		r.Meals[0].StrArea,
		r.Meals[0].StrInstructions,
		r.Meals[0].StrMealThumb,
		r.Meals[0].StrTags,
		r.Meals[0].StrYoutube,
		r.Meals[0].StrIngredient1,
		r.Meals[0].StrIngredient2,
		r.Meals[0].StrIngredient3,
		r.Meals[0].StrIngredient4,
		r.Meals[0].StrIngredient5,
		r.Meals[0].StrIngredient6,
		r.Meals[0].StrIngredient7,
		r.Meals[0].StrIngredient8,
		r.Meals[0].StrIngredient9,
		r.Meals[0].StrIngredient10,
		r.Meals[0].StrIngredient11,
		r.Meals[0].StrIngredient12,
		r.Meals[0].StrIngredient13,
		r.Meals[0].StrIngredient14,
		r.Meals[0].StrIngredient15,
		r.Meals[0].StrIngredient16,
		r.Meals[0].StrIngredient17,
		r.Meals[0].StrIngredient18,
		r.Meals[0].StrIngredient19,
		r.Meals[0].StrIngredient20,
		r.Meals[0].StrMeasure1,
		r.Meals[0].StrMeasure2,
		r.Meals[0].StrMeasure3,
		r.Meals[0].StrMeasure4,
		r.Meals[0].StrMeasure5,
		r.Meals[0].StrMeasure6,
		r.Meals[0].StrMeasure7,
		r.Meals[0].StrMeasure8,
		r.Meals[0].StrMeasure9,
		r.Meals[0].StrMeasure10,
		r.Meals[0].StrMeasure11,
		r.Meals[0].StrMeasure12,
		r.Meals[0].StrMeasure13,
		r.Meals[0].StrMeasure14,
		r.Meals[0].StrMeasure15,
		r.Meals[0].StrMeasure16,
		r.Meals[0].StrMeasure17,
		r.Meals[0].StrMeasure18,
		r.Meals[0].StrMeasure19,
		r.Meals[0].StrMeasure20,
		r.Meals[0].StrSource,
		r.Meals[0].StrImageSource,
		r.Meals[0].StrCreativeCommonsConfirmed,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		fmt.Println(err)
	}
	return nil
}
