/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Rubanik-Alexei/AppCLI/cmd"
	"github.com/Rubanik-Alexei/AppCLI/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tarantool/go-tarantool"
)

var (
	opts tarantool.Opts = tarantool.Opts{
		User: "guest",
	}
)

var conn *tarantool.Connection

func main() {
	var isEmpty bool
	var err error
	conn, err = tarantool.Connect("127.0.0.1:3301", opts)
	if err != nil {
		log.Fatalf("Connection refused")
	}
	defer conn.Close()
	spaces := conn.Schema.Spaces
	if _, ok := spaces["toDo"]; !ok {
		_, err = conn.Call("box.schema.sequence.create", []interface{}{"S", map[string]bool{"if_not_exists": true}})
		if err != nil {
			log.Fatal(err)
		}
		_, err = conn.Call("box.schema.space.create", []interface{}{"toDo", map[string]bool{"if_not_exists": true}})
		if err != nil {
			errtxt := err.Error()
			if errtxt != "unsupported Lua type 'function' (0x20)" {
				log.Fatal(err)
			}
		}
		_, err = conn.Call("box.space.toDo:format", [][]map[string]string{
			{
				{"name": "id", "type": "integer"},
				{"name": "name", "type": "string"},
				{"name": "description", "type": "string"},
				{"name": "completedAt", "type": "integer"}, //storing as UNIX time
			}})
		if err != nil {
			log.Fatal(err)
		}
		_, err = conn.Call("box.space.toDo:create_index", []interface{}{
			"primary",
			map[string]interface{}{
				"parts":         []string{"id"},
				"type":          "TREE",
				"unique":        true,
				"sequence":      "S",
				"if_not_exists": true}})
		if err != nil {
			log.Fatal(err)
		}
		_, err = conn.Call("box.space.toDo:create_index", []interface{}{
			"toDoIndex",
			map[string]interface{}{
				"parts":         []string{"completedAt"},
				"type":          "TREE",
				"unique":        false,
				"if_not_exists": true}})
		if err != nil {
			log.Fatal(err)
		}
		conn.Close()
		conn, err = tarantool.Connect("127.0.0.1:3301", opts)
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		isEmpty = true
	}
	cmd.Conn = conn
	if isEmpty {
		cmd.AddTask([]string{"cliApp", "FinishCLIApp"})
		cmd.AddTask([]string{"JobsSearching", "Check new jobs"})
	}
	//cmd.DoTask([]string{"4"})
	// cmd.ListTasks()
	// cmd.DoTask([]string{"17"})
	// cmd.ListTasks()
	//cmd.RemoveTask([]string{"1"})
	//cmd..DoTask([]string{"14"})
	cmd.Execute()
	InitModel()

}

func InitModel() {
	cmd.P = tea.NewProgram(tui.NewModel(cmd.ListTasksTUI()))
	if err := cmd.P.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
