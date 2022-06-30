/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	Execute()
	//myredis.AddPkgToRedis2()
	/*	fmt.Println(myredis.GetPkgRepoFromRedis())
		fmt.Println(myredis.GetPkgNameFromRedis())*/
	/*s := fetch.GetRespContent("goquery")
	fmt.Println(s)*/
	/*	pkg := fetch.GetPkg("beego")
		fmt.Println(pkg)*/
}
