package audiobook

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

type AudioBook struct {
	BookName  string
	Author    string
	Reader    string
	AudioUrl  string
	StorePath string //set the value after download -> BookName-Author/Reader/
}

// Function to download an audiobook
func (ab *AudioBook) DownloadAudiobook() {
	abStart := time.Now()
	audioBooks := []string{}
	pagesUrls := []string{}
	downloadPath := "./DownloadedContent/"
	if ab.Reader != "" {
		downloadPath += ab.BookName + "-" + ab.Author + "/" + ab.Reader
	} else {
		downloadPath += ab.BookName + "-" + ab.Author
	}
	downloadPath = strings.Trim(downloadPath, `'`)

	// Step-1: Loop through paginator; end in case of error
	for i := 1; ; i++ {
		var resp *http.Response
		var err error
		var url string
		if i == 1 {
			url = ab.AudioUrl
		} else {
			url = fmt.Sprintf("%s%d/", ab.AudioUrl, i)
		}
		resp, err = http.Get(url)

		if (resp.Request.URL.String() != url) || (err != nil || resp.StatusCode != 200) {
			// fmt.Println("Page finding broke")
			break
		} else {
			pagesUrls = append(pagesUrls, url)
		}
		// fmt.Println(resp.StatusCode)
		// fmt.Println("found ", url)
		// fmt.Println("is ", resp.Request.URL.String())
	}

	// Step-2: Find and loop through to find mp3 urls; add them to an array of strings
	for i := 0; i < len(pagesUrls); i++ {
		// Find mp3 files by regex on each page and add them to audioBooks
		audioBooks = append(audioBooks, GetMp3UrlsFromPage(pagesUrls[i])...)
	}

	// fmt.Println(audioBooks)

	// Step-3: Download mp3 files at a location (BookName-Author/Reader); update StorePath with location
	DownloadAudios(audioBooks, downloadPath)
	ab.StorePath = downloadPath

	abSecs := time.Since(abStart).Seconds()
	fmt.Printf("\nAudiobook time: %.2fs\n", abSecs)
}

// Function to get Urls of MP3 files from the page
func GetMp3UrlsFromPage(url string) []string {
	urlFetchStart := time.Now()
	bookUrls := []string{}
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)
	contentString := string(content)

	// regex for mp3 files on each page:
	re := regexp.MustCompile(`href=".*.mp3"`)
	for _, urlString := range re.FindAllString(contentString, -1) {
		urlString = urlString[5:]
		urlString = strings.Trim(urlString, `"`)
		bookUrls = append(bookUrls, urlString)
	}

	urlSecs := time.Since(urlFetchStart).Seconds()
	fmt.Printf("\nURL fetch time: %.2fs\n", urlSecs)

	return bookUrls
}

// Function to download audioclips to a specified folder
func DownloadAudios(urls []string, downloadLocation string) {
	downStartTime := time.Now()
	var wg sync.WaitGroup
	wg.Add(len(urls))

	err := os.MkdirAll(downloadLocation, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	for iter, url := range urls {
		go func(url string, iter int) {
			defer wg.Done()

			fileName := fmt.Sprintf("%s/chapter%d.mp3", downloadLocation, iter+1)
			fmt.Println("Downloading", url, "to", fileName)

			output, err := os.Create(fileName)
			if err != nil {
				log.Fatal("Error while creating", fileName, "-", err)
			}
			defer output.Close()

			res, err := http.Get(url)
			if err != nil {
				log.Fatal("http get error: ", err)
			} else {
				defer res.Body.Close()
				_, err = io.Copy(output, res.Body)
				if err != nil {
					log.Fatal("Error while downloading", url, "-", err)
				} else {
					fmt.Println("Downloaded", fileName)
				}
			}
		}(url, iter)
	}
	wg.Wait()
	downSecs := time.Since(downStartTime).Seconds()
	fmt.Printf("\nURL fetch time: %.2fs\n", downSecs)
}
