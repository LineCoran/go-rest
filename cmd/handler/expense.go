package handler

import (
	"net/http"
	"strconv"

	todo "github.com/LineCoran/go-api"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createExpense(c *gin.Context) {

	id, ok := c.Get(userCtx)

	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user id is not found")
	}
	var input todo.Expense
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.ExpenseList.Create(id.(int), input)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllExpense(c *gin.Context) {
	userId, ok := c.Get(userCtx)

	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user id is not found")
		return
	}

	expenses, err := h.services.ExpenseList.GetAllByUserId(userId.(int))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": expenses,
	})

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
	deletedId, err := h.services.ExpenseList.Delete(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "failed to delete expense by id")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": deletedId,
	})
}

func (h *Handler) updateExpenseById(c *gin.Context) {
	expenseIdStr := c.Param("id")
	var input todo.Expense

	expenseId, err := strconv.Atoi(expenseIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id format, must be integer",
		})
		return
	}

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	expense, err := h.services.ExpenseList.Update(expenseId, input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, expense)
}
