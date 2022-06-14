package main

import (
	gvars "GoldenFlask/gvars"
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get(gvars.Site)
	checkErr(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	checkErr(err)
	content := string(body)

	fmt.Println(content)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
