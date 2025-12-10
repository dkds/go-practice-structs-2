package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println("Text: ", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	jsonContent, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, jsonContent, 0644)
	if err != nil {
		return nil
	}
	return nil
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Content cannot be empty")
	}
	return Todo{
		Text: content,
	}, nil
}
