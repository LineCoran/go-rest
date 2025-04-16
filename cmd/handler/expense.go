package handler

import (
	"net/http"
	"strconv"

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
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id format, must be integer",
		})
		return
	}
	expense, err := h.services.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get expense: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, expense)
}

func (h *Handler) deleteExpenseById(c *gin.Context) {
	id := c.Param("id")
	h.services.Delete(id)
}

func (h *Handler) updateExpenseById(c *gin.Context) {

}
