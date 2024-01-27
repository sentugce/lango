package main

/* import (
	"context"
	"fmt"
	"time"

	"tuce/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGODB_URI = "mongodb://localhost:27017/godb"

func main() {
	saveUser("yp", "qwe", "asd@asd.com")
	pagesSave("ey", "tuc", "kullanıcı", "/users")
}

func saveUser(name string, userName string, email string) {
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

	coll := client.Database("godb").Collection("users")
	newUser := models.User{
		Name:     name,
		UserName: userName,
		Email:    email,
	}
	result, err := coll.InsertOne(ctx, newUser)
	fmt.Println(result)
}

func pagesSave(name string, userName string, desc string, uri string) {
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
	colls := client.Database("godb").Collection("pages")
	ey := models.Page{
		Name:     name,
		UserName: userName,
		Desc:     desc,
		URI:      uri,
	}

	result, err := colls.InsertOne(ctx, ey)
	if err != nil {
		panic(err)
	}
	fmt.Print(result)

} */
