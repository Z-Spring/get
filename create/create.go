package create

import (
	"context"
	"fmt"
	"get/fetch"
	"get/myredis"
	"get/utils"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

var (
	searchCmd = &cobra.Command{
		Use:   "search",
		Short: "you can search get-cli's support packages.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			pkgNames := myredis.GetNamesFromRedis()
			m := utils.ConvertSliceToMap(pkgNames)
			for _, arg := range args {
				// todo 这里也可以优化，从string里面取 【这里是从list里面取】
				if utils.IsContain2(arg, m) {
					pkgName, _ := myredis.GetPkg(arg)
					fmt.Printf("found this package: %s\n", pkgName)
					fmt.Printf("you can use [get %s] to get this package.", arg)
				} else {
					pkg := fetch.GetPkg(arg)
					if pkg == (fetch.Pkg{}) {
						fmt.Printf("can't find [%s] package!\n", arg)
					} else {
						myredis.AddNameToRedis(arg, pkg.FullName)
						myredis.AddNameToRedis2(arg)
						fmt.Printf("found this package: %s\n", pkg.FullName)
						fmt.Printf("you can use [get %s] to get this package.", arg)
					}
				}
			}
		},
	}
	pkgName string
)

func CmdCreate(args []string) *cobra.Command {
	if len(args) == 1 {
		fmt.Printf("please input more args!\n\n")
		return &cobra.Command{}
	}
	var name string
	name = args[1]
	if name == "search" {
		return searchCmd
	}
	if name == "-h" {
		return &cobra.Command{}
	}
	pkgName = GetPkgName(name)
	if pkgName == "" {
		return &cobra.Command{}
	}
	cmd := DefineCmd(name)
	return cmd

}

func DefineCmd(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use: name,
		Run: func(cmd *cobra.Command, args []string) {
			var cmddArg string
			// 有些包用get，有些包用install
			if name == "ioc-golang" || name == "wire" {
				cmddArg = "install"
			} else {
				cmddArg = "get"
			}
			cmdd := exec.Command("go", cmddArg, "-u", pkgName)

			cmdd.Stdout = os.Stdout
			cmdd.Stderr = os.Stderr

			if err := cmdd.Run(); err != nil {
				log.Fatal(err)
			}
			fmt.Println("all downloads are done.")
		},
	}
	return cmd
}

// GetPkgName input name find pkg
func GetPkgName(name string) string {
	names := myredis.GetNamesFromRedis()
	m := utils.ConvertSliceToMap(names)

	var err error
	if !utils.IsContain2(name, m) {
		pkg := fetch.GetPkg(name)
		if pkg == (fetch.Pkg{}) {
			_, cancle := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancle()
			fmt.Printf("can't find [%s] package!\n", name)
			return ""
		}
		pkgName = pkg.FullName
		// 查找到后添加到redis中
		// todo : 这里参数命名不规范 name pkgName分不清
		myredis.AddNameToRedis(name, pkgName)
		myredis.AddNameToRedis2(name)
		return pkgName
	} else {
		pkgName, err = myredis.GetPkg(name)
		if err != nil {
			fmt.Println(err)
		}
		return pkgName
	}

}

func Ask() {
	var y, n = "y", "n"
	var input string
	fmt.Printf("found this package: %s ,do you want to get it?(%s/%s)", pkgName, y, n)
	if _, err := fmt.Scanf("%s", &input); err != nil {
		fmt.Printf("fmt.Scanf failed with '%s'\n", err)
	}
	if input == "n" || input == "Y" {

	}
}
