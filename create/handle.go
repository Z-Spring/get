package create

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/z-spring/get/fetch"
	// "github.com/z-spring/get/myredis"
	"github.com/z-spring/get/registry"
	// "github.com/z-spring/get/utils"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

var (
	pkgName string
)

func HandleCommand(args []string) *cobra.Command {
	if len(args) == 1 {
		return &cobra.Command{}
	}
	name := args[1]
	pkgName = GetPkgName2(args[2])

	switch name {
	case "search":
		return registry.NewSearchCommand()
	case "-h":
		return &cobra.Command{}
	case "":
		return &cobra.Command{}
	case "weather":
		return registry.NewWeatherCommand()
	default:
		cmd := NewCommand(name)
		return cmd
	}

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
/* func GetPkgName(name string) string {
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

} */

// GetPkgName2 GetPkgName get full names from web
func GetPkgName2(name string) string {
	pkg := fetch.GetFirstPkgInfo(name)
	if pkg == (fetch.Pkg{}) {
		fmt.Printf("Can't find [%s] package!\n", name)
		return ""
	}
	pkgName = pkg.FullName
	return pkgName

}
