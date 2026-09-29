package api

import (
	"github.com/gin-gonic/gin"
	controller "gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/mock-api-server/api/controllers"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/mock-api-server/internal/configs"
)

type GinMapping struct {
	applicationInstanceController *controller.ApplicationInstanceController
	config                        *configs.Config
	router                        *gin.Engine
}

func NewGinMapping(config *configs.Config) *GinMapping {
	return &GinMapping{
		router:                        gin.Default(),
		applicationInstanceController: controller.NewApplicationInstanceController(),
		config:                        config,
	}
}

func (gm *GinMapping) Run() error {
	err := gm.create()
	if err != nil {
		return err
	}

	return gm.router.Run(":" + gm.config.HTTP.Port)
}

func Healthz(c *gin.Context) {
	c.JSON(200, gin.H{"status": "OK"})
}

func (gm *GinMapping) create() error {

	err := gm.router.SetTrustedProxies(nil)

	if err != nil {
		return err
	}

	gm.router.GET("/healthz", Healthz)
	gm.router.Use(func(c *gin.Context) {
		if c.GetHeader("PRIVATE-TOKEN") != "test-token" {
			c.AbortWithStatusJSON(401, gin.H{"error": "test token required"})
			return
		}
		c.Next()
	})

	v3 := gm.router.Group("/v3/applicationInstance")

	v3.GET("", gm.applicationInstanceController.List)
	v3.PUT("/game/:applicationId/empty/allocate", gm.applicationInstanceController.Create)
	v3.GET("/:instanceId", gm.applicationInstanceController.Get)
	v3.PUT("/:instanceId", gm.applicationInstanceController.Update)
	v3.POST("/:instanceId/restart", gm.applicationInstanceController.Restart)
	gm.router.GET("/_test/state", gm.applicationInstanceController.State)
	gm.router.POST("/_test/allocation", gm.applicationInstanceController.SetAllocationBehavior)
	gm.router.POST("/_test/instances/:instanceId/status", gm.applicationInstanceController.SetStatus)

	return nil
}
