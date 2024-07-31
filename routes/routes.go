package routes

import (
	"github.com/NavindaHewawickrama/dasaratha-jewellery/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouters(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewverAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
}
