package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (note Note) Display() {
	fmt.Println("Title: ", note.title)
	fmt.Println("Content: ", note.content)
}

func New(title string, content string) (Note, error) {
	if title == "" {
		return Note{}, errors.New("Title cannot be empty")
	}
	if content == "" {
		return Note{}, errors.New("Content cannot be empty")
	}
	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}
