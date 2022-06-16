package main

import (
	gvars "GoldenFlask/gvars"
	"fmt"
)

func main() {
	fmt.Println("Scraping for AudioBooks in given Urls")
	fmt.Println(gvars.AudioBooksDatabase)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
