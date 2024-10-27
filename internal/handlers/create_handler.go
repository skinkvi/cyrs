package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skinkvi/cyrs/internal/models"
	"github.com/skinkvi/cyrs/pkg/db"
)

// TODO: нужно подумать как добавить сюда логи потому что я не могу просто передать сюда логгер это будет
// грубой ошибкой так делать. Есть вариант создать просто структуру хендлер и потом просто делать методы
// на ее основе а в ней как раз будет логгер
func CreateReaderHandler(c *gin.Context) {
	var reader *models.Reader
	if err := c.ShouldBindJSON(&reader); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.DB.Create(&reader).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, reader)
}
