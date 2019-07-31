package httpClientKit

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type HandlerFunc_Response func(resp *http.Response)

type HttpClient struct {
	client        *http.Client
	gCurCookieJar *cookiejar.Jar
	gCurCookies   []*http.Cookie
	start         time.Time
}

func GetHttpClient(guid string) *HttpClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Dial: func(netw, addr string) (net.Conn, error) {
			conn, err := net.DialTimeout(netw, addr, time.Second*15)
			if err != nil {
				return nil, err
			}
			conn.SetDeadline(time.Now().Add(time.Second * 15))
			return conn, nil
		},
		ResponseHeaderTimeout: time.Second * 15,
	}

	client := &HttpClient{}
	client.client = &http.Client{Transport: tr}
	//client.client = &http.Client{}
	client.gCurCookieJar, _ = cookiejar.New(nil)

	client.start = time.Now()
	return client
}
func (this *HttpClient) Clear() {
	this.gCurCookieJar, _ = cookiejar.New(nil)
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

	LoopI := 0
	//time.Sleep(time.Millisecond * 200)
	u, err := url.Parse(strUrl)
	if err != nil {
		return "", err
	}
	this.client.Jar = this.gCurCookieJar

	LoopIA := 0
LoopGet:
	LoopIA++
	LoopI++;
	resp, err := this.client.Get(strUrl)
	this.gCurCookies = this.gCurCookieJar.Cookies(u)

	if err != nil {
		strE := err.Error()
		if strings.Contains(strE, "timeout") && LoopI < 4 {
			goto LoopGet
		}
		fmt.Println(strE)
		return "", err
	}
	defer resp.Body.Close()

	LoopI = 0
LoopReadAll:
	LoopI++
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		strE := err.Error()
		if strings.Contains(strE, "timeout") && LoopI < 4 {
			goto LoopReadAll
		}

		if LoopIA < 3 {
			goto LoopGet
		}

		fmt.Println(strE)
		return "", err
	}
	strType := resp.Header.Get("Content-Type")
	if strings.Contains(strType, "charset=gb2312") {
		decodeBytes, _ := simplifiedchinese.GBK.NewDecoder().Bytes(body)
		return string(decodeBytes), err
	}
	return string(body), err
}
func (this *HttpClient) GetBytesEx(strUrl string) ([]byte, error) {
	LoopI := 0
	u, err := url.Parse(strUrl)
	if err != nil {
		return nil, err
	}

	this.client.Jar = this.gCurCookieJar
	LoopIA := 0
LoopGet:
	LoopIA++
	LoopI++;
	resp, err := this.client.Get(strUrl)
	this.gCurCookies = this.gCurCookieJar.Cookies(u)

	if err != nil {
		strE := err.Error()
		if strings.Contains(strE, "timeout") && LoopI < 4 {
			goto LoopGet
		}
		fmt.Println(strE)
		return nil, err
	}
	defer resp.Body.Close()

	LoopI = 0
LoopReadAll:
	LoopI++
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		strE := err.Error()
		if strings.Contains(strE, "timeout") && LoopI < 4 {
			goto LoopReadAll
		}

		if LoopIA < 3 {
			goto LoopGet
		}

		fmt.Println(strE)
		return nil, err
	}
	return body, err
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

func (this *HttpClient) GetStringCallBark(strUrl string, callBark HandlerFunc_Response) (string, error) {
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
	if callBark != nil {
		callBark(resp)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), err
}

func (this *HttpClient) DoRequest(method, strUrl string, paramsHeader map[string]string, data io.Reader) ([]byte, error) {
	u, err := url.Parse(strUrl)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest(method, strUrl, data)
	this.client.Jar = this.gCurCookieJar

	if paramsHeader != nil {
		for k, v := range paramsHeader {
			req.Header.Set(k, v)
		}
	}

	resp, err := this.client.Do(req)
	this.gCurCookies = this.gCurCookieJar.Cookies(u)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}

func (this *HttpClient) GetHeader(strUrl string, paramsHeader map[string]string) ([]byte, error) {
	u, err := url.Parse(strUrl)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("GET", strUrl, nil)
	this.client.Jar = this.gCurCookieJar
	if paramsHeader != nil {
		for k, v := range paramsHeader {
			req.Header.Set(k, v)
		}
	}

	resp, err := this.client.Do(req)
	this.gCurCookies = this.gCurCookieJar.Cookies(u)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return body, err
}
  
func (this *HttpClient) PostFormHeader(strUrl string, paramsHeader map[string]string, data url.Values) (string, error) {
	u, err := url.Parse(strUrl)
	if err != nil {
		return "", err
	}

	req, _ := http.NewRequest("POST", strUrl, strings.NewReader(data.Encode()))
	this.client.Jar = this.gCurCookieJar

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if paramsHeader != nil {
		for k, v := range paramsHeader {
			req.Header.Set(k, v)
		}
	}

	resp, err := this.client.Do(req)
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

func (this *HttpClient) PostForm(strUrl string, data url.Values) (string, error) {
	return this.PostFormCallBark(strUrl, data, nil)
}

func (this *HttpClient) PostFormCallBark(strUrl string, data url.Values, callBark HandlerFunc_Response) (string, error) {
	u, err := url.Parse(strUrl)
	if err != nil {
		return "", err
	}

	this.client.Jar = this.gCurCookieJar
	//cost := time.Since(this.start)
	//if cost < time.Millisecond*200 {
	//	time.Sleep(time.Millisecond*200 - cost)
	//}
	resp, err := this.client.PostForm(strUrl, data)
	//this.start = time.Now()

	this.gCurCookies = this.gCurCookieJar.Cookies(u)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if callBark != nil {
		callBark(resp)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), err
}

func (this *HttpClient) PostJson(strUrl, strJsonData string) (string, error) {
	u, err := url.Parse(strUrl)
	if err != nil {
		return "", err
	}

	this.client.Jar = this.gCurCookieJar
	contentType := `application/json`
	resp, err := this.client.Post(strUrl, contentType, strings.NewReader(strJsonData))
	this.gCurCookies = this.gCurCookieJar.Cookies(u)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}
