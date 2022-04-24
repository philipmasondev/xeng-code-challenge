package repository

import "web-application-template/pkg/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertRecipiesToDB(recipe models.RecipiesJSON) error
}
