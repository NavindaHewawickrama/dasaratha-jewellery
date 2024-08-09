package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/NavindaHewawickrama/dasaratha-jewellery/models"
	"github.com/gin-gonic/gin"
	"gituhb.com/NavindaHewawickrama/dasaratha-jewellery/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{
				"Error": "invalid code",
			})
			c.Abort()
			return
		}

		address, err:= ObjectIDFromHex(user_id)
		if err != nil {
			c.IndentedJSON(500, "Internal server error")
		}

		var addresses models.Address

		addresses.Address_id = primitive.NewObjectID()

		if err := c.ShouldBindJSON(&addresses); err != nil {


		address := models.Address{}
		if err := c.BindJSON(&address); err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		match_filter := bson.D{{Key:"$match",Value: bson.D{primitive.E{Key: "_id",Value: address}}}}
		unwind := bson.D{{Key: "$unwind",Value: bson.D{primitive.E{Key: "path",Value: "$address"}}}}






		usert_id, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.IndentedJSON(500, "Internal server error")
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.D{primitive.E{Key: "_id", Value: usert_id}}
		update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "address", Value: address}}}}
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(404, "Wrong command")
			return
		}

		defer cancel()
		ctx.Done()
		c.IndentedJSON(200, "Successfully Added")
	}
}

func EditHomeAddress() gin.HandlerFunc {

}

func EditWOrkAddress() gin.HandlerFunc {

}

func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")

		if user_id == "" {
			c.Header("Contect-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{
				"Error": "invalid search index",
			})
			c.Abort()
			return
		}

		addresses := make([]models.Address, 0)
		usert_id, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.IndentedJSON(500, "Internal server error")
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.D{primitive.E{Key: "_id", Value: usert_id}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(404, "Wrong command")
			return
		}

		defer cancel()
		ctx.Done()
		c.IndentedJSON(200, "Successfully Deleted")

	}
}
