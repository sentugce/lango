package main

import (
	"errors"
	//"fmt"
	"net/http"
	//"time"

	//"tuce/controllers"
	"tuce/routes"

	"github.com/gin-gonic/gin"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	Id        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{Id: "1", Item: "clean room", Completed: false},
	{Id: "2", Item: "read book", Completed: false},
	{Id: "3", Item: "record video", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo Todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := selectGetTodo(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := selectGetTodo(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func selectGetTodo(id string) (*Todo, error) {
	for i, t := range todos {
		if t.Id == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func main() {
	router := gin.Default()
	routes.UserRoute(router)
	//controllers.GetGod()
	/* routes.UserRoute(router)
	getGod()
	//controllers.CreateGod() */
	//router.GET("/todos", controllers.getGod())
	router.Run("localhost:9090")
	/* 	router.GET("/todos/:id", getTodo)
	   	router.PATCH("/todos/:id", toggleTodoStatus)
	   	router.POST("/todos", addTodo)
	   	router.Run("localhost:9090") */

}
