package handler

import (
	"github.com/LineCoran/go-api/cmd/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/exist", h.checkExistByUserName)
	}

	api := router.Group("/api", h.authMiddleware)

	{
		expenses := api.Group("/expenses")
		{
			expenses.POST("/", h.createExpense)
			expenses.GET("/", h.getAllExpense)
			expenses.GET("/:id", h.getExpenseById)
			expenses.DELETE("/:id", h.deleteExpenseById)
			expenses.PUT("/:id", h.updateExpenseById)
		}

		category := api.Group("/category")
		{
			category.POST("/", h.createCategory)
			category.GET("/", h.getAllCategory)
			category.DELETE("/:id", h.deleteCategory)
			// category.DELETE("/:id", h.deleteExpenseById)
			// category.PUT("/:id", h.updateExpenseById)
		}
	}

	return router
}
