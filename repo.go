/*
Copyright © 2022 murphy <murphyqq1@gmail.com>
*/
package main

import (
	"os"

	"github.com/z-spring/get/create"
)

var (
	pkgName      = os.Args
	yourAddedCmd = create.HandleCommand(pkgName)
)

func init() {
	rootCmd.AddCommand(yourAddedCmd)

}
