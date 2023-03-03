package create

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"github.com/z-spring/get/fetch"
	"github.com/z-spring/get/myredis"
	"github.com/z-spring/get/registry"
	"github.com/z-spring/get/utils"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

var (
	pkgName string
)

func HandleCommand(args []string) []*cobra.Command {
	//go registry.Spinner(100 * time.Millisecond)

	if len(args) == 1 {
		fmt.Printf("please input more args!\n\n")
		return &cobra.Command{}
	}
	name := args[1]

	if name == "search" {
		return []*cobra.Command{registry.NewSearchCommand()}
	}
	if name == "-h" {
		return []*cobra.Command{&cobra.Command{}}
	}
	if name == "weather" {
		return []*cobra.Command{registry.NewWeatherCommand()}
	}

	var cmds []*cobra.Command
	//var cmd *cobra.Command
	cmdChan := make(chan *cobra.Command, 10)
	var wg sync.WaitGroup
	// get multi pkgnames
	for _, argName := range args[1:] {
		wg.Add(1)
		go func(name string) {
			if name == "beego" {
				pkgName = "github.com/beego/beego/v2@latest"
			}
			pkgName = GetPkgName(name)

			// todo: 这里怎么实现？

			cmdChan <- NewCommand(pkgName)
			//cmd
			//log.Println(cmds)

			log.Printf("%s has downloaded!", pkgName)
			wg.Done()
		}(argName)
	}
	go func() {
		wg.Wait()
		close(cmdChan)
	}()

	for cmddd := range cmdChan {

		cmds = append(cmds, cmddd)
	}
	//defer close(cmdChan)

	/*	if pkgName == "" {
		return []*cobra.Command{&cobra.Command{}}
	}*/
	//cmd := NewCommand(name)
	return cmds

}

func NewCommand(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use: name,
		Run: func(cmd *cobra.Command, args []string) {
			var cmddArg string
			// 有些包用get，有些包用install
			// todo: 这里可以弄成slice
			if name == "ioc-golang" || name == "wire" || name == "get" {
				cmddArg = "install"
			} else {
				cmddArg = "get"
			}
			cmdd := exec.Command("go", cmddArg, "-u", pkgName)
			//cmdd := exec.Command("go", cmddArg, "-u", name)

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
	// find whether the pkg is in redis
	names := myredis.GetNamesFromRedis()
	m := utils.ConvertSliceToMap(names)

	var err error
	if !utils.IsContain2(name, m) {
		pkg := fetch.GetFirstPkgInfo(name)
		if pkg == (fetch.Pkg{}) {
			_, cancle := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancle()
			fmt.Printf("Timeout! Can't find [%s] package!\n", name)
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
