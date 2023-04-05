package router

import (
	controllers "cs-backend/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Routes() *gin.Engine {
	router := gin.Default()

	organizationRoutes := router.Group("v1/organizations")
	{
		organizationRoutes.GET("/", controllers.GetOrganizations)
		organizationRoutes.GET("/:id/users", controllers.GetOrganizations)
		organizationRoutes.GET("/:id/org-data", controllers.GetOrganizations)
	}

	return router
}
