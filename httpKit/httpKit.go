package httpKit

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/TtMyth123/kit/fileKit"
	"github.com/TtMyth123/kit/httpKit/TmpFileKit"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
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

func GetImgUrl(rootUrl, ImgUrl string) string {
	if ImgUrl == "" {
		return ""
	}
	//ImgUrl = strings.ToLower(ImgUrl)
	if ImgUrl[1:1] == `/` {
		return rootUrl + ImgUrl
	} else {
		if strings.Index(ImgUrl, `://`) > 0 {
			return ImgUrl
		} else {
			return rootUrl + `/` + ImgUrl
		}
	}
}

func UploadFile(f multipart.File, head *multipart.FileHeader, path, fileName, ext string) (filePath string, err error) {
	defer f.Close()
	file, _ := ioutil.ReadAll(f)
	ext1 := filepath.Ext(head.Filename)
	fileKit.CreateMutiDir(path)

	if ext != "" {
		ext = "." + ext
	} else {
		ext = ext1
	}

	localpath := fmt.Sprintf("%s/%s%s", path, fileName, ext)
	//保存的路径
	err = ioutil.WriteFile(localpath, file, 0666)
	if err != nil {
		return "", err
	}
	return localpath, nil
}

func UploadTmpFile(t int, f multipart.File, head *multipart.FileHeader, path, fileName, ext string) (filePath string, err error) {
	filePath, err = UploadFile(f, head, path, fileName, ext)
	if err == nil {
		TmpFileKit.AddFile(fileName, filePath, t)
	}

	return filePath, err
}

func GetUrl(url1 string, param map[string]interface{}) (string, error) {
	U, e := url.Parse(url1)
	if e != nil {
		return url1, e
	}

	newUrl := url1
	strP := ``
	for k, v := range param {
		strP += fmt.Sprintf(`&%s=%v`, k, v)
	}

	if len(strP) > 0 {
		if len(U.RawQuery) > 0 {
			newUrl = fmt.Sprintf(`%s&%s`, newUrl, strP[1:])
		} else {
			newUrl = fmt.Sprintf(`%s?%s`, newUrl, strP[1:])
		}

	}

	return newUrl, nil
}

func GetParamValue(r *http.Request) map[string]string {
	mp := make(map[string]string)
	if r == nil {
		return mp
	}
	if r.Form == nil {
		r.ParseForm()
	}
	for k, v := range r.Form {
		mp[k] = v[0]
	}

	return mp
}

func GetPayloadData(r *http.Request) map[string]string {
	mp := make(map[string]string)
	buf := make([]byte, 1024)
	n, _ := r.Body.Read(buf)
	err := json.Unmarshal(buf[0:n], &mp)
	if err == nil {
		return mp
	}
	return mp
}

func GetParamObject(s map[string]string, data interface{}) ([]string, error) {
	ff := make([]string, 0)
	getV := reflect.ValueOf(data)
	getV = getV.Elem()
	var e error
	getType := reflect.TypeOf(getV.Interface())
	for fieldName, v := range s {
		_, ok := getType.FieldByName(fieldName)
		if ok {
			bb := getV.FieldByName(fieldName)
			k := bb.Kind()
			switch k {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				vv, e := strconv.ParseInt(v, 10, 0)
				if e != nil {
					return ff, e
				}

				ff = append(ff, fieldName)
				bb.SetInt(vv)

			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				vv, e := strconv.ParseUint(v, 10, 0)
				if e != nil {
					return ff, e
				}
				ff = append(ff, fieldName)
				bb.SetUint(vv)
			case reflect.Float32, reflect.Float64:
				vv, e := strconv.ParseFloat(v, 10)
				if e != nil {
					return ff, e
				}
				ff = append(ff, fieldName)
				bb.SetFloat(vv)
			case reflect.String:
				ff = append(ff, fieldName)
				bb.SetString(v)
			default:
				e = fmt.Errorf("无法识别%s的%s", fieldName, k)
			}
		}

	}

	return ff, e
}
