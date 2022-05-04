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

	recipe := models.MealsJSONinsert{}

	json.Unmarshal([]byte(rec), &recipe)

	sqlStatment := `insert into recipesnonjson (id, meal, drinkalternate, category, area, instructions, mealthumd, tags, youtube, ingredient1, ingredient2, ingredient3, ingredient4,
		ingredient5, ingredient6, ingredient7, ingredient8, ingredient9, ingredient10, ingredient11, ingredient12, ingredient13, ingredient14, ingredient15, ingredient16,
		ingredient17, ingredient18, ingredient19, ingredient20, measure1, measure2, measure3, measure4, measure5, measure6, measure7, measure8, measure9, measure10, measure11,
		measure12, measure13, measure14, measure15, measure16, measure17, measure18, measure19, measure20, source, imageSource, creativecommonsconfirmed, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35,
				$36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54)`

	_, err := m.DB.Exec(sqlStatment,
		recipe.Meals[0].IDMeal,
		recipe.Meals[0].StrMeal,
		recipe.Meals[0].StrDrinkAlternate,
		recipe.Meals[0].StrCategory,
		recipe.Meals[0].StrArea,
		recipe.Meals[0].StrInstructions,
		recipe.Meals[0].StrMealThumb,
		recipe.Meals[0].StrTags,
		recipe.Meals[0].StrYoutube,
		recipe.Meals[0].StrIngredient1,
		recipe.Meals[0].StrIngredient2,
		recipe.Meals[0].StrIngredient3,
		recipe.Meals[0].StrIngredient4,
		recipe.Meals[0].StrIngredient5,
		recipe.Meals[0].StrIngredient6,
		recipe.Meals[0].StrIngredient7,
		recipe.Meals[0].StrIngredient8,
		recipe.Meals[0].StrIngredient9,
		recipe.Meals[0].StrIngredient10,
		recipe.Meals[0].StrIngredient11,
		recipe.Meals[0].StrIngredient12,
		recipe.Meals[0].StrIngredient13,
		recipe.Meals[0].StrIngredient14,
		recipe.Meals[0].StrIngredient15,
		recipe.Meals[0].StrIngredient16,
		recipe.Meals[0].StrIngredient17,
		recipe.Meals[0].StrIngredient18,
		recipe.Meals[0].StrIngredient19,
		recipe.Meals[0].StrIngredient20,
		recipe.Meals[0].StrMeasure1,
		recipe.Meals[0].StrMeasure2,
		recipe.Meals[0].StrMeasure3,
		recipe.Meals[0].StrMeasure4,
		recipe.Meals[0].StrMeasure5,
		recipe.Meals[0].StrMeasure6,
		recipe.Meals[0].StrMeasure7,
		recipe.Meals[0].StrMeasure8,
		recipe.Meals[0].StrMeasure9,
		recipe.Meals[0].StrMeasure10,
		recipe.Meals[0].StrMeasure11,
		recipe.Meals[0].StrMeasure12,
		recipe.Meals[0].StrMeasure13,
		recipe.Meals[0].StrMeasure14,
		recipe.Meals[0].StrMeasure15,
		recipe.Meals[0].StrMeasure16,
		recipe.Meals[0].StrMeasure17,
		recipe.Meals[0].StrMeasure18,
		recipe.Meals[0].StrMeasure19,
		recipe.Meals[0].StrMeasure20,
		recipe.Meals[0].StrSource,
		recipe.Meals[0].StrImageSource,
		recipe.Meals[0].StrCreativeCommonsConfirmed,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		fmt.Errorf("Error at insert: %v", err)
	}

	return nil
}

