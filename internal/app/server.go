package app

import (
	"fmt"
	"home-utils/internal/geyser"
	"home-utils/internal/user"

	"github.com/gin-gonic/gin"
)

type geyserContext struct {
	controller geyser.GeyserController
	service    geyser.GeyserService
}

type userContext struct {
	controller user.UserController
	service    user.UserService
}

func initGeyserContext(appCtx AppContext) geyserContext {
	service := geyser.NewGeyserService(appCtx.Logger, appCtx.Sql)
	controller := geyser.NewGeyserController(appCtx.Logger, service)
	return geyserContext{
		controller: controller,
		service:    service,
	}
}

func initUserContext(appCtx AppContext) userContext {
	service := user.NewUserService(appCtx.Logger, appCtx.Sql)
	return userContext{
		controller: user.NewUserController(appCtx.Logger, service),
		service:    service,
	}
}

func CreateAndStartServer(appCtx AppContext) {
	router := gin.Default()
	userCtx := initUserContext(appCtx)
	geyserCtx := initGeyserContext(appCtx)

	// user APIs
	userGroup := router.Group("/user")
	userGroup.GET("", userCtx.controller.ValidateKey)

	// geyser APIs
	geyserGroup := router.Group("/geyser", userCtx.controller.ValidateKeyMiddleware)
	geyserGroup.GET("/status", geyserCtx.controller.GetStatus)
	geyserGroup.POST("/action", geyserCtx.controller.DoGeyserAction)
  geyserGroup.GET("/history", geyserCtx.controller.GetGeyserHistory)

	err := router.Run(fmt.Sprintf("0.0.0.0:%d", appCtx.ServerPort))
	if err != nil {
		panic(err)
	}
}
