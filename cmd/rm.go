/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		RemoveTask(args)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func RemoveTask(args []string) {
	for _, taskId := range args {
		id, err := strconv.Atoi(taskId)
		if err != nil {
			fmt.Printf("Got error with parsing %v task\n", taskId)
			continue
		}
		_, err = Conn.Delete("toDo", "primary", []interface{}{id})
		if err != nil {
			fmt.Printf("Got error:%s with removing task %v\n", err, taskId)
		} else {
			fmt.Printf("Successfully removed task %v", taskId)
		}
	}
}
