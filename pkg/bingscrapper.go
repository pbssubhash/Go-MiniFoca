package GoMiniFoca

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

func DownloadDocument(url string, dest string, ext string, wg *sync.WaitGroup) {
	splitter := strings.Split(url, "/")
	out, err := os.Create(dest + `\` + splitter[len(splitter)-1] + "." + ext)
	if err != nil {
		log.Fatalf("Error")
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error")
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalf("Error")
	}
	wg.Done()
}

func Scrap(website string, extension string, pages int, dest string) (map[string]string, error) {
	var wg sync.WaitGroup
	c := colly.NewCollector()
	result := make(map[string]string, 0)
	// website := os.Args[1]
	// extension := os.Args[2]
	// pages, _ := strconv.Atoi(os.Args[3])
	// dest := os.Args[4]
	c.OnHTML("#b_results li:nth-child(n)  h2  a", func(e *colly.HTMLElement) {
		go DownloadDocument(e.Attr("href"), dest, extension, &wg)
		splitter := strings.Split(e.Attr("href"), "/")
		result[dest+`\`+splitter[len(splitter)-1]+"."+extension] = e.Attr("href")
		wg.Add(1)
	})
	count := 0
	for {
		c.Visit(`https://www.bing.com/search?q=site:` + website + `%20filetype:` + extension + `&first=` + strconv.Itoa((count*10)+1))
		if count >= pages {
			break
		}
		count = count + 1
	}
	wg.Wait()
	return result, nil
}
