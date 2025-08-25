package launcher

import (
	"NovaPlus/module/mmcll"
	"context"
	"encoding/json"
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

func (account *AccountMethod) SetAccountConfig(accounts AccountList) ExceptionHandler[any] {
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

func (account *AccountMethod) GetAccountConfig() ExceptionHandler[AccountList] {
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
		return NewExceptionHandler(201, true, "OK!", AccountList{Accounts: []AccountType{}})
	} else {
		var arr AccountList
		err = json.Unmarshal([]byte(content), &arr)
		if err != nil {
			return exception(err)
		}
		return NewExceptionHandler(200, true, "OK!", arr)
	}
}
