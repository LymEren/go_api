package services

import (
	"go_api/dto"
	"go_api/models"
	"go_api/repository"
)

// We are using insert (in repository)
type DefaultTodoServices struct {
	Repo repository.ToDoRepository
}

type TodoService interface {
	TodoInsert(todo models.Todo) (*dto.TodoDTO, error)
}

func (t DefaultTodoServices) TodoInsert(todo models.Todo) (*dto.TodoDTO, error) {
	var response dto.TodoDTO
	if len(todo.Title) <= 2 {
		response.Status = false
		return &response, nil
	}

	result, err := t.Repo.Insert(todo)

	if err != nil || result == false {
		response.Status = false
		return &response, err
	}

	response = dto.TodoDTO{Status: result}
	return &response, nil
}

//
func NewTodoServise(Repo repository.ToDoRepository) DefaultTodoServices {
	return DefaultTodoServices{Repo: Repo}
}
