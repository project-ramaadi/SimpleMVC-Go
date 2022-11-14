package models

import "errors"

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos []Todo

func GetTodos() []Todo {
	return todos
}

func GetTodoById(id string) (*Todo, error) {
	for i, todo := range todos {
		if todo.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func CreateNewTodo(todo Todo) Todo {
	todos = append(todos, todo)
	return todo
}

func UpdateTodoTitleById(id string, newTitle string) (*Todo, error) {
	todo, err := GetTodoById(id)
	if err != nil {
		return nil, err
	}
	todo.Title = newTitle
	return todo, nil
}

func MarkTodoAsDoneById(id string, done bool) (*Todo, error) {
	todo, err := GetTodoById(id)
	if err != nil {
		return nil, err
	}
	todo.Done = done
	return todo, nil
}

func DeleteTodoById(id string) error {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}

	return errors.New("todo not found")
}
