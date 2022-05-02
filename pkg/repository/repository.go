package repository

type DatabaseRepo interface {
	AllUsers() bool

	InsertRecipe(rec string) error
	GetAllAPI() string
	SearchAPI(string) string
}
