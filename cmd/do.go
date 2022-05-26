/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Complete specified task",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		DoTask(args)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func DoTask(args []string) {
	fmt.Println("You've completed following tasks:")
	var tmpTime int64
	for _, v := range args {
		tmpTime = time.Now().Unix()
		id, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("Got error:%s for task %s\n", err, v)
		}
		resp, err := Conn.Update("toDo", "primary", []interface{}{id}, []interface{}{[]interface{}{"=", 3, tmpTime}})
		if err != nil {
			fmt.Printf("Got error:%s for task %s\n", resp.Error, v)
		} else {
			fmt.Printf("%s at %s\n", strings.Split(resp.String(), " ")[3], time.Unix(tmpTime, 0))
		}
	}
}
