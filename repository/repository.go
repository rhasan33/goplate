package repository

import (
	"context"

	"github.com/rhasan33/goplate/models"
)

// UserRepo ..
type UserRepo interface {
	CreateUser(ctx context.Context, us *models.UserSettings) (*models.UserSettings, error)
	GetUser(ctx context.Context, id int) (*models.UserSettings, error)
}
