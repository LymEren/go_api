package repository

import (
	"context"
	"time"

	"go_api/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepositoryDB struct {
	// With this strutct, we can delete, update etc. in our database
	ToDoCollection *mongo.Collection
}

// We declared an interface for the run
type ToDoRepository interface {
	Insert(todo models.Todo) (bool, error)
}

func (t TodoRepositoryDB) Insert(todo models.Todo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defers the execution of a function until the surrounding function returns
	defer cancel()

	// contex and document (we sent todo to MongoDB)
	result, err := t.ToDoCollection.InsertOne(ctx, todo)

	if result.InsertedID == nil || err != nil {
		//errors.New("Problem todoRepositoryDB")
		return false, err
	}
	return true, nil
}

// main.go collection will come here for client

func NewTodoRepositoryDB(dbClient *mongo.Collection) TodoRepositoryDB {
	return TodoRepositoryDB{ToDoCollection: dbClient}
}
