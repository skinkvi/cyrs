package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skinkvi/cyrs/internal/models"
	"gorm.io/gorm"
)

func (h *Handler) CreateBookHandler(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		h.Logger.Sugar().Error("Failed to bind JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.DB.Create(&book).Error; err != nil {
		h.Logger.Sugar().Error("Failed to create book: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.Logger.Sugar().Info("Successfully created book: ", book)
	c.JSON(http.StatusCreated, book)
}

func (h *Handler) UpdateBookHandler(c *gin.Context) {
	title := c.Param("title")

	var book *models.Book
	if err := h.DB.Where("title = ?", title).First(&book).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			h.Logger.Sugar().Error("Book not found: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Book not found"})
		} else {
			h.Logger.Sugar().Error("DataBase error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	// todo: доделать тут
}
