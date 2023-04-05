package router

import (
	controllers "cs-backend/controllers"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Routes() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	organizationRoutes := router.Group("v1/organizations")
	{
		organizationRoutes.GET("/", controllers.GetOrganizations)
		organizationRoutes.GET("/:id/users", controllers.GetAllUsersInAnOrganization)
		organizationRoutes.GET("/:id/org-data", controllers.GetOrgSpecificData)
	}

	return router
}
