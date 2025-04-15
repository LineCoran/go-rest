package handler

import (
	"net/http"

	todo "github.com/LineCoran/go-api"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createExpense(c *gin.Context) {
	var input todo.Expense
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	h.services.Create(123, input)
}

func (h *Handler) getAllExpense(c *gin.Context) {

}

func (h *Handler) getExpenseById(c *gin.Context) {

}

func (h *Handler) deleteExpenseById(c *gin.Context) {

}

func (h *Handler) updateExpenseById(c *gin.Context) {

}
