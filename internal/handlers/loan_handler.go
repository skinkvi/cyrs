package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skinkvi/cyrs/internal/models"
	"gorm.io/gorm"
)

// CreateLoanHandler godoc
// @Summary      Создание нового займа
// @Description  Создает нового займа в системе
// @Tags         loans
// @Accept       json
// @Produce      json
// @Param        loan  body      models.Loan  true  "Займ"
// @Success      201     {object}  models.Loan
// @Failure      400     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Router       /api/loans [post]
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

// UpdateLoanHandler godoc
// @Summary      Обновление информации об займе
// @Description  Обновляет данные существующего займа
// @Tags         loans
// @Accept       json
// @Produce      json
// @Param        id      path      int               true  "ID займа"
// @Param        loans  body      models.Loan     true  "Данные займа"
// @Success      200     {object}  models.Loan
// @Failure      400     {object}  map[string]string
// @Failure      404     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Router       /api/loans/{id} [put]
func (h *Handler) UpdateLoanHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		h.Logger.Sugar().Error("Invalid author ID: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}

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

// GetLoanHandler godoc
// @Summary      Получение списка займов
// @Description  Возвращает список всех завмов
// @Tags         loans
// @Produce      json
// @Success      200  {array}   models.Loan
// @Failure      500  {object}  map[string]string
// @Router       /api/loans [get]
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

// DeleteLoanHandler godoc
// @Summary      Удаление займа
// @Description  Удаляет займ по ID
// @Tags         loans
// @Produce      json
// @Param        id   path      int  true  "ID займа"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/loans/{id} [delete]
func (h *Handler) DeleteLoanHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		h.Logger.Sugar().Error("Invalid author ID: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}

	var loan *models.Loan
	if err := h.DB.Delete(&loan, id).Error; err != nil {
		h.Logger.Sugar().Error("Failed to delete loan: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.Logger.Info("Successfully delete loan")
	c.JSON(http.StatusOK, gin.H{"message": "Delete"})
}
