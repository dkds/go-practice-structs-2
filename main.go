package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"dkds.me/structs-2/note"
	"dkds.me/structs-2/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

func main() {
	title, content := getNoteData()

	todoText := getUserInput("Todo content: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		fmt.Println(err)
		return
	}

	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(note)
	if err != nil {
		fmt.Println(err)
	}

	printAnything(1)
	printAnything(1.0)
	printAnything("1.0")
	printAnything(note)
	printAnything(todo)
}

func printAnything(data any) {
	switch data.(type) {
	case int:
		fmt.Println("Integer:", data)
	case float64:
		fmt.Println("Float:", data)
	case string:
		fmt.Println("String:", data)
	case displayer:
		fmt.Println("Displayer:", data)
	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Failed to save the data to file")
		return err
	}
	fmt.Println("Successfully saved data")
	return nil
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
