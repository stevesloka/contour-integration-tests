package helper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const externalURL = "http://demo.projectcontour.io"

func GetRequest(url string) (*http.Response, string) {
	// Make a get request
	rs, err := http.Get(url)
	// Process response
	if err != nil {
		panic(err) // More idiomatic way would be to print the error and die unless it's a serious error
	}
	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}

	bodyString := string(bodyBytes)

	return rs, bodyString
}

func GetUrl(paths ...string) (*url.URL, error) {
	return url.Parse(fmt.Sprintf("%s/%s", externalURL, strings.Join(paths, "/")))
}
