package controllers

import (
	"context"
	"fmt"
	"net/http"

	//"net/http"

	"tuce/models"

	//"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GodController struct{}

func NewGodController() *GodController {
	return &GodController{}
}

const MONGODB_URI = "mongodb://localhost:27017/godb"

func GetGod() gin.HandlerFunc {
	return func(c *gin.Context) {
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGODB_URI))
		if err != nil {
			panic(err)
		}

		coll := client.Database("godb").Collection("gods")
		searchResult, err := coll.Find(c, bson.D{})
		var gods []models.God
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
			return
		}

		defer searchResult.Close(c)

		if err = searchResult.All(c, &gods); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gods)
	}
}

/* func getGod() {
ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGODB_URI))
if err != nil {
	panic(err)
}
defer func() {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}()
var user models.God
/* newUser := models.God{
	Id:    primitive.NewObjectID(),
	Name:  user.Name,
	Title: user.Title,
} */

//context.IndentedJSON(http.StatusOK, user)
/* 	coll := client.Database("godb").Collection("gods")
result, err := coll.Find(ctx, user)
fmt.Println(result) */

/* 	var user models.User
   	u := models.User{
   		Id:       primitive.NewObjectID(),
   		Name:     user.Name,
   		UserName: user.UserName,
   		Email:    user.Email,
   	}
   	json.NewDecoder(r.Body).Decode(&u)

   	uj, _ := json.Marshal(u)
   	w.Header().Set("content-type", "application/json")
   	w.WriteHeader(http.StatusCreated)
   	fmt.Fprintf(w, "%s/n", uj) */

//}
func CreateGod() gin.HandlerFunc {
	return func(c *gin.Context) {
		var god models.God
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGODB_URI))
		if err != nil {
			panic(err)
		}

		c.BindJSON(&god)

		defer func() {
			if err := client.Disconnect(context.TODO()); err != nil {
				panic(err)
			}
		}()

		newGod := models.God{
			Name:  god.Name,
			Title: god.Title,
		}

		coll := client.Database("godb").Collection("gods")
		result, err := coll.InsertOne(context.TODO(), newGod)
		print(result)
		c.JSON(http.StatusOK, result)
		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	}
}
