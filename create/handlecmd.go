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
	pkgName = GetPkgName(name)
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

// GetPkgName get full names from web
func GetPkgName(name string) string {
	pkg := fetch.GetFirstPkgInfo(name)
	if pkg == (fetch.Pkg{}) {
		fmt.Printf("Can't find [%s] package!\n", name)
		os.Exit(1)
		return ""
	}
	pkgName = pkg.FullName
	return pkgName

}