// GetAllAPI gets all records from db
func (m *postgresDBRepo) GetAllAPI() string {

	rows, err := m.DB.Query("SELECT * FROM recipesnonjson")
	if err != nil {
		fmt.Println(err)
	}

	var test []models.MealsSearch

	for rows.Next() {
		recipe := models.MealsSearch{}

		err := rows.Scan(
			&recipe.ID,
			&recipe.IDMeal,
			&recipe.StrMeal,
			&recipe.StrDrinkAlternate,
			&recipe.StrCategory,
			&recipe.StrArea,
			&recipe.StrInstructions,
			&recipe.StrMealThumb,
			&recipe.StrTags,
			&recipe.StrYoutube,
			&recipe.StrIngredient1,
			&recipe.StrIngredient2,
			&recipe.StrIngredient3,
			&recipe.StrIngredient4,
			&recipe.StrIngredient5,
			&recipe.StrIngredient6,
			&recipe.StrIngredient7,
			&recipe.StrIngredient8,
			&recipe.StrIngredient9,
			&recipe.StrIngredient10,
			&recipe.StrIngredient11,
			&recipe.StrIngredient12,
			&recipe.StrIngredient13,
			&recipe.StrIngredient14,
			&recipe.StrIngredient15,
			&recipe.StrIngredient16,
			&recipe.StrIngredient17,
			&recipe.StrIngredient18,
			&recipe.StrIngredient19,
			&recipe.StrIngredient20,
			&recipe.StrMeasure1,
			&recipe.StrMeasure2,
			&recipe.StrMeasure3,
			&recipe.StrMeasure4,
			&recipe.StrMeasure5,
			&recipe.StrMeasure6,
			&recipe.StrMeasure7,
			&recipe.StrMeasure8,
			&recipe.StrMeasure9,
			&recipe.StrMeasure10,
			&recipe.StrMeasure11,
			&recipe.StrMeasure12,
			&recipe.StrMeasure13,
			&recipe.StrMeasure14,
			&recipe.StrMeasure15,
			&recipe.StrMeasure16,
			&recipe.StrMeasure17,
			&recipe.StrMeasure18,
			&recipe.StrMeasure19,
			&recipe.StrMeasure20,
			&recipe.StrSource,
			&recipe.StrImageSource,
			&recipe.StrCreativeCommonsConfirmed,
			&recipe.DateModified,
		)
		if err != nil {
			fmt.Println(err)
		}

		test = append(test, recipe)

	}
	recipeByte, _ := json.MarshalIndent(test, "", "\t")

	return (string(recipeByte))
}

// GetAllAPI gets all records from db
func (m *postgresDBRepo) Get() string {

	rows, err := m.DB.Query("SELECT * FROM recipesnonjson")
	if err != nil {
		fmt.Println(err)
	}

	var test []models.MealsSearch

	for rows.Next() {
		recipe := models.MealsSearch{}

		err := rows.Scan(
			&recipe.ID,
			&recipe.IDMeal,
			&recipe.StrMeal,
			&recipe.StrDrinkAlternate,
			&recipe.StrCategory,
			&recipe.StrArea,
			&recipe.StrInstructions,
			&recipe.StrMealThumb,
			&recipe.StrTags,
			&recipe.StrYoutube,
			&recipe.StrIngredient1,
			&recipe.StrIngredient2,
			&recipe.StrIngredient3,
			&recipe.StrIngredient4,
			&recipe.StrIngredient5,
			&recipe.StrIngredient6,
			&recipe.StrIngredient7,
			&recipe.StrIngredient8,
			&recipe.StrIngredient9,
			&recipe.StrIngredient10,
			&recipe.StrIngredient11,
			&recipe.StrIngredient12,
			&recipe.StrIngredient13,
			&recipe.StrIngredient14,
			&recipe.StrIngredient15,
			&recipe.StrIngredient16,
			&recipe.StrIngredient17,
			&recipe.StrIngredient18,
			&recipe.StrIngredient19,
			&recipe.StrIngredient20,
			&recipe.StrMeasure1,
			&recipe.StrMeasure2,
			&recipe.StrMeasure3,
			&recipe.StrMeasure4,
			&recipe.StrMeasure5,
			&recipe.StrMeasure6,
			&recipe.StrMeasure7,
			&recipe.StrMeasure8,
			&recipe.StrMeasure9,
			&recipe.StrMeasure10,
			&recipe.StrMeasure11,
			&recipe.StrMeasure12,
			&recipe.StrMeasure13,
			&recipe.StrMeasure14,
			&recipe.StrMeasure15,
			&recipe.StrMeasure16,
			&recipe.StrMeasure17,
			&recipe.StrMeasure18,
			&recipe.StrMeasure19,
			&recipe.StrMeasure20,
			&recipe.StrSource,
			&recipe.StrImageSource,
			&recipe.StrCreativeCommonsConfirmed,
			&recipe.DateModified,
		)
		if err != nil {
			fmt.Println(err)
		}

		test = append(test, recipe)

	}
	recipeByte, _ := json.MarshalIndent(test, "", "\t")

	return (string(recipeByte))
}
