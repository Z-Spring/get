package fetch

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strings"
)

type Pkg struct {
	ShortName string
	FullName  string
	Imported  string
	Synopsis  string
}

func GetRespContent(keyWord string) string {
	url := fmt.Sprintf("https://pkg.go.dev/search?q=%s", keyWord)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36")

	if err != nil {
		log.Println(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	return string(bytes)
}

func GetPkg(keyWord string) Pkg {
	content := GetRespContent(keyWord)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		// define errors
		log.Println(err)
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
