package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string) {
	todo := Todo{Title: title, Completed: false, CompletedAt: nil, CreatedAt: time.Now()}

	*todos = append(*todos, todo)
}

func (todos *Todos) Delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = slices.Delete(t, index, index+1)

	return nil
}

func (todos *Todos) Toggle(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed
	if !isCompleted {
		completedAt := time.Now()
		t[index].CompletedAt = &completedAt
	} else {
		t[index].CompletedAt = nil
	}

	t[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) Edit(index int, title string) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title

	return nil
}

func (todos *Todos) Print() {
	todoTable := table.New(os.Stdout)

	todoTable.SetRowLines(true)
	todoTable.SetLineStyle(table.StyleBlue)
	todoTable.SetHeaderStyle(table.StyleBold)
	todoTable.SetHeaderVerticalAlignment(table.AlignCenter)
	todoTable.SetHeaders("#", "Title", "Completion status", "Created At", "Completed At")

	for i, todo := range *todos {
		completed := "❌"
		completedAt := "-"
		createdAt := todo.CreatedAt.Format(time.RFC1123)

		if todo.Completed {
			completed = "✅"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC1123)
			}
		}

		todoTable.AddRow(strconv.Itoa(i), todo.Title, completed, createdAt, completedAt)
	}

	todoTable.Render()
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		errText := "invalid index"
		fmt.Println(errText)
		return errors.New(errText)
	}

	return nil
}
