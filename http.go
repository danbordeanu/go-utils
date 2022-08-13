package utils

import (
	"fmt"
	"net/http"
	"strings"
)

// GetHttpRequestScheme determines the protocol scheme used in the request call. Returns either "http" or "https", always lowercase
func GetHttpRequestScheme(r *http.Request) string {
	var scheme string
	switch {
	case r.URL.Scheme == "https":
		scheme = "https"
	case r.TLS != nil:
		scheme = "https"
	case strings.HasPrefix(strings.ToLower(r.Proto), "https"):
		scheme = "https"
	case strings.ToLower(r.Header.Get("X-Forwarded-Proto")) == "https":
		scheme = "https"
	default:
		scheme = "http"
	}
	return scheme
}

// GetHttpRequestBaseUrl returns a canonical hostname for the current request in the form of <scheme>://<host or ip>:<port>
func GetHttpRequestBaseUrl(r *http.Request) string {
	return GetHttpRequestScheme(r) + "://" + r.Host
}

// ComputeFullPath will add the current canonical hostname to a URL, if it doesn't already contain the canonical hostname (i.e. it is only a path)
func ComputeFullPath(url string, r *http.Request) string { // TODO: Find a better name for this function after smoking a couple of joints
	//goland:noinspection HttpUrlsUsage
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = fmt.Sprintf("%s/%s", GetHttpRequestBaseUrl(r), strings.TrimPrefix(url, "/"))
	}
	return url
}
