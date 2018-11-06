package reader

import (
	"context"

	"github.com/rhasan33/goplate/conn"
	"github.com/rhasan33/goplate/models"
	"github.com/rhasan33/goplate/repository"
)

type readerRepo struct {
}

// CreateUser create user repo
func (rr *readerRepo) CreateUser(ctx context.Context, us *models.UserSettings) (*models.UserSettings, error) {
	db := conn.PostgresDB()
	err := db.Create(&us).Error
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (rr *readerRepo) GetUser(ctx context.Context, id int) (*models.UserSettings, error) {
	var userData models.UserSettings
	db := conn.PostgresDB()
	data := db.Where("id = ", id).Find(&userData)
	if data.Error != nil {
		return nil, data.Error
	}
	return &userData, nil
}

// NewReader exportable repository
func NewReader() repository.UserRepo {
	return &readerRepo{}
}
