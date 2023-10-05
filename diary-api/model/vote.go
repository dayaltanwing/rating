package model

import (
	"diary-api/database"

	"gorm.io/gorm"
)

type Vote struct {
	gorm.Model
	EntryID uint `json:"entry_id"`
	FilmID  uint `json:"film_id"`
	Score   uint `json:"score"`
}

func (v *Vote) Save() (*Vote, error) {
	err := database.Database.Create(&v).Error
	if err != nil {
		return &Vote{}, err
	}
	return v, nil
}
