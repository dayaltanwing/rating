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
	Films     []Film `gorm:"many2many:entry_films;" json:"films"`
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
	err := database.Database.Where("entryID=?", id).Find(&entry).Error
	if err != nil {
		return Entry{}, err
	}
	return entry, nil
}

func DeleteEntry(id uint) (Entry, error) {
	var entry Entry
	err := database.Database.Where("entryID=?", id).Delete(&entry).Error
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

func Count(id uint) (int64, error) {
	var entryCount int64
	if err := database.Database.Where("user_id = ?", id).Count(&entryCount).Error; err != nil {
		return 0, err
	}

	return entryCount, nil
}
