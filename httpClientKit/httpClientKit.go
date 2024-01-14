package httpClientKit

import (
	"crypto/tls"
	"fmt"
	"github.com/TtMyth123/kit"
	"github.com/TtMyth123/kit/httpKit"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

type HandlerFunc_Response func(resp *http.Response)

type HttpClient struct {
	client        *http.Client
	gCurCookieJar *cookiejar.Jar
	gCurCookies   []*http.Cookie
	start         time.Time
	Id            string
	RetryI        int
	isBusy        bool
}

var timeout = 30

func SetTimeout(t int) {
	timeout = t
}

func (this *HttpClient) GetBusyStatus() bool {
	return this.isBusy
}

func (this *HttpClient) setBusyStatus(b bool) {
	this.isBusy = b
}
func (this *HttpClient) GetClient() *http.Client {
	return this.client
}

func GetHttpClient(guid string) *HttpClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Dial: func(netw, addr string) (net.Conn, error) {
			conn, err := net.DialTimeout(netw, addr, time.Second*time.Duration(timeout))
			if err != nil {
				return nil, err
			}
			conn.SetDeadline(time.Now().Add(time.Second * time.Duration(timeout)))
			return conn, nil
		},
		ResponseHeaderTimeout: time.Second * time.Duration(timeout),
	}

	client := &HttpClient{}
	if guid == "" {
		guid = kit.GetGuid()
	}
	client.client = &http.Client{Transport: tr}
	//client.client = &http.Client{}
	client.gCurCookieJar, _ = cookiejar.New(nil)
	client.Id = guid

	client.start = time.Now()
	return client
}

func SetHttpProxy(guid, HttpProxyUrl string) *HttpClient {
	if HttpProxyUrl == "" {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*time.Duration(timeout))
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * time.Duration(timeout)))
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * time.Duration(timeout),
		}

		client := &HttpClient{}
		if guid == "" {
			guid = kit.GetGuid()
		}
		client.client = &http.Client{Transport: tr}
		//client.client = &http.Client{}
		client.gCurCookieJar, _ = cookiejar.New(nil)
		client.Id = guid

		client.start = time.Now()
		return client
	} else {
		ProxyURL, _ := url.Parse(HttpProxyUrl)
		tr := &http.Transport{
			Proxy:           http.ProxyURL(ProxyURL),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*time.Duration(timeout))
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * time.Duration(timeout)))
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * time.Duration(timeout),
		}

		client := &HttpClient{}
		if guid == "" {
			guid = kit.GetGuid()
		}
		client.client = &http.Client{Transport: tr}
		//client.client = &http.Client{}
		client.gCurCookieJar, _ = cookiejar.New(nil)
		client.Id = guid

		client.start = time.Now()
		return client
	}
}

func GetProxyHttp(guid, HttpProxyUrl string) *HttpClient {
	if HttpProxyUrl == "" {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*time.Duration(timeout))
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * time.Duration(timeout)))
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * time.Duration(timeout),
		}

		client := &HttpClient{}
		if guid == "" {
			guid = kit.GetGuid()
		}
		client.client = &http.Client{Transport: tr}
		//client.client = &http.Client{}
		client.gCurCookieJar, _ = cookiejar.New(nil)
		client.Id = guid

		client.start = time.Now()
		return client
	} else {
		ProxyURL, _ := url.Parse(HttpProxyUrl)
		tr := &http.Transport{
			Proxy:           http.ProxyURL(ProxyURL),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*time.Duration(timeout))
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * time.Duration(timeout)))
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * time.Duration(timeout),
		}

		client := &HttpClient{}
		if guid == "" {
			guid = kit.GetGuid()
		}
		client.client = &http.Client{Transport: tr}
		//client.client = &http.Client{}
		client.gCurCookieJar, _ = cookiejar.New(nil)
		client.Id = guid

		client.start = time.Now()
		return client
	}
}

func (this *HttpClient) Clear() {
	this.gCurCookieJar, _ = cookiejar.New(nil)
}

func (this *HttpClient) CurCookieJa() *cookiejar.Jar {
	return this.gCurCookieJar
}
func (this *HttpClient) GetString1(strUrl string) (string, error) {
	this.isBusy = true
	defer this.setBusyStatus(false)
	resp, err := this.client.Get(strUrl)
	if err != nil {
		return "", err
	}
	aa := resp.Header.Get("Set-Cookie")
	resp.Header.Get("")
	aaa := resp.Cookies()
	fmt.Println("Cookies:", aaa)
	fmt.Println("Set-Cookie:", aa)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	LoopI := 0
LoopReadAll:
	LoopI++
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		strE := err.Error()
		if strings.Contains(strE, "timeout") && LoopI < this.RetryI {
			goto LoopReadAll
		}
		return "", err
	}

	return string(body), err
}

func (this *HttpClient) Get302Location(strUrl string) (string, error) {
	this.isBusy = true
	defer this.setBusyStatus(false)

	u, _ := url.Parse(strUrl)
	this.client.Jar = this.gCurCookieJar
	resp, _ := this.client.Get(strUrl)
	this.gCurCookies = this.gCurCookieJar.Cookies(u)
	defer resp.Body.Close()
	strLocation := resp.Request.Response.Header.Get("Location")

	return strLocation, nil
}

