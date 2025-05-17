package handler

import (
	"net/http"
	"strconv"

	todo "github.com/LineCoran/go-api"
	"github.com/gin-gonic/gin"
)


func (h *Handler) createCategory(c *gin.Context) {

	id, ok := c.Get(userCtx)

	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user id is not found")
	}
	var input todo.ExpenseCategory
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CategoryList.CreateOne(id.(int), input)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if id == 0 {
		NewErrorResponse(c, http.StatusBadRequest, "category already exist")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllCategory(c *gin.Context) {
	userId, ok := c.Get(userCtx)

	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user id is not found")
		return
	}

	expenses, err := h.services.CategoryList.GetAllByUserId(userId.(int))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": expenses,
	})

}

func (h *Handler) deleteCategory(c *gin.Context) {

	userId, ok := c.Get(userCtx)

	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user id is not found")
	}

	idStr := c.Param("id")
	categoryId, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id format, must be integer",
		})
		return
	}

	id, err := h.services.CategoryList.DeleteCategory(userId.(int), categoryId)

	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
