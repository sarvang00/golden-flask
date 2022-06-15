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
	parsedTagData := getAnchorTags(content)
	for tagData := range parsedTagData {
		fmt.Println(parsedTagData[tagData])
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func getAnchorTags(httpStringContent string) []string {
	regexPattern := "<a.*>.*</a>"
	re := regexp.MustCompile(regexPattern)
	matchedTags := re.FindAllString(httpStringContent, -1)
	matchedTags = removeDuplicateStrings(matchedTags)
	return matchedTags
}

func removeDuplicateStrings(anchors []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range anchors {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
