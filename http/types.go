package http

import (
	"regexp"
	"sync"
)

const (
	SUCCESS_HTTP_CODE = 0
)

var (
	reg         *regexp.Regexp
	onceDNS1123 sync.Once
)
