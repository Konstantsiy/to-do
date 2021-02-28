package handler

import (
	"github.com/Konstantsiy/todo/pkg/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input entity.User

	if err := c.BindJSON(&input); err != nil {
		NewResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		NewResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string
	Password string
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		NewResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		NewResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
