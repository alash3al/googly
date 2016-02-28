package main

import "fmt"
import "flag"
import "strings"
import "net/url"
import "github.com/PuerkitoBio/goquery"

// -------------------------

var (
	resultFormat  	string 		= 
`
%d)- %s
%s<%s>
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
`

	Welcome 		string 		=
`
Googly v1.0, by Mohammed Al Ashaal, <http://alash3al.xyz> 
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
`
)

// -------------------------

var (
	Query 		*string		=	flag.String("query", "", "the keywords you want to search about, (default empty)")
	Num 		*int64		=	flag.Int64("num", 10, "the number of results to show per page")
	Offset		*int64 		=	flag.Int64("offset", 0, "continuing displaying results from this offset, (default 0)")
	Personal 	*int		=	flag.Int("personal", 0, "whether to display personal results or not, (default 0)")
	Site		*string 	=	flag.String("site", "", "only display results from this website, (default empty)")
	Similar		*string 	=	flag.String("similar", "", "get results from sites that may be similar to the provided site, (default empty)")
	Extension	*string 	=	flag.String("ext", "", "get results of files that end with the provided extensions, (default empty)")
	Summary		*int 		=	flag.Int("summary", 1, "wether to show summary or not")
)

// -------------------------

func init() {
	flag.Parse()
	fmt.Println(Welcome)
}

// -------------------------

func main() {
	res, err := goquery.NewDocument(fmt.Sprintf(
		"https://www.google.com/search?oe=UTF-8&hl=EN&q=%s&num=%d&start=%d&pws=%d&as_sitesearch=%s&as_rq=%s&as_filetype=%s",
		strings.Replace(*Query, " ", "+", -1),
		*Num,
		*Offset,
		*Personal,
		*Site,
		*Similar,
		*Extension,
	))
	if err != nil {
		fmt.Println(err)
		return
	}
	res.Find("#res .g").Each(func(i int, item *goquery.Selection) {
		title := item.Find(".r").Text()
		summary := item.Find(".st").Text() + "\n"
		link := (func() string {
			href, _ := item.Find("a").Attr("href")
			u, _ := url.Parse(href)
			return u.Query().Get("q")
		})()
		if *Summary < 1 {
			summary = ""
		}
		// fmt.Printf(resultFormat, i + 1, ansi.Color(title, "green+b:white+h"), summary, ansi.Color(link, "red+B:white+h"))
		fmt.Printf(resultFormat, i + 1, title, summary, link)
	})
}
