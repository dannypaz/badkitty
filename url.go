package badkitty

import (
	"fmt"
	"net/url"
	"strings"
)

func bruteForceURL(u string) string {
	u = strings.Replace(u, "http://", "", -1)
	u = strings.Replace(u, "https://", "", -1)
	offset := strings.Index(u, "?")
	if offset != -1 {
		u = u[0:offset]
	}
	return u
}

// NormalizeURL ...
func NormalizeURL(turl string) string {
	parsedURL, err := url.Parse(turl)
	if err != nil {
		return bruteForceURL(turl)
	}
	purl := strings.ToLower(fmt.Sprintf("%s%s", parsedURL.Host, parsedURL.EscapedPath()))
	return purl
}
