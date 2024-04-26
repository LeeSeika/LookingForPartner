package routes

import (
	"github.com/gin-gonic/gin"
	"lookingforpartner/common/middleware"
	"lookingforpartner/pkg/logger"
	"lookingforpartner/service/user/api/internal/config"
	"lookingforpartner/service/user/api/internal/handler"
	"lookingforpartner/service/user/api/internal/svc"
	"net/http"
)

func SetupRouter(c *config.Config, svcCtx *svc.ServiceContext) *gin.Engine {
	if c.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(logger.GinLogger(), logger.GinRecovery(true))
	engine.Use(middleware.InjectSvcCtxMiddleware(svcCtx))

	engine.LoadHTMLFiles("./templates/index.html")
	engine.Static("/static", "./static")

	engine.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	engine.GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "hello")
	})

	group := engine.Group("/api/v1")

	// user service
	group.POST("/signup", handler.SingUpHandler)
	//group.POST("/login", httpHandler.LoginHandler)
	//group.POST("/refresh", httpHandler.RefreshTokenHandler)
	//
	group.Use(middleware.JWTAuthMiddleware())

	return engine
}
