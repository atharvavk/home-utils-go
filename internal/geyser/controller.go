package geyser

import (
	"errors"
	"home-utils/internal/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GeyserController struct {
	logger        *zap.Logger
	geyserService GeyserService
}

func NewGeyserController(logger *zap.Logger, geyserService GeyserService) GeyserController {
	return GeyserController{
		logger:        logger,
		geyserService: geyserService,
	}
}

func (gc GeyserController) GetStatus(c *gin.Context) {
	resposne, err := gc.geyserService.GetGeyserStatus()
	if err != nil {
		var badReqErr models.BadReqErr
		if errors.As(err, &badReqErr) {
			c.JSON(400, badReqErr)
			return
		}
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(200, resposne)
}

func (gc GeyserController) DoGeyserAction(c *gin.Context) {
	var req models.GeyserActionRequest

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(400, models.NewBadReqError(models.INVALID_REQUEST))
		return
	}

	resposne, err := gc.geyserService.DoGeyserAction(c.GetHeader("key"), req.TurnGeyserOn)
	if err != nil {
		var badReqErr models.BadReqErr
		if errors.As(err, &badReqErr) {
			c.JSON(400, badReqErr)
			return
		}
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(200, resposne)
}