package registry

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/z-spring/get/fetch"
	"github.com/z-spring/get/spinner"
)

const (
	Header   = "%s\t%s\t%s\n"
	NAME     = "NAME"
	PKG      = "PKG"
	IMPORTED = "IMPORTED"
	DATA     = "%s\t%s\t%s\n"
)

type SpecialPkg struct {
	ShortName string
	FullName  string
}

func NewSearchCommand() *cobra.Command {
	searchCmd := &cobra.Command{
		Use:   "search",
		Short: "you can search get-cli's support packages.",
		Args:  cobra.ExactArgs(1),
		Run:   runSearch,
	}
	return searchCmd
}

func runSearch(cmd *cobra.Command, args []string) {
	go spinner.Spinner(100 * time.Millisecond)

	keyWord := args[0]
	pkgs := fetch.GetAllPkgInfos(keyWord)
	if len(pkgs) == 0 {
		fmt.Printf("\rcan't find [%s] package", keyWord)
		return
	}
	sl, fl := GetMaxLen(pkgs)
	// input header
	// todo: 这里可以优化 用 tabwriter
	n := fmt.Sprintf("\r%-"+strconv.Itoa(sl)+"s", NAME)
	p := fmt.Sprintf("%-"+strconv.Itoa(fl)+"s", PKG)
	i := IMPORTED
	fmt.Printf(Header, n, p, i)

	HandleSpecialPkg(sl, fl, keyWord, pkgs)
	// input table data
	for _, pkg := range pkgs {
		s := fmt.Sprintf("%-"+strconv.Itoa(sl)+"s", pkg.ShortName)
		f := fmt.Sprintf("%-"+strconv.Itoa(fl)+"s", pkg.FullName)
		i := pkg.Imported
		fmt.Printf(DATA, s, f, i)
	}
}

func GetMaxLen(pkgs []fetch.Pkg) (ShortNameMaxLen int, FullNameMaxLen int) {
	i := pkgs[0]
	ShortNameMaxLen = len(i.ShortName)
	FullNameMaxLen = len(i.FullName)
	for _, v := range pkgs {
		if ShortNameMaxLen < len(v.ShortName) {
			ShortNameMaxLen = len(v.ShortName)
		}
		if FullNameMaxLen < len(v.FullName) {
			FullNameMaxLen = len(v.FullName)
		}
	}

	return
}

func HandleSpecialPkg(sl, fl int, keyWord string, pkgs []fetch.Pkg) {
	s := fmt.Sprintf("%-"+strconv.Itoa(sl)+"s", keyWord)
	i := "[OFFICIAL]"

	switch keyWord {
	case "beego":
		pkg := "github.com/beego/beego/v2@latest"
		f := fmt.Sprintf("%-"+strconv.Itoa(fl)+"s", pkg)
		fmt.Printf(DATA, s, f, i)
	case "get":
		pkg := "github.com/z-spring/get@latest"
		f := fmt.Sprintf("%-"+strconv.Itoa(fl)+"s", pkg)
		fmt.Printf(DATA, s, f, i)
	default:
		fmt.Printf("\r")
	}
}
