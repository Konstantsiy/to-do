package handler

import (
	"github.com/Konstantsiy/todo/pkg/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input entity.TodoList
	if err := c.BindJSON(&input); err != nil {
		NewResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.TodoList.Create(userId, input)
	if err != nil {
		NewResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []entity.TodoList
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	lists, err := h.service.TodoList.GetAll(userId)
	if err != nil {
		NewResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	list, err := h.service.TodoList.GetById(userId, id)
	if err != nil {
		NewResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.TodoList.Delete(userId, id)
	if err != nil {
		NewResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	var input entity.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		NewResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.Update(userId, id, input); err != nil {
		NewResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
