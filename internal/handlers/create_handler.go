package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skinkvi/cyrs/internal/models"
	"github.com/skinkvi/cyrs/pkg/db"
)

func (h *Handler) CreateReaderHandler(c *gin.Context) {
	var reader *models.Reader
	if err := c.ShouldBindJSON(&reader); err != nil {
		h.Logger.Sugar().Error("Falied to bing JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.DB.Create(&reader).Error; err != nil {
		h.Logger.Sugar().Error("Falied to create reader: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.Logger.Sugar().Info("Successful created reader: ", reader)
	c.JSON(http.StatusCreated, reader)
}
