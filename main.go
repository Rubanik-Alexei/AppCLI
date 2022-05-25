/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"log"

	"github.com/Rubanik-Alexei/AppCLI/cmd"
	"github.com/tarantool/go-tarantool"
)

var (
	opts tarantool.Opts = tarantool.Opts{
		User: "guest",
	}
)

var conn *tarantool.Connection

func main() {
	var err error
	conn, err = tarantool.Connect("127.0.0.1:3301", opts)
	if err != nil {
		log.Fatalf("Connection refused")
	}
	defer conn.Close()
	spaces := conn.Schema.Spaces
	if _, ok := spaces["toDo"]; !ok {
		_, err = conn.Call("box.schema.space.create", []interface{}{"toDo", map[string]bool{"if_not_exists": true}})
		if err != nil {
			errtxt := err.Error()
			if errtxt != "unsupported Lua type 'function' (0x20)" {
				log.Fatal(err)
			}
		}
		_, err = conn.Call("box.space.toDo:format", [][]map[string]string{
			{
				{"name": "name", "type": "string"},
				{"name": "description", "type": "string"},
				{"name": "completedAt", "type": "unsigned"}, //storing as UNIX time
			}})
		if err != nil {
			log.Fatal(err)
		}
		_, err = conn.Call("box.space.toDo:create_index", []interface{}{
			"todoIndex",
			map[string]interface{}{
				"parts":         []string{"name"},
				"type":          "TREE",
				"unique":        true,
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
	}
	cmd.Conn = conn
	//cmd.AddTask([]string{"cliApp", "FinishCLIApp"})
	cmd.Execute()
}
