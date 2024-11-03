package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/skinkvi/cyrs/internal/handlers"
)

func SetupRoutes(router *gin.Engine, h *handlers.Handler) {
	api := router.Group("/api")
	{
		books := api.Group("/books")
		{
			books.POST("/", h.CreateBookHandler)
			books.PUT("/:id", h.UpdateBookHandler)
			books.GET("/", h.GetBookHandler)
			books.DELETE("/:id", h.DeleleBookHandler)
		}

		authors := api.Group("/authors")
		{
			authors.POST("/", h.CreateAuthorHandler)
			authors.PUT("/:id", h.UpdateAuthorHandler)
			authors.GET("/", h.GetAuthorHandler)
			authors.DELETE("/:id", h.DeleteAuthorHandler)
		}

		readers := api.Group("/readers")
		{
			readers.POST("/", h.CreateReaderHandler)
			readers.PUT("/:id", h.UpdateReaderHandler)
			readers.GET("/", h.GetReaderHandler)
			readers.DELETE("/:id", h.DeleteReaderHandler)
		}

		loans := api.Group("/loans")
		{
			loans.POST("/", h.CreateLoanHandler)
			loans.PUT("/:id", h.UpdateLoanHandler)
			loans.GET("/", h.GetLoanHandler)
			loans.DELETE("/:id", h.DeleteLoanHandler)
		}
	}
}
