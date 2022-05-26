/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tarantool/go-tarantool"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ListTasks()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ListTasks() {
	var tasks []Task
	err := Conn.SelectTyped("toDo", "toDoIndex", 0, 10, tarantool.IterEq, []interface{}{0}, &tasks)
	if err != nil {
		log.Fatal(err)
	}
	if len(tasks) == 0 {
		fmt.Println("Congrats, there's no incompleted tasks!")
		return
	}
	fmt.Println("List of incompleted tasks:")
	for _, task := range tasks {
		fmt.Printf("%v. %s : %s\n", task.Id, task.Name, task.Description)
	}
}
