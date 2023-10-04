package controller

import (
	"context"
	"diary-api/database"
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

	savedEntry, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
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

func GetFilmRatingScore(context *gin.Context) {
	filmIDParam := context.Param("filmID")

	filmID, err := strconv.ParseUint(filmIDParam, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film ID"})
		return
	}

	var film model.Film

	if err := database.Database.First(&film, filmID).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Film not found"})
		return
	}

	ratingScore := calculateRatingScore(film)

	context.JSON(http.StatusOK, gin.H{"rating_score": ratingScore})
}

func calculateRatingScore(ratingsParam string) float64 {
    // Split the ratingsParam string into individual ratings
    ratings := strings.Split(ratingsParam, ",")
    
    if len(ratings) == 0 {
        return 0.0
    }

    totalRating := 0.0
    for _, ratingStr := range ratings {
        // Convert each rating string to a float64
        rating, err := strconv.ParseFloat(ratingStr, 64)
        if err != nil {
            // Handle parsing error if needed
            return 0.0
        }
        totalRating += rating
    }

    averageRating := totalRating / float64(len(ratings))

    return averageRating
}

