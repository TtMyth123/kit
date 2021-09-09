package HttpClientPoolKit

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"time"
)

type HttpClient struct {
	client        *http.Client
	isBusy bool
	mHttpClientData *HttpClientData
}

func (this *HttpClient) GetBusyStatus()bool {
	return this.isBusy
}

func (this *HttpClient) setBusyStatus(b  bool) {
	this.isBusy = b
}

func newHttpClient(timeout int, aHttpClientData *HttpClientData) *HttpClient  {
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
	client.client = &http.Client{Transport: tr}
	client.mHttpClientData = aHttpClientData
	return client
}
func (this *HttpClient) setClientData()  {
	if this.mHttpClientData!=nil{
		this.client.Jar = this.mHttpClientData.CurCookieJar
	}
}
func (this *HttpClient) setClientDataByClient(u *url.URL)  {
	if this.mHttpClientData!=nil{
		this.mHttpClientData.Cookies = this.mHttpClientData.CurCookieJar.Cookies(u)
	}
}

func (this *HttpClient) Get302Location(strUrl string) (string, error) {
	this.isBusy = true
	defer this.setBusyStatus(false)

	u, _ := url.Parse(strUrl)
	this.setClientData()

}