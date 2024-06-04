package handler

import (
	"NotificationOfBirthdays"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createSubscription(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input NotificationOfBirthdays.Subscription
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	subscriptionId, err := h.services.Subscription.CreateSubscription(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"subscriptionId": subscriptionId,
	})
}

type BirthdayUserId struct {
	BirthdayUserId int `json:"birthday_user_id" binding:"required"`
}

func (h *Handler) deleteSubscription(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input BirthdayUserId
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Subscription.DeleteSubscription(userId, input.BirthdayUserId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
