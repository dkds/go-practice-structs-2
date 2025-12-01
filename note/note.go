package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

func (note Note) Display() {
	fmt.Println("Title: ", note.Title)
	fmt.Println("Content: ", note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	jsonContent, err := json.Marshal(note)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, jsonContent, 0644)
	if err != nil {
		return nil
	}
	return nil
}

func New(title string, content string) (Note, error) {
	if title == "" {
		return Note{}, errors.New("Title cannot be empty")
	}
	if content == "" {
		return Note{}, errors.New("Content cannot be empty")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
