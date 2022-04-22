package handler

import (
	"github.com/gin-gonic/gin"
	shop "local/shop/backend"
	"local/shop/backend/pkg"
	"net/http"
)

func (h *Handler) createUser(c *gin.Context) {
	var request shop.User

	if err := c.BindJSON(&request); err != nil {
		pkg.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateUser(request)
	if err != nil {
		pkg.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{} {
		"id": id,
	})
}
