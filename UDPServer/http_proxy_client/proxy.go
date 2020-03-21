package http_proxy_client

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"os"
)

func NewClientWithProxy() (client *http.Client) {
	proxyIp, exist := os.LookupEnv("PROXY")
	if !exist {
		log.Fatal("NOT FOUND PROXY")
		return
	}
	proxyString := "socks5://" + proxyIp
	proxyUrl, _ := url.Parse(proxyString)

	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxyUrl),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{}
	client.Transport = tr
	return
}
