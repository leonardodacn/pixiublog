package utils

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

func HttpGet(url string, param map[string]string) (string, error) {

	if url == "" {
		return "", errors.Errorf("url %s is not exists", url)
	}
	paramStr := ""
	for k, v := range param {
		paramStr += k + "=" + v + "&"
	}
	paramStr = strings.TrimRight(paramStr, "&")

	if paramStr != "" {
		if strings.Contains(url, "?") {
			url += "&" + paramStr
		} else {
			url += "?" + paramStr
		}
	}

	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func HttpPost(url string, contentType string, body io.Reader) (string, error) {

	resp, err := http.Post(url, contentType, body)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(resBody), nil
}
