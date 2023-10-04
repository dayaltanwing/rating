package controller

import (
	"diary-api/helper"
	"diary-api/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddEntry(context *gin.Context) {
	var input model.Entry

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)

	input.UserID = user.ID
	savedUser, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func GetAllEntries(context *gin.Context) {
	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user.Entries})
}

func DeleteEntry(context *gin.Context) {
	id := context.Param("entryID")
	entryID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid entry ID"})
		return
	}

	entry, err := model.DeleteEntry(uint(entryID))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete entry"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted successfully": entry})
}

func GetEntrybyId(context *gin.Context) {
	id := context.Param("entryID")
	entryID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid entry ID"})
		return
	}

	entry, err := model.FindEntryById(uint(entryID))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete entry"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted successfully": entry})
}

func CountEntry(context *gin.Context) {
	id := context.Param("user_id")
	entryID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid entry ID"})
		return
	}

	var entryCount int64

	entryCount, countErr := model.Count(uint(entryID))

	if countErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count entries"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"count": entryCount})
}
