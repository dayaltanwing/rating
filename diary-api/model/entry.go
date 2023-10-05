package model

import (
	"diary-api/database"

	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Content   string `gorm:"type:text" json:"content"`
	CountFilm uint
	UserID    uint
	FilmID    uint
	Votes  []Vote `gorm:"many2many:film_votes;" json:"votes"`
}

func (entry *Entry) Save() (*Entry, error) {
	err := database.Database.Create(&entry).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}

func FindEntryById(id uint) (Entry, error) {
	var entry Entry
	err := database.Database.Where("ID=?", id).Find(&entry).Error
	if err != nil {
		return Entry{}, err
	}
	return entry, nil
}

func DeleteEntry(id uint) (Entry, error) {
	var entry Entry
	err := database.Database.Where("ID=?", id).Delete(&entry).Error
	if err != nil {
		return Entry{}, err
	}
	return entry, nil
}

func GetAllEntries() ([]Entry, error) {
	var entries []Entry
	if err := database.Database.Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}
