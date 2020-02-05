package httpKit

import (
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func GetData(url string) ([]byte, error) {
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
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)

	_ = resp.Body.Close()
	return body, err
}

 