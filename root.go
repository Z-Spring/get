/*
Copyright © 2022 murphy <murphyqq1@gmail.com>

*/
package main

import (
	"github.com/spf13/cobra"
	"log"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "get",
	Short: "you can get go packages easily",
	Long: `Have you ever had the problem that when you want to import a dependency package for your project, 
you only remember its name, but not its full path? 
For example: you want to import the gin package, 
but you don't remember the full path to import the package (github.com/gin-gonic/gin). 
Now you can use this tool to get the package you want more easily by typing a simple statement: [get gin]. 
It's that simple!`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.get.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().String("gin", "$ go get -u github.com/gin-gonic/gin", "Help message for toggle")
}
