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

	fmt.Println("this is rec", rec, "\n")

	r := models.MealsJSON{}

	json.Unmarshal([]byte(rec), &r)

	fmt.Println("this is r", r, "\n")

	// sqlStatment := `insert into recipesnonjson (meal, drinkalternate, category, area, instructions, mealthumd, tags, youtube, ingredient1, ingredient2, ingredient3, ingredient4,
	// 	ingredient5, ingredient6, ingredient7, ingredient8, ingredient9, ingredient10, ingredient11, ingredient12, ingredient13, ingredient14, ingredient15, ingredient16,
	// 	ingredient17, ingredient18, ingredient19, ingredient20, measure1, measure2, measure3, measure4, measure5, measure6, measure7, measure8, measure9, measure10, measure11,
	// 	measure12, measure13, measure14, measure15, measure16, measure17, measure18, measure19, measure20, source, imageSource, creativecommonsconfirmed, created_at, updated_at)
	// 	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35,
	// 			$36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53)`

	// _, err := m.DB.Exec(sqlStatment,
	// 	r.StrMeal,
	// 	r.StrDrinkAlternate,
	// 	r.StrCategory,
	// 	r.StrArea,
	// 	r.StrInstructions,
	// 	r.StrMealThumb,
	// 	r.StrTags,
	// 	r.StrYoutube,
	// 	r.StrIngredient1,
	// 	r.StrIngredient2,
	// 	r.StrIngredient3,
	// 	r.StrIngredient4,
	// 	r.StrIngredient5,
	// 	r.StrIngredient6,
	// 	r.StrIngredient7,
	// 	r.StrIngredient8,
	// 	r.StrIngredient9,
	// 	r.StrIngredient10,
	// 	r.StrIngredient11,
	// 	r.StrIngredient12,
	// 	r.StrIngredient13,
	// 	r.StrIngredient14,
	// 	r.StrIngredient15,
	// 	r.StrIngredient16,
	// 	r.StrIngredient17,
	// 	r.StrIngredient18,
	// 	r.StrIngredient19,
	// 	r.StrIngredient20,
	// 	r.StrMeasure1,
	// 	r.StrMeasure2,
	// 	r.StrMeasure3,
	// 	r.StrMeasure4,
	// 	r.StrMeasure5,
	// 	r.StrMeasure6,
	// 	r.StrMeasure7,
	// 	r.StrMeasure8,
	// 	r.StrMeasure9,
	// 	r.StrMeasure10,
	// 	r.StrMeasure11,
	// 	r.StrMeasure12,
	// 	r.StrMeasure13,
	// 	r.StrMeasure14,
	// 	r.StrMeasure15,
	// 	r.StrMeasure16,
	// 	r.StrMeasure17,
	// 	r.StrMeasure18,
	// 	r.StrMeasure19,
	// 	r.StrMeasure20,
	// 	r.StrSource,
	// 	r.StrImageSource,
	// 	r.StrCreativeCommonsConfirmed,
	// 	time.Now(),
	// 	time.Now(),
	// )

	fmt.Printf("this is after the sqlStatement: %+v\n", r)

	for i := 0; i < len(r.Meals); i++ {
		fmt.Println("Meal ID:", r.Meals[i].IDMeal)
		fmt.Println("Meal Area:", r.Meals[i].StrArea)

	}

	// if err != nil {
	// 	fmt.Println("InsertRecipe failed with ", err)
	// }

	return nil
}
