/*
Copyright © 2022 murphy <murphyqq1@gmail.com>

*/
package main

import (
	"get/create"
	"os"
)

var (
	// todo 这里要限制一下输入参数的数量  1个
	pkgName      = os.Args
	yourAddedCmd = create.CmdCreate(pkgName)
)

func init() {
	rootCmd.AddCommand(
		yourAddedCmd,
	)
}
