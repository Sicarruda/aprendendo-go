// https://www.alura.com.br/artigos/criando-uma-simples-aplicacao-web-com-go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

func initDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database("tasker").Collection("tasks")
	fmt.Println("Collection instance created!")
}

func handler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "../static/index.html")
}

// https://levelup.gitconnected.com/build-a-todo-app-in-golang-mongodb-and-react-e1357b4690a6

// CreateTask create task route
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task models.ToDoList
	_ = json.NewDecoder(r.Body).Decode(&task)
	// fmt.Println(task, r.Body)
	insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}

func main() {
	initDB()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}
