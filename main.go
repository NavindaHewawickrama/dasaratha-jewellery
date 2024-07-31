package main

import (
	"github.com/NavindaHewawickrama/dasaratha-jewellery/controllers"
	"github.com/NavindaHewawickrama/dasaratha-jewellery/database"
	"github.com/NavindaHewawickrama/dasaratha-jewellery/middleware"
	"github.com/NavindaHewawickrama/dasaratha-jewellery/routes"
	"github.com/gin-gonic/gin"
)

func main (){
	port:=os.Getenv("PORT")
	if port==""{
		port="8080"
	}	

	app:= controllers.NewApplication{database.ProductData(database.Client,"Products"),database.UserData(database.Client,"Users")}

	router := gin.New()
	router.Use(gin.Logger())\

	routes.UserRoutes(router)

	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/buynow", app.BuyNow())

	log.Fatal(router.Run(":"+port))

}