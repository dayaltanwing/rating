package model

import (
	"diary-api/database"

	"gorm.io/gorm"
)

type Film struct {
	gorm.Model
	Name        string `gorm:"type:text" json:"content"`
	CountEntry  uint
	RatingScore uint
	EntryID     uint
	Entries     []Entry
}

func (f *Film) Save() (*Film, error) {
	err := database.Database.Create(&f).Error
	if err != nil {
		return &Film{}, err
	}
	return f, nil
}

func FindFilmByName(name string) (Film, error) {
	var film Film
	err := database.Database.Where("name=?", name).Find(&film).Error
	if err != nil {
		return Film{}, err
	}
	return film, nil
}

func GetAllFilms() ([]Film, error) {
	var films []Film
	if err := database.Database.Find(&films).Error; err != nil {
		return nil, err
	}
	return films, nil
}

func DeleteFilm(name string) (Film, error) {
	var film Film
	err := database.Database.Where("name=?", film).Delete(&film).Error
	if err != nil {
		return Film{}, err
	}
	return film, nil
}

func FindFilmById(id uint) (Film, error) {
	var film Film
	err := database.Database.Where("ID=?", id).Find(&film).Error
	if err != nil {
		return Film{}, err
	}
	return film, nil
}