package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skinkvi/cyrs/internal/models"
	"gorm.io/gorm"
)

// CreateReaderHandler godoc
// @Summary      Создание нового читателя
// @Description  Создает нового читателя в системе
// @Tags         readers
// @Accept       json
// @Produce      json
// @Param        reader  body      models.Reader  true  "Читатель"
// @Success      201     {object}  models.Reader
// @Failure      400     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Router       /api/readers [post]
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

// UpdateReaderHandler godoc
// @Summary      Обновление информации об читателе
// @Description  Обновляет данные существующего читателя
// @Tags         readers
// @Accept       json
// @Produce      json
// @Param        id      path      int               true  "ID читателя"
// @Param        reader  body      models.Reader     true  "Данные читателя"
// @Success      200     {object}  models.Reader
// @Failure      400     {object}  map[string]string
// @Failure      404     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Router       /api/readers/{id} [put]
func (h *Handler) UpdateReaderHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		h.Logger.Sugar().Error("Invalid author ID: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}

	var reader *models.Reader
	if err := h.DB.First(&reader, id).Error; err != nil {
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

// GetReaderHandler godoc
// @Summary      Получение списка читателей
// @Description  Возвращает список всех читателей
// @Tags         readers
// @Produce      json
// @Success      200  {array}   models.Reader
// @Failure      500  {object}  map[string]string
// @Router       /api/readers [get]
func (h *Handler) GetReaderHandler(c *gin.Context) {
	var readers []models.Reader
	if err := h.DB.Find(&readers).Error; err != nil {
		h.Logger.Sugar().Error("Failed to get readers: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.Logger.Sugar().Info("Successfully to get readers: ", readers)
	c.JSON(http.StatusOK, readers)
}

// DeleteReaderHandler godoc
// @Summary      Удаление читателя
// @Description  Удаляет читателя по ID
// @Tags         readers
// @Produce      json
// @Param        id   path      int  true  "ID читателя"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/readers/{id} [delete]
func (h *Handler) DeleteReaderHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		h.Logger.Sugar().Error("Invalid author ID: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}

	var reader *models.Reader
	if err := h.DB.Delete(&reader, id).Error; err != nil {
		h.Logger.Sugar().Error("Failed to detele reader: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.Logger.Info("Successfully reader deleted")
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
