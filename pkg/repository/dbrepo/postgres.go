package dbrepo

import (
	"context"
	"time"
	"web-application-template/pkg/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// InsetRecipiesToDB inserts recipies to recipe table.
func (m *postgresDBRepo) InsertRecipiesToDB(recipe models.RecipiesJSON) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	InsStatment := `insert into recipies (recipe)
		values ($1)`

	_, err := m.DB.ExecContext(ctx, InsStatment,
		recipe.RecipeJsonString,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}
