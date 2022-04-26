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

<<<<<<< HEAD
=======
	fmt.Println("this is rec", rec, "\n")

>>>>>>> 29e31f792293455a169f30fdba11d417a4a11970
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
	// for i := 0; i < len(r.Meals); i++ {
	// 	m.sqlDBUpdate(sqlStatment, r)
	// }

	// for i := 0; i < len(r.Meals); i++ {
	// 	fmt.Println(r.Meals[i].IDMeal, r.Meals[i].StrArea)
	// }

	return nil
}

// // interate over multiple structs
// func (m *postgresDBRepo) sqlDBUpdate(sqlQuery string, r models.MealsJSON) error {

// 	var error1 error

// 	for i := 0; i < len(r.Meals); i++ {

// 		_, err := m.DB.Exec(sqlQuery,
// 			r.Meals[i].IDMeal,
// 			r.Meals[i].StrDrinkAlternate,
// 			r.Meals[i].StrCategory,
// 			r.Meals[i].StrArea,
// 			r.Meals[i].StrInstructions,
// 			r.Meals[i].StrMealThumb,
// 			r.Meals[i].StrTags,
// 			r.Meals[i].StrYoutube,
// 			r.Meals[i].StrIngredient1,
// 			r.Meals[i].StrIngredient2,
// 			r.Meals[i].StrIngredient3,
// 			r.Meals[i].StrIngredient4,
// 			r.Meals[i].StrIngredient5,
// 			r.Meals[i].StrIngredient6,
// 			r.Meals[i].StrIngredient7,
// 			r.Meals[i].StrIngredient8,
// 			r.Meals[i].StrIngredient9,
// 			r.Meals[i].StrIngredient10,
// 			r.Meals[i].StrIngredient11,
// 			r.Meals[i].StrIngredient12,
// 			r.Meals[i].StrIngredient13,
// 			r.Meals[i].StrIngredient14,
// 			r.Meals[i].StrIngredient15,
// 			r.Meals[i].StrIngredient16,
// 			r.Meals[i].StrIngredient17,
// 			r.Meals[i].StrIngredient18,
// 			r.Meals[i].StrIngredient19,
// 			r.Meals[i].StrIngredient20,
// 			r.Meals[i].StrMeasure1,
// 			r.Meals[i].StrMeasure2,
// 			r.Meals[i].StrMeasure3,
// 			r.Meals[i].StrMeasure4,
// 			r.Meals[i].StrMeasure5,
// 			r.Meals[i].StrMeasure6,
// 			r.Meals[i].StrMeasure7,
// 			r.Meals[i].StrMeasure8,
// 			r.Meals[i].StrMeasure9,
// 			r.Meals[i].StrMeasure10,
// 			r.Meals[i].StrMeasure11,
// 			r.Meals[i].StrMeasure12,
// 			r.Meals[i].StrMeasure13,
// 			r.Meals[i].StrMeasure14,
// 			r.Meals[i].StrMeasure15,
// 			r.Meals[i].StrMeasure16,
// 			r.Meals[i].StrMeasure17,
// 			r.Meals[i].StrMeasure18,
// 			r.Meals[i].StrMeasure19,
// 			r.Meals[i].StrMeasure20,
// 			r.Meals[i].StrSource,
// 			r.Meals[i].StrImageSource,
// 			r.Meals[i].StrCreativeCommonsConfirmed,
// 			time.Now(),
// 			time.Now(),
// 		)
// 		if err != nil {
// 			error1 = err
// 		}
// 	}
// 	return error1
// }
