package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func NewHttpConnection(timeout int, proxy bool, proxyStr ...string) (*http.Client, error) {
	if !proxy {
		httpClient := &http.Client{
			Timeout: time.Duration(timeout) * time.Millisecond,
		}
		return httpClient, nil
	} else {
		proxyURL, err := url.Parse("http://" + proxyStr[0])
		if err != nil {
			return nil, err
		}
		fmt.Println(proxyURL)
		httpClient := &http.Client{

			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			},
			Timeout: time.Duration(timeout) * time.Millisecond,
		}
		return httpClient, err
	}
}

func NetworkAvailable(url string, timeout int, proxy bool, proxyStr ...string) (available bool, latency int, statusCode int) {
	httpclient, _ := NewHttpConnection(timeout, proxy, proxyStr...)

	startTime := time.Now()

	resp, err := httpclient.Get(url)
	if err != nil {
		return false, timeout, 0
	}

	defer resp.Body.Close()

	rtt := int(time.Since(startTime) / time.Millisecond)

	if resp.StatusCode == http.StatusOK {
		return true, rtt, resp.StatusCode
	} else {
		return false, rtt, resp.StatusCode
	}

}
