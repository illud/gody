package router

import (
	"net/http"
	"strings"

	actionsController "github.com/gody-server/app/actions/aplication"
	tokenController "github.com/gody-server/app/token/aplication"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	usersController "github.com/gody-server/app/users/aplication"

	jwtService "github.com/gody-server/adapters/jwt"
	docs "github.com/gody-server/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Token Auth Middleware
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//If one of this paths will be ignore
		if c.FullPath() == "/swagger/*any" ||
			strings.HasPrefix(c.FullPath(), "/gody") ||
			c.FullPath() == "/users/login" ||
			c.FullPath() == "/token/verify" {
		} else {
			// auth := c.GetHeader("Authorization")
			// validateToken := jwt.ValidateToken(auth)

			// token := c.Request.Cookies()
			token := c.GetHeader("Authorization")

			if len(token) != 0 {
				err := jwtService.ValidateToken(token)
				if err != "Ok" {
					c.JSON(403, gin.H{
						"data": "Invalid Token",
					})
					c.Abort()
				}
			} else {
				c.JSON(403, gin.H{
					"data": "No token provided",
				})
				c.Abort()
			}
		}
		c.Next()
	}
}

func Router() *gin.Engine {
	//this sets gin to release mode
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"},
		AllowHeaders:  []string{"Accept", "Authorization", "Content-Type", "Content-Length", "X-CSRF-Token", "Token", "session", "Origin", "Host", "Connection", "Accept-Encoding", "Accept-Language", "X-Requested-With"},
		ExposeHeaders: []string{"X-Total-Count", "Content-Range"},
	}))

	//Token Auth Middleware
	router.Use(TokenAuthMiddleware())

	// Serve static files from the build folder
	router.StaticFS("/gody", http.Dir("./gody"))

	//SWAGGER
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//users
	router.POST("/users/login", usersController.Login)
	router.POST("/users", usersController.CreateUsers)
	router.GET("/users", usersController.GetUsers)
	router.GET("/users/:usersId", usersController.GetOneUsers)
	router.PUT("/users/:usersId", usersController.UpdateUsers)
	router.DELETE("/users/:usersId", usersController.DeleteUsers)

	//actions
	router.POST("/actions/run", actionsController.Run)
	router.POST("/actions", actionsController.CreateActions)
	router.GET("/actions", actionsController.GetActions)
	router.GET("/actions/:actionsId", actionsController.GetOneActions)
	router.PUT("/actions/:actionsId", actionsController.UpdateActions)
	router.DELETE("/actions/:actionsId", actionsController.DeleteActions)

	//token
	router.POST("/token/verify", tokenController.Verify)
	router.POST("/token", tokenController.CreateToken)
	router.GET("/token", tokenController.GetToken)
	router.GET("/token/:tokenId", tokenController.GetOneToken)
	router.PUT("/token/:tokenId", tokenController.UpdateToken)
	router.DELETE("/token/:tokenId", tokenController.DeleteToken)

	return router
}
