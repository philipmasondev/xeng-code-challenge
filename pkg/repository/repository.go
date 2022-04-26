package repository

type DatabaseRepo interface {
	AllUsers() bool

	InsertRecipesToDB(res string) error

	InsertRecipe(rec string) error
}
