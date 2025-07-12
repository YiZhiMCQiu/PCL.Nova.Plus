package mmcll

import (
	"bytes"
	"io"
	"net/http"
)

type UrlMethod struct {
	// url 需要请求的网址
	url string
}

// NewUrlMethod 生成一个UrlMethod实例
func NewUrlMethod(url string) *UrlMethod {
	return &UrlMethod{url: url}
}

// baseRequest 基础的请求函数
func (url *UrlMethod) baseRequest(reqMethod string, header map[string]string, data []byte) ([]byte, error) {
	if reqMethod == "" {
		reqMethod = http.MethodGet
	}
	if header == nil {
		header = make(map[string]string)
		header["Content-Type"] = "application/x-www-form-urlencoded"
	}
	header["User-Agent"] = LauncherUserAgent
	req, err := http.NewRequest(reqMethod, url.url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Post 请求
func (url *UrlMethod) Post(body string, isJSON bool) (string, error) {
	header := make(map[string]string)
	if isJSON {
		header["Content-Type"] = "application/json;charset=UTF-8"
		header["Accept"] = "application/json"
	} else {
		header["Content-Type"] = "application/x-www-form-urlencoded"
	}
	res, err := url.baseRequest(http.MethodPost, header, []byte(body))
	return string(res), err
}

// Get 请求，附带一个验证key的请求
func (url *UrlMethod) Get(authorization string) (string, error) {
	res, err := url.baseRequest(http.MethodGet, map[string]string{"Authorization": authorization}, []byte(""))
	return string(res), err
}

// GetDefault 默认请求，不带Authorization的请求
func (url *UrlMethod) GetDefault() (string, error) {
	res, err := url.baseRequest(http.MethodGet, nil, []byte(""))
	return string(res), err
}

type AccountResult struct {
	name         string
	uuid         string
	accessToken  string
	refreshToken string
	clientToken  string
	base         string
}

// AccountLogin key值在正版登录时，请输入client_id，如果是外置登录，请输入请求根网址【需要精确到/api/yggdrasil】。
type AccountLogin struct {
	key string
}
