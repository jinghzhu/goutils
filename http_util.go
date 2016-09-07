package GoUtils

import (
    "bytes"
    "errors"
    "io"
    "io/ioutil"
    "net/http"
)

func IsURL(url string) bool {
    urlRegeXp := "^((http|https|ftp)\\://)?([a-zA-Z0-9\\.\\-]+(\\:[a-zA-Z0-9\\.&amp;\\$\\-]+)*@)?((25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9])\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[0-9])|([a-zA-Z0-9\\-]+\\.)*[a-zA-Z0-9\\-]+\\.[a-zA-Z]{2,4})(\\:[0-9]+)?(/[^/][a-zA-Z0-9\\.\\,\\?\\'\\/\\+&amp;\\$#\\=~_\\-@]*)*$"
    match, err := regexp.MatchString(urlRegeXp, url)
    if err != nil {
        Logger.Error(errMsgRegeXp)
        panic(errMsgRegeXp)
    }
    return match
}

func ResponseToString(resp *http.Response) (string, error) {
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        errMsg := "Can't Resolve Http Response. ioutil error: " + err.Error()
        Logger.Error(errMsg)
        return "", errors.New(errMsg)
    } else {
        return string(body), nil
    }
}

func HttpGet(url, username, password string) (*http.Response, error) {
    return Http(HTTP_METHOD_GET, url, username, password, nil)
}

func HttpPost(url, username, password string, data []byte) (*http.Response, error) {
    return Http(HTTP_METHOD_POST, url, username, password, data)
}

func HttpDelete(url, username, password string, data []byte) (*http.Response, error) {
    return Http(HTTP_METHOD_DELETE, url, username, password, data)
}

func HttpPut(url, username, password string, data []byte) (*http.Response, error) {
    return Http(HTTP_METHOD_PUT, url, username, password, data)
}

func Http(method, url, username, password string, data []byte) (*http.Response, error) {
    if !IsValidHttpMethod(method) {
        errMsg := "fail to send http request, parameter <method> is null."
        Logger.Error(errMsg)
        return nil, errors.New(errMsg)
    }

    if url == "" {
        errMsg := "fail to send http request, parameter <url> is null."
        Logger.Error(errMsg)
        return nil, errors.New(errMsg)
    }

    Logger.Info("send http request, method:" + method + ", url:" + url)

    var b io.Reader = nil

    if data != nil {
        b = bytes.NewBuffer(data)
        Logger.Info("send http request, data:" + string(data))
    }

    client := &http.Client{}
    req, err := http.NewRequest(method, url, b)
    if err != nil {
        errMsg := "fail to send http request, " + err.Error()
        Logger.Error(errMsg)
        return nil, errors.New(errMsg)
    }
    req.Header.Add("Content-Type", "application/json")
    if username != "" && password != "" {
        req.SetBasicAuth(username, password)
    }
    resp, err := client.Do(req)
    if err != nil {
        errMsg := "fail to send http request, " + err.Error()
        Logger.Error(errMsg)
        return nil, errors.New(errMsg)
    }
    return resp, nil
}
