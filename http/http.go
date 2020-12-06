package http

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// IsIPv6 checks whether the input is a valid IPv6 address.
func IsIPv6(ip string) bool {
	ips := strings.Split(ip, ":")
	l := len(ips)
	if l != 8 {
		return false
	}
	validStr := "0123456789ABCDEFabcdef"
	for _, v := range ips {
		if len(v) < 1 || len(v) > 4 {
			return false
		}
		for _, x := range v {
			if !strings.Contains(validStr, string(x)) {
				return false
			}
		}
	}

	return true
}

// IsIPv4 checks whether the string is a valid IPv4 address.
func IsIPv4(ip string) bool {
	ips := strings.Split(ip, ".")
	l := len(ips)
	if l != 4 {
		return false
	}
	for _, v := range ips {
		lv, tmp := len(v), 0
		if lv > 3 || lv < 1 {
			return false
		}
		if v[0] == '0' && lv > 1 {
			return false
		}
		for _, x := range v {
			if x < '0' || x > '9' {
				return false
			}
			tmp = 10*tmp + int(x-'0')
		}
		if tmp > 255 {
			return false
		}
	}

	return true
}

func getDNS1123Reg() *regexp.Regexp {
	onceDNS1123.Do(func() {
		format := "^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$"
		reg = regexp.MustCompile(format)
	})

	return reg
}

// IsDNS1123 checks if it is a valid DNS-1123 subdomain. It must consist of lower case alphanumeric characters,
// '-' or '.', and must start and end with an alphanumeric character.
func IsDNS1123(data string) bool {
	return getDNS1123Reg().MatchString(data)
}

// IsEmail validates input is a valid email address or not.
func IsEmail(email string) bool {
	email = strings.TrimSpace(email)
	errMsgRegeXp := "Error in running regular expression"
	urlRegeXp := "^([a-zA-Z0-9_\\-\\.]+)@((\\[[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3\\}\\.)|(([a-zA-Z0-9\\-]+\\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\\]?)$"
	match, err := regexp.MatchString(urlRegeXp, email)
	if err != nil {
		panic(errMsgRegeXp)
	}
	return match
}

// IsURL validates input is a valid url address or not.
func IsURL(url string) bool {
	urlRegeXp := "^((http|https|ftp)\\://)?([a-zA-Z0-9\\.\\-]+(\\:[a-zA-Z0-9\\.&amp;\\$\\-]+)*@)?((25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9])\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[0-9])|([a-zA-Z0-9\\-]+\\.)*[a-zA-Z0-9\\-]+\\.[a-zA-Z]{2,4})(\\:[0-9]+)?(/[^/][a-zA-Z0-9\\.\\,\\?\\'\\/\\+&amp;\\$#\\=~_\\-@]*)*$"
	match, err := regexp.MatchString(urlRegeXp, url)
	if err != nil {
		panic(err)
	}
	return match
}

// IsRequestURL checks if the string rawurl, assuming it was received in an HTTP request, is a valid
// URL confirm to RFC 3986.
func IsRequestURL(rawurl string) (bool, error) {
	url, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return false, err
	}
	if len(url.Scheme) == 0 {
		return false, nil
	}

	return true, nil
}

// IsRequestURI checks if the string rawurl, assuming it was received in an HTTP request, is an
// absolute URI or an absolute path.
func IsRequestURI(rawurl string) bool {
	_, err := url.ParseRequestURI(rawurl)
	return err == nil
}

func RespToStr(resp *http.Response) (string, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errMsg := "Can't Resolve Http Response. ioutil error: " + err.Error()
		fmt.Errorf(errMsg)
		return "", errors.New(errMsg)
	} else {
		return string(body), nil
	}
}

// HttpGet performs HTTP GET with given URL, user name and password.
func HttpGet(url, username, password string) (*http.Response, error) {
	return Http(http.MethodGet, url, username, password, nil)
}

// HttpPods performs HTTP POST with given URL, user name, password and data.
func HttpPost(url, username, password string, data []byte) (*http.Response, error) {
	return Http(http.MethodPost, url, username, password, data)
}

// HttpDelete performs HTTP DELTE with given URL, user name, password and data.
func HttpDelete(url, username, password string, data []byte) (*http.Response, error) {
	return Http(http.MethodDelete, url, username, password, data)
}

func HttpPut(url, username, password string, data []byte) (*http.Response, error) {
	return Http(http.MethodPut, url, username, password, data)
}

func IsValidHttpMethod(method string) bool {
	return strings.EqualFold(method, http.MethodGet) || strings.EqualFold(method, http.MethodPost) || strings.EqualFold(method, http.MethodPut) || strings.EqualFold(method, http.MethodDelete)
}

func Http(method, url, username, password string, data []byte) (*http.Response, error) {
	if !IsValidHttpMethod(method) {
		errMsg := "fail to send http request, parameter <method> is illegal."
		fmt.Errorf(errMsg)
		return nil, errors.New(errMsg)
	}

	if url == "" {
		errMsg := "fail to send http request, parameter <url> is empty."
		fmt.Errorf(errMsg)
		return nil, errors.New(errMsg)
	}

	var b io.Reader = nil

	if data != nil {
		b = bytes.NewBuffer(data)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, b)
	if err != nil {
		errMsg := "fail to send http request, " + err.Error()
		fmt.Errorf(errMsg)
		return nil, errors.New(errMsg)
	}
	req.Header.Add("Content-Type", "application/json")
	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := "fail to send http request, " + err.Error()
		fmt.Errorf(errMsg)
		return nil, errors.New(errMsg)
	}
	return resp, nil
}
