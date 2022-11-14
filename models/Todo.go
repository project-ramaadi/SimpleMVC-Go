package models

import "errors"

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos = []Todo{
	{ID: 1, Title: "Task 1", Done: false},
	{ID: 2, Title: "Task 2", Done: true},
}

func GetTodos() []Todo {
	return todos
}

func GetTodoById(id int) (*Todo, error) {
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

func UpdateTodoById(id int, newTitle string) (*Todo, error) {
	todo, err := GetTodoById(id)
	if err != nil {
		return nil, err
	}
	todo.Title = newTitle
	return todo, nil
}

func DeleteTodoById(id int) error {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}

	return errors.New("todo not found")
}
