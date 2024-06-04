package handler

import (
	"NotificationOfBirthdays"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getAuthorResponse struct {
	Users []NotificationOfBirthdays.Author `json:"users"`
}

func (h *Handler) getUsers(c *gin.Context) {
	users, err := h.services.Profile.GetUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAuthorResponse{
		Users: users,
	})
}
