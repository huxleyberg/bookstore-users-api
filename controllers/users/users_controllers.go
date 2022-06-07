package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huxleyberg/bookstore-users-api/domain/users"
	"github.com/huxleyberg/bookstore-users-api/services"
	"github.com/huxleyberg/bookstore-users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.RestErr{
			Message: "invalide json body",
			Code:    http.StatusBadRequest,
			Error:   "bad_request",
		}
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	res, saveErr := services.CreateUser(&user)
	if saveErr != nil {
		//handle error
		return
	}

	c.JSON(http.StatusCreated, res)

}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
