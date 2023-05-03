package repository

import "gorm.io/gorm"

type Repository interface {
	InsertData(input InputDataGame) (InputDataGame, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) InsertData(input InputDataGame) (InputDataGame, error) {
	err := r.db.Create(input).Error
	if err != nil {
		return input, err
	}

	return input, nil
}
