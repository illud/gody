package router

import (
	"log"
	"net/http"
	"strings"

	executionhistoryController "github.com/gody-server/app/executionhistory/aplication"

	actionsController "github.com/gody-server/app/actions/aplication"
	tokenController "github.com/gody-server/app/token/aplication"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	usersController "github.com/gody-server/app/users/aplication"

	jwtService "github.com/gody-server/adapters/jwt"
	configFile "github.com/gody-server/config"
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
			c.FullPath() == "/token/verify" ||
			c.FullPath() == "/config" {
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

	configFile, err := configFile.ConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     configFile.AllowOrigins,
		AllowMethods:     []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "Content-Length", "X-CSRF-Token", "Token", "session", "Origin", "Host", "Connection", "Accept-Encoding", "Accept-Language", "X-Requested-With"},
		ExposeHeaders:    []string{"X-Total-Count", "Content-Range"},
		AllowCredentials: true,
	}))

	//Token Auth Middleware
	// router.Use(TokenAuthMiddleware())

	// Serve static files from the build folder
	router.StaticFS("/gody", http.Dir("./gody"))

	//SWAGGER
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.GET("/config", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": configFile,
		})
	})

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

	//executionhistory
	router.POST("/execution-history", executionhistoryController.CreateExecutionhistory)
	router.GET("/execution-history/all/by-action/:actionId", executionhistoryController.GetExecutionhistory)
	router.GET("/execution-history/:executionHistoryId", executionhistoryController.GetOneExecutionhistory)
	router.PUT("/execution-history/:executionHistoryId", executionhistoryController.UpdateExecutionhistory)
	router.DELETE("/execution-history/:executionHistoryId", executionhistoryController.DeleteExecutionhistory)

	return router
}
