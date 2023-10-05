package model

import (
	"diary-api/database"

	"gorm.io/gorm"
)

type Film struct {
	gorm.Model
	Name        string `gorm:"type:text" json:"name"`
	CountEntry  uint
	RatingScore uint
	Votes       []Vote `gorm:"many2many:film_votes;" json:"votes"`
	EntryID     uint
}

type ScoreResult struct {
    TotalScore int
    VoteCount  int
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

func CalculateAverageScoreForFilm(filmID uint) (float64, error) {
    var result ScoreResult

    if err := database.Database.Model(&Film{}).Where("id = ?", filmID).
        Joins("JOIN film_votes ON films.id = film_votes.film_id").
        Select("SUM(film_votes.score) AS total_score, COUNT(film_votes.score) AS vote_count").
        Scan(&result).Error; err != nil {
        return 0.0, err
    }

    if result.VoteCount == 0 {
        return 0.0, nil 
    }

    averageScore := float64(result.TotalScore) / float64(result.VoteCount)
    return averageScore, nil
}

