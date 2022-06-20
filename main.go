package main

import (
	gvars "GoldenFlask/gvars"
	"fmt"
)

func main() {
	fmt.Println("Scraping for AudioBooks in given Urls")
	// fmt.Println(gvars.AudioBooksDatabase)
	for _, value := range gvars.AudioBooksDatabase {
		value.DownloadAudiobook()
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
