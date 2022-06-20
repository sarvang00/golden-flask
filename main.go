package main

import (
	gvars "GoldenFlask/gvars"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Scraping for AudioBooks in given Urls")
	start := time.Now()
	// fmt.Println(gvars.AudioBooksDatabase)
	for _, value := range gvars.AudioBooksDatabase {
		value.DownloadAudiobook()
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("\nTotal time: %.2fs\n", secs)
}
