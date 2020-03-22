package http_proxy_client

import (
	"fmt"
	"golang.org/x/net/proxy"
	"log"
	"net/http"
)

const (
	PROXY_IP       = "45.151.100.139"
	PROXY_PORT     = "8000"
	PROXY_USERNAME = "V062gV"
	PROXY_PASS     = "QnHsds"
)

func NewClientWithProxy() (client *http.Client) {
	//proxyIp, exist := os.LookupEnv("PROXY")
	//if !exist {
	//	log.Fatal("NOT FOUND PROXY")
	//	return
	//}

	fmt.Println(PROXY_IP, PROXY_PORT, PROXY_USERNAME, PROXY_PASS)

	auth := proxy.Auth{
		User:     PROXY_USERNAME,
		Password: PROXY_PASS,
	}

	dialer, err := proxy.SOCKS5("tcp", PROXY_IP+":"+PROXY_PORT, &auth, proxy.Direct)
	if err != nil {
		log.Fatal("can't connect to the proxy:", err)
	}

	tr := &http.Transport{Dial: dialer.Dial}
	client = &http.Client{
		Transport: tr,
	}

	//proxyString := "socks5://" + proxyIp
	//proxyUrl, _ := url.Parse(proxyString)
	//
	//tr := &http.Transport{
	//	Proxy:           http.ProxyURL(proxyUrl),
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	//client = &http.Client{}
	//client.Transport = tr
	return
}
