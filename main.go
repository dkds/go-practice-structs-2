package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"dkds.me/structs-2/note"
	"dkds.me/structs-2/todo"
)

func main() {
	title, content := getNoteData()

	todoText := getUserInput("Todo content: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()

	err = todo.Save()
	if err != nil {
		fmt.Println("Failed to save the todo to a file")
		return
	}
	fmt.Println("Successfully saved todo")

	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	note.Display()

	err = note.Save()
	if err != nil {
		fmt.Println("Failed to save the note to a file")
		return
	}
	fmt.Println("Successfully saved note")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
