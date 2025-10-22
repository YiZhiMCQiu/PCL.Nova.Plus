package launcher

import (
	"NovaPlus/module/mmcll"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type DownloadMethod struct {
	Ctx context.Context
}

func UnmarshalKeyValue(data string, v *map[string]string) error {
	// 检查v是否为nil
	if v == nil {
		return fmt.Errorf("target map cannot be nil")
	}

	// 如果目标map为nil，创建一个新的map
	if *v == nil {
		*v = make(map[string]string)
	}

	// 如果输入为空字符串，直接返回
	if data == "" {
		return nil
	}

	// 按照分号分割键值对
	pairs := strings.Split(data, ";")

	for i, pair := range pairs {
		// 去除首尾空格
		trimmedPair := strings.TrimSpace(pair)

		// 如果分割后的字符串为空，说明有两个连续的分号
		if trimmedPair == "" {
			return fmt.Errorf("empty key-value pair at position %d", i)
		}

		// 查找第一个等号的位置
		equalIndex := strings.Index(trimmedPair, "=")

		// 如果没有找到等号，返回错误
		if equalIndex == -1 {
			return fmt.Errorf("no '=' found in pair '%s' at position %d", trimmedPair, i)
		}

		// 分割键和值（只分割第一个等号）
		key := strings.TrimSpace(trimmedPair[:equalIndex])
		value := strings.TrimSpace(trimmedPair[equalIndex+1:])

		// 检查键是否为空
		if key == "" {
			return fmt.Errorf("empty key in pair '%s' at position %d", trimmedPair, i)
		}

		// 将键值对存入map
		(*v)[key] = value
	}

	return nil
}

func (d DownloadMethod) StartDownload(url string, savePath string) string {
	rw := &ReaderWriter{
		Ctx: d.Ctx,
	}
	var headers map[string]string = nil
	h := rw.ReadConfig(rw.GetConfigIniPath(), "Download", "Headers")
	if h != "" {
		err := json.Unmarshal([]byte(h), &headers)
		if err != nil {
			headers = nil
		}
	}
	var cookies map[string]string = nil
	c := rw.ReadConfig(rw.GetConfigIniPath(), "Download", "Cookies")
	if c != "" {
		err := UnmarshalKeyValue(c, &cookies)
		if err != nil {
			cookies = nil
		}
	}
	bt := rw.ReadConfig(rw.GetConfigIniPath(), "Version", "ThreadBiggest")
	btn, err := strconv.Atoi(bt)
	if err != nil || btn < 1 || btn > 256 {
		btn = 64
	}
	proxyType := rw.ReadConfig(rw.GetConfigIniPath(), "Download", "ProxyType")
	proxyUrl := rw.ReadConfig(rw.GetConfigIniPath(), "Download", "ProxyUrl")
	proxyPort := rw.ReadConfig(rw.GetConfigIniPath(), "Download", "ProxyPort")
	proxyUsername := rw.ReadConfig(rw.GetConfigIniPath(), "Download", "ProxyUsername")
	proxyPassword := rw.ReadConfig(rw.GetConfigIniPath(), "Download", "ProxyPassword")
	isSocks := proxyType == "3"
	isHttps := proxyType == "2"
	proxy := mmcll.NewProxy(proxyUrl, proxyPort, proxyUsername, proxyPassword, isHttps, isSocks)
	dwhc := mmcll.NewDownloadWithHeadersCookies(url, savePath, btn, 0, proxy, headers, cookies)
	dwhc.StartDownload(func(url string, _ int, errmessage string, status int, all int64, current int64) {
		runtime.EventsEmit(d.Ctx, "DownloadCurrent", url, errmessage, status, all, current)
	})
	return ""
}
