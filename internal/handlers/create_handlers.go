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

func (h *Handler) CreateBookHandler(c *gin.Context) {
	var book *models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		h.Logger.Sugar().Error("Falied to bing JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.DB.Create(&book).Error; err != nil {
		h.Logger.Sugar().Error("Falied to create book: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.Logger.Sugar().Info("Successful created reader: ", book)
	c.JSON(http.StatusCreated, book)
}

func (h *Handler) CreateAuthorHandler(c *gin.Context) {
	var author *models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		h.Logger.Sugar().Error("Falied to bing JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.DB.Create(&author).Error; err != nil {
		h.Logger.Sugar().Error("Falied to create author: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.Logger.Sugar().Info("Successful created reader: ", author)
	c.JSON(http.StatusCreated, author)
}

func (h *Handler) CreateLoanHandler(c *gin.Context) {
	var loan *models.Loan
	if err := c.ShouldBindJSON(&loan); err != nil {
		h.Logger.Sugar().Error("Falied to bing JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.DB.Create(&loan).Error; err != nil {
		h.Logger.Sugar().Error("Falied to create loan: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.Logger.Sugar().Info("Successful created reader: ", loan)
	c.JSON(http.StatusCreated, loan)

}
