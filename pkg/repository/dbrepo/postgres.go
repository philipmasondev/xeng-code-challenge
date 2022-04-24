package dbrepo

import (
	"fmt"
	"time"
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
