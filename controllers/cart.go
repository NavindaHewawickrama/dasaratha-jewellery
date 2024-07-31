package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gituhb.com/NavindaHewawickrama/dasaratha-jewellery/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct{
	productCollection *mongo.Collection
	userCollaction *mongo.Collection
}

func NewApplication(productCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		productCollection: productCollection,
		userCollection: userCollection 
	}
}

func (app *Application) AddToCart() gin.Handler {
	return func(c *gin.Context){
		productQueryID := c.Query("id")
		if productQueryID == ""{
			log.Println("product is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}

		userQueryID := c.Query("userID")
		if userQueryID == ""{
			log.Println("user is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return

		}

		productID,err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil{
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		err = database.AddProductToCart(ctx, app.productCollection, productID,userQueryID)
		if err != nil{
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(200,"Successfully Added to cart")
	}
}

func (app *Application) RemoveItem() gin.Handlerfunc {
	return func(c *gin.Context){
		productQueryID := c.Query("id")
		if productQueryID == ""{
			log.Println("product is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}

		userQueryID := c.Query("userID")
		if userQueryID == ""{
			log.Println("user is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return

		}

		productID,err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil{
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		err = database.RemoveCartItem(ctx, app.productCollection, app.userCollaction, productID,userQueryID)
		if err != nil{
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(200,"Successfully Removed from cart")
	}
}

func GetItemFromCart() gin.Handlerfunc {
	
}

func BuyFromCart() gin.Handlerfunc {

}

func (app *Application) InstantBuy() gin.Handlerfunc {
	return func(c *gin.Context){
		productQueryID := c.Query("id")
		if productQueryID == ""{
			log.Println("product is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}

		userQueryID := c.Query("userID")
		if userQueryID == ""{
			log.Println("user is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return

		}

		productID,err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil{
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		err = database.InstantBuyer(ctx, app.productCollection, app.userCollaction, productID,userQueryID)
		if err != nil{
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(200,"Successfully Placed the order and Bought the product")
	}
}
