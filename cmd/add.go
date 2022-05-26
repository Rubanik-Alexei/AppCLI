/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Command for adding tasks",
	Long:  `You can add a task with it's description in to-do list`,
	Run: func(cmd *cobra.Command, args []string) {
		// get the flag value, its default value is false
		// fstatus, _ := cmd.Flags().GetBool("float")
		// if fstatus { // if status is true, call addFloat
		// 	addFloat(args)
		// } else {
		// 	addInt(args)
		// }
		AddTask(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	//addCmd.Flags().BoolP("float", "f", false, "Add Floating Numbers")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func AddTask(args []string) {
	task := Task{}
	if len(args) < 2 {
		fmt.Printf("Not enough arguments for 'add' command")
	}
	//arg.MustParse(&task)
	var tuple []interface{} = []interface{}{nil}
	for _, v := range args {
		tuple = append(tuple, v)
	}
	tuple = append(tuple, 0)
	err := Conn.InsertTyped("toDo", tuple, &task)
	if err != nil && err.Error() != "msgpack: invalid code 94 decoding int64" {
		fmt.Printf("Error in 'add' command: %s", err)
	} else {
		fmt.Printf("Successfully added the task: '%s': %s\n", tuple[1], tuple[2])
	}
}

// func addInt(args []string) {
// 	var result int
// 	for _, number := range args {
// 		if tmp, err := strconv.Atoi(number); err == nil {
// 			result += tmp
// 		}
// 	}
// 	fmt.Printf("Addition of numbers %s is %d", args, result)
// }

// func addFloat(args []string) {
// 	var result float64
// 	for _, v := range args {
// 		// convert string to float64
// 		ftemp, err := strconv.ParseFloat(v, 64)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		result = result + ftemp
// 	}
// 	fmt.Printf("Sum of floating numbers %s is %f", args, result)
// }
