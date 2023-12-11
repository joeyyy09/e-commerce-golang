package routes

import (
	"github.com/joeyyy09/e-commerce-golang/controllers"
	"github.com/gin-gonic/gin"
)

func userRoutes(incomingRoutes "gi.Engine")
{
	incomingRoutes.POST("/users/signup", controllers.Signup())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.GET("/admin/addproduct",	controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview",controllers.SearchProduct())
	incomingRoutes.GET("/users/search",controllers.SearchProductbyQuery())
}
