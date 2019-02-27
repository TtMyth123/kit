package httpClientKit

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

type HttpClient struct {
	client        *http.Client
	gCurCookieJar *cookiejar.Jar
	gCurCookies   []*http.Cookie
}

func GetHttpClient(guid string) *HttpClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Dial: func(netw, addr string) (net.Conn, error) {
			conn, err := net.DialTimeout(netw, addr, time.Second*5)
			if err != nil {
				return nil, err
			}
			conn.SetDeadline(time.Now().Add(time.Second * 5))
			return conn, nil
		},
		ResponseHeaderTimeout: time.Second * 5,
	}

	client := &HttpClient{}
	client.client = &http.Client{Transport: tr}
	//client.client = &http.Client{}
	client.gCurCookieJar, _ = cookiejar.New(nil)
	return client
}

func (this *HttpClient) GetString1(strUrl string) (string, error) {
	resp, err := this.client.Get(strUrl)
	aa := resp.Header.Get("Set-Cookie")
	resp.Header.Get("")
	aaa := resp.Cookies()
	fmt.Println("Cookies:", aaa)
	fmt.Println("Set-Cookie:", aa)

	//urlX, _ := url.Parse("http://zhanzhang.baidu.com")
	//j.SetCookies(urlX, clist)
	//
	//fmt.Printf("Jar cookie : %v", j.Cookies(urlX))
	//
	//fmt.Println("Set-Cookie11:", this.client.Jar)
	//this.client.Jar.SetCookies()

	//resp.Header.
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), err
}

func (this *HttpClient) GetString(strUrl string) (string, error) {
	u, err := url.Parse(strUrl)
	if err != nil {
		return "", err
	}
	this.client.Jar = this.gCurCookieJar
	resp, err := this.client.Get(strUrl)
	this.gCurCookies = this.gCurCookieJar.Cookies(u)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), err
}

func (this *HttpClient) GetBytes(strUrl string) ([]byte, error) {
	u, err := url.Parse(strUrl)
	if err != nil {
		return nil, err
	}

	this.client.Jar = this.gCurCookieJar
	resp, err := this.client.Get(strUrl)
	this.gCurCookies = this.gCurCookieJar.Cookies(u)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func (this *HttpClient) PostForm(strUrl string, data url.Values) (string, error) {
	u, err := url.Parse(strUrl)
	if err != nil {
		return "", err
	}
	this.client.Jar = this.gCurCookieJar
	resp, err := this.client.PostForm(strUrl, data)
	this.gCurCookies = this.gCurCookieJar.Cookies(u)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), err
}
