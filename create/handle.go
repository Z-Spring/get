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

	/* if len(args) == 1 {
		fmt.Printf("please input more args!\n\n")
		return &cobra.Command{}
	} */
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

	cmds := multiNames(args)

	return cmds

}

func multiNames(args []string) []*cobra.Command {
	var (
		cmds []*cobra.Command
		wg   sync.WaitGroup
	)

	// max commands: 10
	cmdChan := make(chan *cobra.Command, 10)

	// get multi pkgnames
	for _, argName := range args[1:] {
		wg.Add(1)
		go func(name string) {
			if name == "beego" {
				pkgName = "github.com/beego/beego/v2@latest"
			}
			// don't get infos from redis
			pkgName = GetPkgName2(name)
			cmdChan <- NewCommand(pkgName)
			// log.Printf("%s has downloaded!", pkgName)
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

	return cmds
}

// convert package names to command
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
// Discarded.
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
		// myredis.AddNameToRedis(name, pkgName)
		// myredis.AddNameToRedis2(name)
		return pkgName
	} else {
		pkgName, err = myredis.GetPkg(name)
		if err != nil {
			fmt.Println(err)
		}
		return pkgName
	}

}

// GetPkgName get full names from web
func GetPkgName2(name string) string {
	pkg := fetch.GetFirstPkgInfo(name)
	if pkg == (fetch.Pkg{}) {
		// 5秒后超时
		_, cancle := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancle()
		fmt.Printf("Timeout! Can't find [%s] package!\n", name)
		return ""
	}
	pkgName = pkg.FullName
	return pkgName

}
