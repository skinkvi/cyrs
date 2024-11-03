package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skinkvi/cyrs/internal/models"
	"gorm.io/gorm"
)

func (h *Handler) CreateLoanHandler(c *gin.Context) {
	var loan *models.Loan
	if err := c.ShouldBindJSON(&loan); err != nil {
		h.Logger.Sugar().Error("Falied to bing JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.DB.Create(&loan).Error; err != nil {
		h.Logger.Sugar().Error("Falied to create loan: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.Logger.Sugar().Info("Successful created reader: ", loan)
	c.JSON(http.StatusCreated, loan)

}

func (h *Handler) UpdateLoanHandler(c *gin.Context) {
	id := c.Param("id")

	var loan *models.Loan
	if err := h.DB.First(&loan, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			h.Logger.Sugar().Error("loan not found: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Loan not found"})
		} else {
			h.Logger.Sugar().Error("DataBase error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	var input struct {
		BookID     *uint   `json:"book_id"`
		ReaderID   *uint   `json:"reader_id"`
		LoanDate   *string `json:"loan_date"`
		ReturnDate *string `json:"return_date"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		h.Logger.Sugar().Error("Failed to bing JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.BookID != nil {
		loan.BookID = *input.BookID
	}

	if input.ReaderID != nil {
		loan.ReaderID = *input.ReaderID
	}

	if input.LoanDate != nil {
		loan.LoanDate = *input.LoanDate
	}

	if input.ReturnDate != nil {
		loan.ReturnDate = *input.ReturnDate
	}

	if err := h.DB.Save(&loan).Error; err != nil {
		h.Logger.Sugar().Error("Failed to update loan: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update loan"})
		return
	}

	h.Logger.Sugar().Info("Successfully updated loan", loan)
	c.JSON(http.StatusOK, loan)
}

func (h *Handler) GetLoanHandler(c *gin.Context) {
	var loans []models.Loan
	if err := h.DB.Find(&loans).Error; err != nil {
		h.Logger.Sugar().Error("Failed to get loans: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.Logger.Sugar().Info("Successfully get loans: ", loans)
	c.JSON(http.StatusOK, loans)
}

func (h *Handler) DeleteLoanHandler(c *gin.Context) {
	id := c.Param("id")

	var loan *models.Loan
	if err := h.DB.Where("id = ?", id).Delete(&loan).Error; err != nil {
		h.Logger.Sugar().Error("Failed to delete loan: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.Logger.Info("Successfully delete loan")
	c.JSON(http.StatusOK, "Delete")
}
