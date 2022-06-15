package main

import (
	gvars "GoldenFlask/gvars"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get(gvars.Site)
	checkErr(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	checkErr(err)
	content := string(body)

	// fmt.Println(content)
	parsedTagData := parseAnchorTags(content)
	for tagData := range parsedTagData {
		fmt.Println(parsedTagData[tagData])
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func parseAnchorTags(httpStringContent string) []string {
	regexPattern := `<a.*>.*</a>|<a.*/>`
	re := regexp.MustCompile(regexPattern)
	matchedTags := re.FindAllString(httpStringContent, -1)
	return matchedTags
}
