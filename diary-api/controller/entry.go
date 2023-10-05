package controller

import (
	"diary-api/database"
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

	context.JSON(http.StatusCreated, gin.H{"entry": savedUser})
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
	id := context.Param("ID")
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
	id := context.Param("ID")
	entryID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid entry ID"})
		return
	}

	entry, err := model.FindEntryById(uint(entryID))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get entry"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"get successfully": entry})
}

func CountEntry(context *gin.Context) {
	id := context.Param("user_id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid entry ID"})
		return
	}

	entryCount, countErr := Count(uint(userID))

	if countErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count entries"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"count": entryCount})
}

func Count(userID uint) (int64, error) {
	var count int64
	if err := database.Database.Model(&model.Entry{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
