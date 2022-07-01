/*
Copyright © 2022 murphy <murphyqq1@gmail.com>

*/
package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "get",
	Short: "you can get go packages easily",
	Long: `Have you ever had the problem that when you want to import a dependency package for your project, 
but you only remember its name, but not its full path? 
For example: you want to get the gin package,but you don't remember its full path (github.com/gin-gonic/gin). 
Now you can use this tool to get the package you want easier by input a simple command: [get gin]. 
It's that simple!`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func init() {
}
