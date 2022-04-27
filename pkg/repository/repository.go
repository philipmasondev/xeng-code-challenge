package repository

type DatabaseRepo interface {
	AllUsers() bool

	InsertRecipe(rec string) error
}
