package fetch

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Pkg struct {
	ShortName string
	FullName  string
	Imported  string
	Synopsis  string
}

func getPkgContent(keyWord string) string {
	url := fmt.Sprintf("https://pkg.go.dev/search?q=%s", keyWord)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36")

	if err != nil {
		return "can not send request, please try again later"
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "can not send request, please try again later"
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "fetch error, please try again later"
		// log.Println(err)
	}
	defer resp.Body.Close()
	return string(bytes)
}

func GetFirstPkgInfo(keyWord string) Pkg {
	content := getPkgContent(keyWord)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		// define errors
		fmt.Println(err)
	}

	sum := doc.Find("div.SearchResults-summary > h1 > strong").Text()
	modules, _ := strconv.Atoi(sum)
	if modules == 0 {
		return Pkg{}
	}
	var pkg Pkg

	selection := doc.Find("div.SearchSnippet").First()

	shortName := selection.Find("div.SearchSnippet-headerContainer > h2 > a").Text()
	shortName = strings.Split(shortName, "\n")[1]
	shortName = strings.ReplaceAll(shortName, "\n", "")
	shortName = strings.TrimSpace(shortName)

	fullName := selection.Find("div.SearchSnippet-headerContainer > h2 > a > span").Text()
	fullName = strings.Trim(fullName, "()")

	synopsis := selection.Find("p").Text()
	synopsis = strings.TrimSpace(synopsis)

	imported := selection.Find("div.SearchSnippet-infoLabel >a > strong").Text()

	pkg = Pkg{
		ShortName: shortName,
		FullName:  fullName,
		Imported:  imported,
		Synopsis:  synopsis,
	}

	return pkg

}

func GetAllPkgInfos(keyWord string) []Pkg {
	content := getPkgContent(keyWord)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		err = errors.New("can't parse the fetch content")
		fmt.Println(err.Error())
	}

	sum := doc.Find("div.SearchResults-summary > h1 > strong").Text()
	modules, _ := strconv.Atoi(sum)
	if modules == 0 {
		return nil
	}
	var pkgs []Pkg

	doc.Find("div.SearchSnippet").Each(func(i int, selection *goquery.Selection) {
		shortName := selection.Find("div.SearchSnippet-headerContainer > h2 > a").Text()
		shortName = strings.Split(shortName, "\n")[1]
		shortName = strings.ReplaceAll(shortName, "\n", "")
		shortName = strings.TrimSpace(shortName)

		fullName := selection.Find("div.SearchSnippet-headerContainer > h2 > a > span").Text()
		fullName = strings.Trim(fullName, "()")

		synopsis := selection.Find("p").Text()
		synopsis = strings.TrimSpace(synopsis)

		imported := selection.Find("div.SearchSnippet-infoLabel >a > strong").Text()

		pkg := Pkg{
			ShortName: shortName,
			FullName:  fullName,
			Imported:  imported,
			Synopsis:  synopsis,
		}
		pkgs = append(pkgs, pkg)

	})

	return pkgs

}
