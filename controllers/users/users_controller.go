package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nickx720/bookstore_users-api/domain/users"
	"github.com/nickx720/bookstore_users-api/services"
	"github.com/nickx720/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	/* shouldBindJson replaces the following
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// todo handle Error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		// todo handle json error
		return
	} */
	if err := c.ShouldBindJSON(&user); err != nil {
		// Todo handle json error
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// Todo handle user creation error
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
