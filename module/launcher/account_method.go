package launcher

import (
	"NovaPlus/module/mmcll"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
)

type AccountType struct {
	AType        string  `json:"type"`
	Name         string  `json:"name"`
	UUID         string  `json:"uuid"`
	AccessToken  *string `json:"access_token,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	ClientToken  *string `json:"client_token,omitempty"`
	Server       *string `json:"server,omitempty"`
	BaseCode     *string `json:"base_code,omitempty"`
	HeadSkin     string  `json:"head_skin"`
}
type AccountList struct {
	Accounts []AccountType `json:"accounts"`
}
type AccountMethod struct {
	Ctx context.Context
}

func (m *AccountMethod) SetAccountConfig(accounts AccountList) ExceptionHandler[any] {
	exception := func(err error) ExceptionHandler[any] {
		return NewExceptionHandler[any](500, false, "File System Error: "+err.Error(), nil)
	}
	realJson, err := json.MarshalIndent(accounts, "", "\t")
	if err != nil {
		return exception(err)
	}
	str := string(realJson)
	home, err := GetHomeDir()
	if err != nil {
		exePath, err := os.Executable()
		if err != nil {
			return exception(err)
		}
		home = filepath.Join(filepath.Dir(exePath), "PCL.Nova", "config")
	}
	err = mmcll.SetFile(filepath.Join(home, "AccountJSON.json"), str)
	if err != nil {
		return exception(err)
	}
	return NewExceptionHandler[any](201, true, "Created!", nil)
}

func (m *AccountMethod) GetAccountConfig() ExceptionHandler[AccountList] {
	exception := func(err error) ExceptionHandler[AccountList] {
		return NewExceptionHandler(500, false, "File System Error: "+err.Error(), AccountList{Accounts: []AccountType{}})
	}
	home, err := GetHomeDir()
	if err != nil {
		exePath, err := os.Executable()
		if err != nil {
			return exception(err)
		}
		home = filepath.Join(filepath.Dir(exePath), "PCL.Nova", "config")
	}
	cur := filepath.Join(home, "AccountJSON.json")
	if content, _ := mmcll.GetFile(cur); content == "" {
		err := mmcll.SetFile(cur, "{\n\t\"accounts\": []\n}")
		if err != nil {
			return exception(err)
		}
		return NewExceptionHandler(201, true, "Created!", AccountList{Accounts: []AccountType{}})
	} else {
		var arr AccountList
		err = json.Unmarshal([]byte(content), &arr)
		if err != nil {
			return exception(err)
		}
		return NewExceptionHandler(200, true, "OK!", arr)
	}
}

// GetUserCode 微软登录，但是此处用到了 ClientID，这个是在 privacy.go 里面被忽略了的，请手动新建并添加一个~
func (m *AccountMethod) GetUserCode() ExceptionHandler[[]string] {
	exception := func(err error) ExceptionHandler[[]string] {
		return NewExceptionHandler(500, false, "Get User Code has Exception: "+err.Error(), []string{})
	}
	url := mmcll.NewAccountLogin(ClientId())
	userCode, deviceCode, err := url.GetUserCode()
	if err != nil {
		return exception(err)
	}
	return NewExceptionHandler(200, true, "OK!", []string{userCode, deviceCode})
}

// GetAccessToken 获取到 AccessToken！
func (m *AccountMethod) GetAccessToken(deviceCode string) ExceptionHandler[AccountType] {
	exception := func(code int, err error) ExceptionHandler[AccountType] {
		return NewExceptionHandler(code, false, "Get Access Token has Exception: "+err.Error(), AccountType{})
	}
	url := mmcll.NewAccountLogin(ClientId())
	// 如果 err 的 code == -6 的话，说明没有完成登录，需要登录一次才能继续~
	accountLogin, err := url.LoginMicrosoft(deviceCode)
	if err != nil {
		var err2 mmcll.ErrorMMCLL
		errors.As(err, &err2)
		return exception(mmcll.If(err2.Code() != -6, 500, 401).(int), err)
	}
	var accessToken = accountLogin.AccessToken()
	var refreshToken = accountLogin.RefreshToken()
	var accountType = AccountType{
		AType:        "Microsoft",
		Name:         accountLogin.Name(),
		UUID:         accountLogin.Uuid(),
		AccessToken:  &accessToken,
		RefreshToken: &refreshToken,
	}
	return NewExceptionHandler(200, true, "OK!", accountType)
}

// GetUserHeadSkin 通过 用户 UUID 获取到用户的 头像！
func (m *AccountMethod) GetUserHeadSkin(uuid string) ExceptionHandler[string] {
	exception := func(err error) ExceptionHandler[string] {
		return NewExceptionHandler(500, false, "Get Head Skin has Exception: "+err.Error(), "")
	}
	url := mmcll.NewUrlMethod(fmt.Sprintf("https://sessionserver.mojang.com/session/minecraft/profile/%s", uuid))
	get, err := url.GetDefault()
	if err != nil {
		return exception(err)
	}
	var profile map[string]any
	err = json.Unmarshal(get, &profile)
	if err != nil {
		return exception(err)
	}
	profileValue := mmcll.Safe(profile, "", "properties", 0, "value").(string)
	if profileValue == "" {
		return exception(errors.New("cannot get head skin profile json"))
	}
	decoded, err := base64.StdEncoding.DecodeString(profileValue)
	if err != nil {
		return exception(err)
	}
	var skin map[string]any
	err = json.Unmarshal(decoded, &skin)
	skinUrl := mmcll.Safe(skin, "", "textures", "SKIN", "url").(string)
	if skinUrl == "" {
		return exception(errors.New("cannot get head skin base64 decode"))
	}
	url2 := mmcll.NewUrlMethod(skinUrl)
	skinValue, err := url2.GetDefault()
	if err != nil {
		return exception(err)
	}
	// 以下是经过裁剪的！如果不需要裁剪，请删掉以下片段，并且将上述 skinValue 直接包装成 base64！
	img, _, err := image.Decode(bytes.NewBuffer(skinValue))
	if err != nil {
		return exception(err)
	}
	rgba := img.(*image.NRGBA)
	subHead := rgba.SubImage(image.Rect(8, 8, 16, 16)).(*image.NRGBA)
	var buf bytes.Buffer
	err = png.Encode(&buf, subHead)
	if err != nil {
		return exception(err)
	}
	result := base64.StdEncoding.EncodeToString(buf.Bytes())
	return NewExceptionHandler(200, true, "OK!", result)
}
