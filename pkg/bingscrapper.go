package GoMiniFoca

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gocolly/colly"
)

func DownloadDocument(url string, dest string, ext string, wg *sync.WaitGroup) {
	// splitter := strings.Split(url, "/")
	out, err := os.Create(dest + ext)
	if err != nil {
		log.Fatalf("Error-11")
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error-12")
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalf("Error-13")
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

		// splitter := strings.Split(e.Attr("href"), "/")
		algorithm := md5.New()
		algorithm.Write([]byte(e.Attr("href")))
		result[dest+string(os.PathSeparator)+string(hex.EncodeToString(algorithm.Sum(nil)))+".zip"] = e.Attr("href")
		go DownloadDocument(e.Attr("href"), dest, string(hex.EncodeToString(algorithm.Sum(nil))+".zip"), &wg)
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
