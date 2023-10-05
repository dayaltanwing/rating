package controller

import (
	"diary-api/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddFilm(context *gin.Context) {
	var input model.Film

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedFilm, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedFilm})
}

func GetFilmbyName(context *gin.Context) {
	name := context.Param("name")
	film, err := model.FindFilmByName(name)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find film"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"find successfully": film})
}

func DeleteFilm(context *gin.Context) {
	name := context.Param("name")

	_, err := model.DeleteFilm(name)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete film"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Film delete successfully"})
}

func GetAllFilm(context *gin.Context) {
	films, err := model.GetAllFilms()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve films"})
		return
	}

	context.JSON(http.StatusOK, films)
}

func RatingScore(context *gin.Context) {
	filmID := context.Param("filmID")
	ID, err := strconv.ParseUint(filmID, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film ID"})
		return
	}
	averageScore, err := model.CalculateAverageScoreForFilm(uint(ID))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to caculate"})
		return
	}

	context.JSON(http.StatusOK, averageScore)
}
