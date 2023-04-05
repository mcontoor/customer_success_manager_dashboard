package main

import (
	controllers "cs-backend/controllers"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Routes() gin.Router {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	organizationRoutes := router.Group("v1/organizations")
	{
		organizationRoutes.use("/", controllers.GetOrganizations)
		organizationRoutes.use("/:id/users", controllers.GetOrganizations)
		organizationRoutes.use("/:id/org-data", controllers.GetOrganizations)

	}

	return router
}
