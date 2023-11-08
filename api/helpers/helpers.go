package helpers

import (
	"os"
	"strings"
)

func EnforceHttp(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

func RemoveDomainErr(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}

	newUrl := strings.Replace(url, "http://", "", 1)
	newUrl = strings.Replace(newUrl, "https://", "", 1)
	newUrl = strings.Replace(newUrl, "www.", "", 1)
	newUrl = strings.Split(newUrl, "/")[0]

	if newUrl == os.Getenv("DOMAIN") {
		return false
	}

	return true
}
