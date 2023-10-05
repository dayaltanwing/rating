package controller

import (
	"diary-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddVote(context *gin.Context) {
	var input model.Vote

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
