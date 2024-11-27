package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skinkvi/cyrs/internal/models"
	"gorm.io/gorm"
)

// @Summary      Создание новой книги
// @Description  Создает новую книгу в системе и записывает ее в базу данных
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body      models.Book  true  "Книга"
// @Success      201     {object}  models.Book
// @Failure      400     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Router       /api/books [post]
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

// UpdateBookHandler godoc
// @Summary      Обновление информации об книге
// @Description  Обновляет данные существующей книги
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id      path      int               true  "ID книги"
// @Param        book  body      models.Book     true  "Данные книги"
// @Success      200     {object}  models.Book
// @Failure      400     {object}  map[string]string
// @Failure      404     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Router       /api/books/{id} [put]
func (h *Handler) UpdateBookHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		h.Logger.Sugar().Error("Invalid author ID: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}

	var book *models.Book
	if err := h.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			h.Logger.Sugar().Error("Book not found: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Book not found"})
		} else {
			h.Logger.Sugar().Error("DataBase error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	var input struct {
		Title         *string `json:"title"`
		AuthorID      *uint   `json:"author_id"`
		PublishedDate *string `json:"published_date"`
		ISBN          *string `json:"isbn"`
		Available     *bool   `json:"available"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		h.Logger.Sugar().Error("Failed to bing JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Title != nil {
		book.Title = *input.Title
	}

	if input.AuthorID != nil {
		book.AuthorID = *input.AuthorID
	}

	if input.PublishedDate != nil {
		book.PublishedDate = *input.PublishedDate
	}

	if input.ISBN != nil {
		book.ISBN = *input.ISBN
	}

	if input.Available != nil {
		book.Available = *input.Available
	}

	if err := h.DB.Save(&book).Error; err != nil {
		h.Logger.Sugar().Error("Failed to update book: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	h.Logger.Sugar().Info("Successfully updated book", book)
	c.JSON(http.StatusOK, book)
}

// GetBookHandler godoc
// @Summary      Получение списка книг
// @Description  Возвращает список всех книг
// @Tags         books
// @Produce      json
// @Success      200  {array}   models.Book
// @Failure      500  {object}  map[string]string
// @Router       /api/books [get]
func (h *Handler) GetBookHandler(c *gin.Context) {
	var books []models.Book
	if err := h.DB.Find(&books).Error; err != nil {
		h.Logger.Sugar().Error("Failed to find book: ", books)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.Logger.Sugar().Info("Successfully find book: ", books)
	c.JSON(http.StatusOK, books)
}

// DeleteBookHandler godoc
// @Summary      Удаление книги
// @Description  Удаляет книгу по ID
// @Tags         books
// @Produce      json
// @Param        id   path      int  true  "ID книги"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/books/{id} [delete]
func (h *Handler) DeleleBookHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		h.Logger.Sugar().Error("Invalid author ID: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}

	var book *models.Book
	if err := h.DB.Delete(&book, id).Error; err != nil {
		h.Logger.Sugar().Error("Failed delete book: ", err.Error())
		return
	}

	h.Logger.Info("Successfully deleted book")
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
