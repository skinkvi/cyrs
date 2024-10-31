package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skinkvi/cyrs/internal/models"
	"gorm.io/gorm"
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

	if err := h.DB.Create(&reader).Error; err != nil {
		h.Logger.Sugar().Error("Falied to create reader: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.Logger.Sugar().Info("Successful created reader: ", reader)
	c.JSON(http.StatusCreated, reader)
}

func (h *Handler) UpdateReaderHandler(c *gin.Context) {
	name := c.Param("name")

	var reader *models.Reader
	if err := h.DB.Where("name = ?", name).First(&reader).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			h.Logger.Sugar().Error("Reader not found: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Reader not found"})
		} else {
			h.Logger.Sugar().Error("DataBase error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	var input struct {
		Name           *string `json:"name"`
		Email          *string `json:"email"`
		MemberShipDate *string `json:"member_ship_date"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		h.Logger.Sugar().Error("Failed to bing JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name != nil {
		reader.Name = *input.Name
	}

	if input.Email != nil {
		reader.Email = *input.Email
	}

	if input.MemberShipDate != nil {
		reader.MembershipDate = *input.MemberShipDate
	}

	if err := h.DB.Save(&reader).Error; err != nil {
		h.Logger.Sugar().Error("Failed to update reader: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update reader"})
		return
	}

	h.Logger.Sugar().Info("Successfully updated reader", reader)
	c.JSON(http.StatusOK, reader)
}
