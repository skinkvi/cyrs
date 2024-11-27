package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skinkvi/cyrs/internal/models"
	"gorm.io/gorm"
)

// CreateAuthorHandler godoc
// @Summary      Создание нового автора
// @Description  Создаёт нового автора в системе
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param        author  body      models.Author  true  "Автор"
// @Success      201     {object}  models.Author
// @Failure      400     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Router       /api/authors [post]
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

// UpdateAuthorHandler godoc
// @Summary      Обновление информации об авторе
// @Description  Обновляет данные существующего автора
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param        id      path      int               true  "ID автора"
// @Param        author  body      models.Author     true  "Данные автора"
// @Success      200     {object}  models.Author
// @Failure      400     {object}  map[string]string
// @Failure      404     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Router       /api/authors/{id} [put]
func (h *Handler) UpdateAuthorHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		h.Logger.Sugar().Error("Invalid author ID: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}

	var author *models.Author
	if err := h.DB.First(&author, id).Error; err != nil {
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

// GetAuthorHandler godoc
// @Summary      Получение списка авторов
// @Description  Возвращает список всех авторов
// @Tags         authors
// @Produce      json
// @Success      200  {array}   models.Author
// @Failure      500  {object}  map[string]string
// @Router       /api/authors [get]
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

// DeleteAuthorHandler godoc
// @Summary      Удаление автора
// @Description  Удаляет автора по ID
// @Tags         authors
// @Produce      json
// @Param        id   path      int  true  "ID автора"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/authors/{id} [delete]
func (h *Handler) DeleteAuthorHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		h.Logger.Sugar().Error("Invalid author ID: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}

	if err := h.DB.Delete(&models.Author{}, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			h.Logger.Sugar().Error("Author not found: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Author not found"})
		} else {
			h.Logger.Sugar().Error("DataBase error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	h.Logger.Sugar().Info("Successful deleted author")
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
