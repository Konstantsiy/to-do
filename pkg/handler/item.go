package handler

import (
	"github.com/Konstantsiy/todo/pkg/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	var input entity.TodoItem
	if err = c.BindJSON(&input); err != nil {
		NewResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err = h.service.TodoItem.Create(userId, id, input)
	if err != nil {
		NewResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllItems(c *gin.Context) {

}

func (h *Handler) getItemById(c *gin.Context) {

}

func (h *Handler) updateItem(c *gin.Context) {

}

func (h *Handler) deleteItem(c *gin.Context) {

}
