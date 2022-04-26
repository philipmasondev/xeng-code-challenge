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

// InsetRecipiesToDB inserts recipies to recipe table.
func (m *postgresDBRepo) InsertRecipesToDB(recipeJSON string) error {

	// auto fill created and update times for db columns
	createdAt := time.Now()
	updatedAt := time.Now()

	//sqlQuery := `SELECT * FROM recipes WHERE recipe->>'strMeal' ? 'Arrabiata'`

	//rows := m.DB.QueryRow(sqlQuery)
	//fmt.Println(rows)
	//if rows {

	sqlStatment := `insert into recipes (recipe, created_at, updated_at)
		values ($1, $2, $3)`

	_, err := m.DB.Exec(sqlStatment, recipeJSON, createdAt, updatedAt)

	if err != nil {
		fmt.Println("error in m.DB.ExecContext. Error - ", err)
	}
	//}

	return nil
}

// InsertRecipe will update recipesNonJson table
func (m *postgresDBRepo) InsertRecipe(rec string) error {

	var mod models.Recipes
	var r models.Recipesslice
	json.Unmarshal([]byte(rec), &r)
	fmt.Println(r.Meals[0])

	sqlStatment := `insert into recipesnonjson (meal, drinkalternate, category, area, instructions, mealthumd, tags, youtube, ingredient1, ingredient2, ingredient3, ingredient4,
		ingredient5, ingredient6, ingredient7, ingredient8, ingredient9, ingredient10, ingredient11, ingredient12, ingredient13, ingredient14, ingredient15, ingredient16,
		ingredient17, ingredient18, ingredient19, ingredient20, measure1, measure2, measure3, measure4, measure5, measure6, measure7, measure8, measure9, measure10, measure11,
		measure12, measure13, measure14, measure15, measure16, measure17, measure18, measure19, measure20, source, imageSource, creativecommonsconfirmed, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35,
				$36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53)`

	_, err := m.DB.Exec(sqlStatment,
		mod.StrMeal,
		mod.StrDrinkAlternate,
		mod.StrCategory,
		mod.StrArea,
		mod.StrInstructions,
		mod.StrMealThumb,
		mod.StrTags,
		mod.StrYoutube,
		mod.Strngredient1,
		mod.Strngredient2,
		mod.Strngredient3,
		mod.Strngredient4,
		mod.Strngredient5,
		mod.Strngredient6,
		mod.Strngredient7,
		mod.Strngredient8,
		mod.Strngredient9,
		mod.Strngredient10,
		mod.Strngredient11,
		mod.Strngredient12,
		mod.Strngredient13,
		mod.Strngredient14,
		mod.Strngredient15,
		mod.Strngredient16,
		mod.Strngredient17,
		mod.Strngredient18,
		mod.Strngredient19,
		mod.Strngredient20,
		mod.StrMeasure1,
		mod.StrMeasure2,
		mod.StrMeasure3,
		mod.StrMeasure4,
		mod.StrMeasure5,
		mod.StrMeasure6,
		mod.StrMeasure7,
		mod.StrMeasure8,
		mod.StrMeasure9,
		mod.StrMeasure10,
		mod.StrMeasure11,
		mod.StrMeasure12,
		mod.StrMeasure13,
		mod.StrMeasure14,
		mod.StrMeasure15,
		mod.StrMeasure16,
		mod.StrMeasure17,
		mod.StrMeasure18,
		mod.StrMeasure19,
		mod.StrMeasure20,
		mod.StrSource,
		mod.StrImageSource,
		mod.StrCreativeCommonsConfirmed,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		fmt.Println("InsertRecipe failed with ", err)
	}

	return nil
}
