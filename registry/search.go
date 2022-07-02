package registry

import (
	"fmt"
	"get/fetch"
	"github.com/spf13/cobra"
	"strconv"
	"time"
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
		Args:  cobra.MinimumNArgs(1),
		Run:   runSearch,
	}
	return searchCmd
}

func runSearch(cmd *cobra.Command, args []string) {
	go spinner(100 * time.Millisecond)

	keyWord := args[0]
	pkgs := fetch.GetAllPkgInfos(keyWord)
	if len(pkgs) == 0 {
		fmt.Printf("\rcan't find [%s] package", keyWord)
		return
	}
	sl, fl := GetMaxLen(pkgs)

	n := fmt.Sprintf("\r%-"+strconv.Itoa(sl)+"s", NAME)
	p := fmt.Sprintf("%-"+strconv.Itoa(fl)+"s", PKG)
	i := IMPORTED
	fmt.Printf(Header, n, p, i)

	/*if keyWord == "beego" {
		s := fmt.Sprintf("%-"+strconv.Itoa(sl)+"s", "beego")
		f := fmt.Sprintf("%-"+strconv.Itoa(fl)+"s", "github.com/beego/beego/v2@latest")
		i := "[OFFICIAL]"
		fmt.Printf(DATA, s, f, i)
	}*/
	HandleSpecialPkg(sl, fl, keyWord, pkgs)

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

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func HandleSpecialPkg(sl, fl int, keyWord string, pkgs []fetch.Pkg) {
	s := fmt.Sprintf("%-"+strconv.Itoa(sl)+"s", keyWord)
	i := "[OFFICIAL]"

	switch keyWord {
	case "beego":
		pkg := "github.com/beego/beego/v2@latest"
		f := fmt.Sprintf("%-"+strconv.Itoa(fl)+"s", pkg)
		fmt.Printf(DATA, s, f, i)
	default:
		fmt.Printf("\r")
	}
}
