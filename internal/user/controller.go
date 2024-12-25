package user

import (
	"errors"
	"home-utils/internal/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	logger      *zap.Logger
	userService UserService
}

func NewUserController(logger *zap.Logger, userService UserService) UserController {
	return UserController{
		logger:      logger,
		userService: userService,
	}
}

func (uc UserController) ValidateKey(c *gin.Context) {
	userKey := c.GetHeader("key")
	response, err := uc.userService.GetUser(userKey)
	if err != nil {
		var badReqErr models.BadReqErr
		if errors.As(err, &badReqErr) {
			c.JSON(400, badReqErr)
			return
		}
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, response)
}

func (uc UserController) ValidateKeyMiddleware(c *gin.Context) {
	userKey := c.GetHeader("key")
	_, err := uc.userService.GetUser(userKey)
	if err != nil {
		var badReqErr models.BadReqErr
		if errors.As(err, &badReqErr) {
			c.AbortWithStatusJSON(400, badReqErr)
			return
		}
		c.AbortWithStatusJSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.Next()
}
