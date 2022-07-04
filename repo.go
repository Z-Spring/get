/*
Copyright © 2022 murphy <murphyqq1@gmail.com>

*/
package main

import (
	"github.com/Z-Spring/get/create"
	"os"
)

var (
	// todo 这里要限制一下输入参数的数量  1个
	pkgName      = os.Args
	yourAddedCmd = create.HandleCommand(pkgName)
)

func init() {
	rootCmd.AddCommand(
		yourAddedCmd,
	)
}
