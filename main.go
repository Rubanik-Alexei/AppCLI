/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"log"

	"github.com/Rubanik-Alexei/AppCLI/cmd"
	"github.com/tarantool/go-tarantool"
)

func main() {
	conn, err := tarantool.Connect("127.0.0.1:3301", tarantool.Opts{
		User: "admin",
		Pass: "pass",
	})

	if err != nil {
		log.Fatalf("Connection refused")
	}
	defer conn.Close()
	_, err = conn.Call("box.schema.space.create", []interface{}{
		"geo",
		map[string]bool{"if_not_exists": true}})
	if err != nil {
		log.Fatal(err)
	}
	cmd.Execute()
}
