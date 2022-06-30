/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	args2 "get/args"
	"get/create"
	"get/fetch"
	"get/myredis"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var (
	ginCmd = &cobra.Command{
		Use:   "gin",
		Short: "Gin is a web framework written in Go (Golang). ",
		Run: func(cmd *cobra.Command, args []string) {
			cmdd := exec.Command("go", "get", "-u", "github.com/gin-gonic/gin")
			CmdRun(cmdd)
		},
	}

	gjsonCmd = &cobra.Command{
		Use:   "gjson",
		Short: "GJSON is a Go package that provides a fast and simple way to get values from a json document. ",
		Run: func(cmd *cobra.Command, args []string) {
			cmdd := exec.Command("go", "get", "-u", "github.com/tidwall/gjson")
			CmdRun(cmdd)
		},
	}

	goqueryCmd = &cobra.Command{
		Use:   "goquery",
		Short: "goquery brings a syntax and a set of features similar to jQuery to the Go language.",
		Run: func(cmd *cobra.Command, args []string) {
			cmdd := exec.Command("go", "get", "-u", "github.com/PuerkitoBio/goquery")
			CmdRun(cmdd)
		},
	}

	beeGoCmd = &cobra.Command{
		Use: "beego",
		Long: `Beego is used for rapid development of enterprise application in Go,
including RESTful APIs, web apps and backend services.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmdd := exec.Command("go", "get", "github.com/beego/beego/v2@latest")
			CmdRun(cmdd)
		},
	}

	redisVersion string
	redisCmd     = &cobra.Command{
		Use:   "myredis",
		Short: "Package myredis implements a Redis client.",
		Long:  "about version,you can choose v9,v8,v7",
		Run: func(cmd *cobra.Command, args []string) {
			cmdd := exec.Command("go", "get", "-u", fmt.Sprintf("github.com/go-myredis/myredis/%s", redisVersion))
			CmdRun(cmdd)
		},
	}

	elasticsearchVersion string
	elasticsearchCmd     = &cobra.Command{
		Use:   "elastic",
		Short: "Package elasticsearch provides a Go client for Elasticsearch.",
		Run: func(cmd *cobra.Command, args []string) {
			cmdd := exec.Command("go", "get", "-u", fmt.Sprintf("github.com/elastic/go-elasticsearch/%s", elasticsearchVersion))
			CmdRun(cmdd)
		},
	}

	jwtCmd = &cobra.Command{
		Use:   "jwt",
		Short: "A go (or 'golang' for search engine friendliness) implementation of JSON Web Tokens.",
		Run: func(cmd *cobra.Command, args []string) {
			cmdd := exec.Command("go", "get", "-u", "github.com/golang-jwt/jwt/v4")
			CmdRun(cmdd)
		},
	}

	searchCmd = &cobra.Command{
		Use:   "search",
		Short: "you can search get-cli's support packages.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			pkgNames := myredis.GetPkgNameFromRedis()
			m := args2.ConvertSliceToMap(pkgNames)

			for _, arg := range args {
				if args2.IsContain2(arg, m) {
					// todo: 这里可以用redis来缓存一下
					pkg := fetch.GetPkg(arg)
					fmt.Printf("found this package: %s\n", pkg.FullName)
					fmt.Printf("you can use [get %s] to get this package.", arg)
				} else {
					fmt.Printf("\ncan't find [%s] package", arg)
				}

			}
		},
	}
	yourOwnCmd *cobra.Command

	pkgName string
	pkgRepo string
	addCmd  = &cobra.Command{
		Use:   "add",
		Short: "you can add go packages into get-cli.",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			pkgNames := myredis.GetPkgNameFromRedis()
			pkgRepos := myredis.GetPkgRepoFromRedis()

			n := args2.ConvertSliceToMap(pkgNames)
			r := args2.ConvertSliceToMap(pkgRepos)

			if args2.IsContain2(pkgName, n) || args2.IsContain2(pkgRepo, r) {
				fmt.Print("get-cli has added this package")
			} else {
				yourOwnCmd = create.CmdCreate(pkgName, pkgRepo)
				myredis.AddPkgToRedis(pkgName, pkgRepo)
				fmt.Println("your package added!")
			}

		},
	}

	otherCmd = &cobra.Command{
		Use: "other",
		Run: func(cmd *cobra.Command, args []string) {
			//cmd.SetFlagErrorFunc()
		},
	}
)

func CmdRun(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("all downloads are done.")
}

func init() {
	rootCmd.AddCommand(
		ginCmd,
		gjsonCmd,
		goqueryCmd,
		beeGoCmd,
		redisCmd,
		elasticsearchCmd,
		searchCmd,
		jwtCmd,
		otherCmd,
		addCmd,
		//yourOwnCmd,
	)
	addCmd.AddCommand(yourOwnCmd)

	redisCmd.Flags().StringVarP(&redisVersion, "myredis", "v", "v9", "choose go-myredis version")
	elasticsearchCmd.Flags().StringVarP(&elasticsearchVersion, "elasticsearch", "v", "v8", "choose go-elasticsearch version")

	addCmd.Flags().StringVarP(&pkgName, "pkgName", "n", "", "input your pkgName")
	addCmd.Flags().StringVarP(&pkgRepo, "pkgRepo", "r", "", "input your pkgRepo")
	addCmd.MarkFlagsRequiredTogether("pkgName", "pkgRepo")
	//searchCmd.Flags().String("package", "", "get search [pkg]")

	// ginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
