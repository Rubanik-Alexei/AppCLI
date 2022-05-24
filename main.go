/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"log"

	"github.com/Rubanik-Alexei/AppCLI/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
