package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Edit   string
	Delete int
	Toggle int
	Print  bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo with title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit an existing todo by index and specify a new title, index:title")
	flag.IntVar(&cf.Delete, "delete", -1, "Specify a todo index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo index to toggle")
	flag.BoolVar(&cf.Print, "print", false, "List all todos")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Exec(todos *Todos) {
	switch {
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Edit != "":
		editArgs := strings.SplitN(cf.Edit, ":", 2)
		if len(editArgs) != 2 {
			fmt.Println("Invalid format for editing todo, please use index:title format")
			os.Exit(1)
		}

		index, err := strconv.Atoi(editArgs[0])
		if err != nil {
			fmt.Println("Invalid index, integer number is required")
			os.Exit(1)
		}

		todos.Edit(index, editArgs[1])
	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle)
	case cf.Delete != -1:
		todos.Delete(cf.Delete)
	case cf.Print:
		todos.Print()
	default:
		fmt.Println("No such command")
		os.Exit(1)
	}
}
