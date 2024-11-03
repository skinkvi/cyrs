package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skinkvi/cyrs/internal/models"
	"gorm.io/gorm"
)

func (h *Handler) CreateAuthorHandler(c *gin.Context) {
	var author *models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		h.Logger.Sugar().Error("Falied to bing JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.DB.Create(&author).Error; err != nil {
		h.Logger.Sugar().Error("Falied to create author: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.Logger.Sugar().Info("Successful created reader: ", author)
	c.JSON(http.StatusCreated, author)
}

func (h *Handler) UpdateAuthorHandler(c *gin.Context) {
	name := c.Param("name")

	var author *models.Author
	if err := h.DB.Where("name = ?", name).First(&author).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			h.Logger.Sugar().Error("Author not found: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Author not found"})
		} else {
			h.Logger.Sugar().Error("DataBase error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	var input struct {
		Name        *string `json:"name"`
		BirthDate   *string `json:"birth_date"`
		Nationality *string `json:"natoinality"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		h.Logger.Sugar().Error("Failed to bing JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name != nil {
		author.Name = *input.Name
	}

	if input.BirthDate != nil {
		author.BirthDate = *input.BirthDate
	}

	if input.Nationality != nil {
		author.Nationality = *input.Nationality
	}

	if err := h.DB.Save(&author).Error; err != nil {
		h.Logger.Sugar().Error("Failed to update author: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update author"})
		return
	}

	h.Logger.Sugar().Info("Successfully updated author", author)
	c.JSON(http.StatusOK, author)
}

func (h *Handler) GetAuthorHandler(c *gin.Context) {
	var authors *[]models.Author
	if err := h.DB.Find(&authors).Error; err != nil {
		h.Logger.Sugar().Error("Failed to get authors: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.Logger.Sugar().Info("Successful to get author: ", authors)
	c.JSON(http.StatusOK, authors)
}

func (h *Handler) DeleteAuthorHandler(c *gin.Context) {
	name := c.Param("name")

	var author *models.Author
	if err := h.DB.Where("name = ?", name).Delete(&author).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			h.Logger.Sugar().Error("Author not found: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Author not found"})
		} else {
			h.Logger.Sugar().Error("DataBase error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	h.Logger.Sugar().Info("Successful deleted author: ", author)
	c.JSON(http.StatusOK, "Deleted")
}