func (this *HttpClient) GetString(strUrl string) (string, error) {
	this.isBusy = true
	defer this.setBusyStatus(false)

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
	LoopI++
	resp, err := this.client.Get(strUrl)
	this.gCurCookies = this.gCurCookieJar.Cookies(u)

	if err != nil {
		strE := err.Error()
		if strings.Contains(strE, "timeout") && LoopI < this.RetryI {
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
		if strings.Contains(strE, "timeout") && LoopI < this.RetryI {
			goto LoopReadAll
		}

		if LoopIA < this.RetryI {
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

func (this *HttpClient) GetStringParam(strUrl string, param map[string]interface{}) (string, error) {
	newUrl, _ := httpKit.GetUrl(strUrl, param)

	return this.GetString(newUrl)
}

func (this *HttpClient) GetBytesEx(strUrl string) ([]byte, error) {
	this.isBusy = true
	defer this.setBusyStatus(false)

	LoopI := 0
	u, err := url.Parse(strUrl)
	if err != nil {
		return nil, err
	}

	this.client.Jar = this.gCurCookieJar
	LoopIA := 0
LoopGet:
	LoopIA++
	LoopI++
	resp, err := this.client.Get(strUrl)
	this.gCurCookies = this.gCurCookieJar.Cookies(u)

	if err != nil {
		strE := err.Error()
		if strings.Contains(strE, "timeout") && LoopI < this.RetryI {
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
		if strings.Contains(strE, "timeout") && LoopI < this.RetryI {
			goto LoopReadAll
		}

		if LoopIA < this.RetryI {
			goto LoopGet
		}

		fmt.Println(strE)
		return nil, err
	}
	return body, err
}

func (this *HttpClient) GetBytes(strUrl string) ([]byte, error) {
	this.isBusy = true
	defer this.setBusyStatus(false)

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
	LoopI := 0
LoopReadAll:
	LoopI++
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		strE := err.Error()
		if strings.Contains(strE, "timeout") && LoopI < this.RetryI {
			goto LoopReadAll
		}
		return body, err
	}
	return body, err
}

func (this *HttpClient) GetStringCallBark(strUrl string, callBark HandlerFunc_Response) (string, error) {
	this.isBusy = true
	defer this.setBusyStatus(false)

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
	LoopI := 0
LoopReadAll:
	LoopI++
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		strE := err.Error()
		if strings.Contains(strE, "timeout") && LoopI < this.RetryI {
			goto LoopReadAll
		}
		return "", err
	}
	return string(body), err
}

func (this *HttpClient) DoRequest(method, strUrl string, paramsHeader map[string]string, data io.Reader) ([]byte, error) {
	this.isBusy = true
	defer this.setBusyStatus(false)

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
	this.isBusy = true
	defer this.setBusyStatus(false)

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
	LoopI := 0
LoopReadAll:
	LoopI++
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		strE := err.Error()
		if strings.Contains(strE, "timeout") && LoopI < this.RetryI {
			goto LoopReadAll
		}
		return body, err
	}

	return body, err
}

func (this *HttpClient) PostFormHeader(strUrl string, paramsHeader map[string]string, data url.Values) (string, error) {
	this.isBusy = true
	defer this.setBusyStatus(false)

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

	LoopI := 0
LoopReadAll:
	LoopI++
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		strE := err.Error()
		if strings.Contains(strE, "timeout") && LoopI < this.RetryI {
			goto LoopReadAll
		}
		return "", err
	}
	return string(body), err
}

func (this *HttpClient) PostForm(strUrl string, data url.Values) (string, error) {
	return this.PostFormCallBark(strUrl, data, nil)
}

func (this *HttpClient) PostFormCallBark(strUrl string, data url.Values, callBark HandlerFunc_Response) (string, error) {
	this.isBusy = true
	defer this.setBusyStatus(false)

	u, err := url.Parse(strUrl)
	if err != nil {
		return "", err
	}

	this.client.Jar = this.gCurCookieJar
	resp, err := this.client.PostForm(strUrl, data)
	if err != nil {
		return "", err
	}
	this.gCurCookies = this.gCurCookieJar.Cookies(u)
	defer resp.Body.Close()
	if callBark != nil {
		callBark(resp)
	}
	LoopI := 0
LoopReadAll:
	LoopI++
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		strE := err.Error()
		if strings.Contains(strE, "timeout") && LoopI < this.RetryI {
			goto LoopReadAll
		}
		return "", err
	}
	return string(body), err
}

func (this *HttpClient) PostJson(strUrl, strJsonData string) (string, error) {
	this.isBusy = true
	defer this.setBusyStatus(false)

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
	LoopI := 0
LoopReadAll:
	LoopI++
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		strE := err.Error()
		if strings.Contains(strE, "timeout") && LoopI < this.RetryI {
			goto LoopReadAll
		}
		return "", err
	}
	return string(body), err
}

func (this *HttpClient) Post302LocationFormHeader(strUrl string, paramsHeader map[string]string, data url.Values) (string, string, error) {
	this.isBusy = true
	defer this.setBusyStatus(false)

	u, err := url.Parse(strUrl)
	if err != nil {
		return "", "", err
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
		return "", "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	strLocation := ""
	if resp.Request.Response != nil {
		strLocation = resp.Request.Response.Header.Get("Location")
	}

	return string(body), strLocation, err

}
