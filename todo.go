package todo

import (
	"fmt"
)

func CreateTodo(title string) string {
	message := fmt.Sprintf("Creating a new todo with title: %s\n", title)
	return message
}

func DeleteTodo(title string) string {
	message := fmt.Sprintf("Deleting todo with title: %s\n", title)
	return message
}

func UpdateTodo(oldTitle, newTitle string) string {
	message := fmt.Sprintf("Updating todo from title: %s to new title: %s\n", oldTitle, newTitle)
	return message
}
