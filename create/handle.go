package create

import (
	"context"
	"fmt"
	"github.com/murphyzz/get/fetch"
	"github.com/murphyzz/get/myredis"
	"github.com/murphyzz/get/registry"
	"github.com/murphyzz/get/utils"
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
	pkgName string
)

func HandleCommand(args []string) *cobra.Command {
	//go registry.Spinner(100 * time.Millisecond)

	if len(args) == 1 {
		fmt.Printf("please input more args!\n\n")
		return &cobra.Command{}
	}
	name := args[1]

	if name == "search" {
		return registry.NewSearchCommand()
	}
	if name == "-h" {
		return &cobra.Command{}
	}
	if name == "weather" {
		return registry.NewWeatherCommand()
	}
	//HandleSpecialCommand(name)

	pkgName = GetPkgName(name)
	if pkgName == "" {
		return &cobra.Command{}
	}
	if name == "beego" {
		pkgName = "github.com/beego/beego/v2@latest"
	}
	cmd := NewCommand(name)
	return cmd

}

func NewCommand(name string) *cobra.Command {
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

// GetPkgName input name to find pkgName
func GetPkgName(name string) string {
	names := myredis.GetNamesFromRedis()
	m := utils.ConvertSliceToMap(names)

	var err error
	if !utils.IsContain2(name, m) {
		pkg := fetch.GetFirstPkgInfo(name)
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
