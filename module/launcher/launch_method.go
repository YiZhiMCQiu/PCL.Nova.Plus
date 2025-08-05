package launcher

import (
	"NovaPlus/module/mmcll"
	"context"
	"encoding/json"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type MCConfig struct {
	Path string `json:"path"`
	Name string `json:"name"`
}
type MCConfigs struct {
	Mc []MCConfig `json:"mc"`
}
type JavaConfig struct {
	Path    string `json:"path"`
	Version string `json:"version"`
	Arch    string `json:"arch"`
	Vendor  string `json:"vendor"`
}
type JavaConfigs struct {
	Java []JavaConfig `json:"java"`
}
type LaunchMethod struct {
	Ctx context.Context
}

func (lm *LaunchMethod) GetMCVersionConfig() MCConfigs {
	cur := filepath.Join(GetCurrentExeDir(), "PCL.Nova", "config", "MCJson.json")
	err := EnsureConfigFile(cur)
	if err != nil {
		panic(err)
	}
	if content, _ := mmcll.GetFile(cur); content == "" {
		err := mmcll.SetFile(cur, "{\n\t\"mc\": []\n}")
		if err != nil {
			panic(err)
		}
		return MCConfigs{Mc: make([]MCConfig, 0)}
	} else {
		var arr MCConfigs
		err = json.Unmarshal([]byte(content), &arr)
		if err != nil {
			panic(err)
		}
		return arr
	}
}

func (lm *LaunchMethod) SetMCVersionConfig(mc MCConfigs) {
	cur := filepath.Join(GetCurrentExeDir(), "PCL.Nova", "config", "MCJson.json")
	err := EnsureConfigFile(cur)
	if err != nil {
		panic(err)
	}
	res, err := json.MarshalIndent(mc, "", "\t")
	if err != nil {
		panic(err)
	}
	err = mmcll.SetFile(cur, string(res))
	if err != nil {
		panic(err)
	}
}

func (lm *LaunchMethod) GetJavaConfig() JavaConfigs {
	cur := filepath.Join(GetCurrentExeDir(), "PCL.Nova", "config", "JavaJson.json")
	err := EnsureConfigFile(cur)
	if err != nil {
		panic(err)
	}
	if content, err := mmcll.GetFile(cur); err != nil {
		err := mmcll.SetFile(cur, "{\n\t\"java\": []\n}")
		if err != nil {
			panic(err)
		}
		return JavaConfigs{Java: make([]JavaConfig, 0)}
	} else {
		var arr JavaConfigs
		err = json.Unmarshal([]byte(content), &arr)
		if err != nil {
			panic(err)
		}
		return arr
	}
}
func (lm *LaunchMethod) SetJavaConfig(java JavaConfigs) {
	cur := filepath.Join(GetCurrentExeDir(), "PCL.Nova", "config", "JavaJson.json")
	err := EnsureConfigFile(cur)
	if err != nil {
		panic(err)
	}
	res, err := json.MarshalIndent(java, "", "\t")
	if err != nil {
		panic(err)
	}
	err = mmcll.SetFile(cur, string(res))
	if err != nil {
		panic(err)
	}
}

// GetMCAllVersion 从 versions 目录下，提取出所有的 MC 版本并返回给前端
func (lm *LaunchMethod) GetMCAllVersion(rootPath string) []string {
	rootPath = filepath.Join(rootPath, "versions")
	stat, err := os.Stat(rootPath)
	if err != nil || !stat.IsDir() {
		return []string{}
	}
	dir, err := os.Open(rootPath)
	if err != nil {
		return []string{}
	}
	defer dir.Close()
	dirs, err := dir.Readdir(-1)
	if err != nil {
		return []string{}
	}
	var result []string
	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}
		result = append(result, filepath.Join(rootPath, d.Name()))
	}
	return result
}

func (lm *LaunchMethod) GetCurrentMinecraftDir() string {
	return filepath.Join(GetCurrentExeDir(), ".minecraft")
}

func (lm *LaunchMethod) GetJavaInfo(path string) JavaConfig {
	cmd := CMD(path, "-XshowSettings:properties", "-version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return JavaConfig{
			Path:    err.Error(),
			Vendor:  "",
			Version: "",
			Arch:    "",
		}
	}
	outStr := string(out)
	outStrSlice := strings.Split(outStr, "\n")
	result := JavaConfig{}
	result.Path = path
	for _, line := range outStrSlice {
		lc := strings.Split(line, "=")
		if len(lc) == 2 {
			lc0 := strings.Trim(lc[0], " \n\r")
			if lc0 == "java.vendor" {
				result.Vendor = strings.Trim(lc[1], " \n\r")
			} else if lc0 == "java.version" {
				result.Version = strings.Trim(lc[1], " \n\r")
			} else if lc0 == "os.arch" {
				result.Arch = strings.Trim(lc[1], " \n\r")
			}
		}
	}
	return result
}
func (lm *LaunchMethod) LaunchGame() string {
	rw := ReaderWriter{
		Ctx: lm.Ctx,
	}
	acc := AccountMethod{
		Ctx: lm.Ctx,
	}

	configIniPath := rw.GetConfigIniPath()
	accounts := acc.GetAccountConfig().Accounts
	accIndexStr := rw.ReadConfig(rw.GetOtherIniPath(), "Account", "SelectAccount")
	accIndex, err := strconv.Atoi(accIndexStr)
	if err != nil {
		return "Cannot convert account index to int"
	}
	account, ok := mmcll.SafeGet(accounts, accIndex)
	if !ok {
		return "Account array index out of range!"
	}
	var accountStruct mmcll.LaunchAccount
	if account.Name != "" && account.UUID != "" {
		accountStruct = mmcll.NewLaunchAccountOffline(account.Name, account.UUID)
		if account.AccessToken != nil && *account.AccessToken != "" {
			accountStruct = mmcll.NewLaunchAccountMicrosoft(account.Name, account.UUID, *account.AccessToken)
			if account.Server != nil && *account.Server != "" && account.BaseCode != nil && *account.BaseCode != "" {
				accountStruct = mmcll.NewLaunchAccountThirdParty(account.Name, account.UUID, *account.AccessToken, *account.BaseCode, *account.Server)
			}
		}
	} else {
		return "Account Name or UUID is Empty!"
	}
	javas := lm.GetJavaConfig().Java
	javaIndexStr := rw.ReadConfig(configIniPath, "Java", "SelectJava")
	javaIndex, err := strconv.Atoi(javaIndexStr)
	if err != nil {
		return "Cannot convert java index to int"
	}
	java, ok := mmcll.SafeGet(javas, javaIndex)
	if !ok {
		return "Java array index out of range!"
	}
	javaPath := java.Path
	roots := lm.GetMCVersionConfig().Mc
	rootIndexStr := rw.ReadConfig(configIniPath, "MC", "SelectMC")
	rootIndex, err := strconv.Atoi(rootIndexStr)
	if err != nil {
		return "Cannot convert root index to int"
	}
	var root MCConfig
	if rootIndex == 0 {
		root = MCConfig{
			Path: lm.GetCurrentMinecraftDir(),
		}
	} else {
		root, ok = mmcll.SafeGet(roots, rootIndex-1)
		if !ok {
			return "Root array index out of range!"
		}
	}
	rootPath := root.Path
	vers := lm.GetMCAllVersion(rootPath)
	verIndexStr := rw.ReadConfig(configIniPath, "MC", "SelectVer")
	verIndex, err := strconv.Atoi(verIndexStr)
	if err != nil {
		return "Cannot convert root version to int"
	}
	ver, ok := mmcll.SafeGet(vers, verIndex)
	if !ok {
		return "Version array index out of range!"
	}
	isolation := rw.ReadConfig(configIniPath, "Document", "SelectIsolation") == "4"
	gamePath := mmcll.If(isolation, ver, rootPath).(string)
	option := mmcll.NewLaunchOption(accountStruct, javaPath, rootPath, ver, gamePath)
	heightStr := rw.ReadConfig(configIniPath, "Document", "WindowHeight")
	if height, err := strconv.Atoi(heightStr); err == nil {
		option.SetWindowHeight(uint32(height))
	}
	widthStr := rw.ReadConfig(configIniPath, "Document", "WindowWidth")
	if width, err := strconv.Atoi(widthStr); err == nil {
		option.SetWindowWidth(uint32(width))
	}
	maxMemoryStr := rw.ReadConfig(configIniPath, "Document", "MaxMemory")
	if maxMemory, err := strconv.Atoi(maxMemoryStr); err == nil {
		option.SetMaxMemory(uint32(maxMemory))
	}
	customInfo := rw.ReadConfig(configIniPath, "Version", "CustomInfo")
	if customInfo != "" {
		option.SetCustomInfo(customInfo)
	} else {
		option.SetCustomInfo("PCL.Nova.Plus")
	}
	additionalJVM := rw.ReadConfig(configIniPath, "Version", "AdditionalJVM")
	if additionalJVM != "" {
		option.SetAdditionalJvm(additionalJVM)
	}
	additionalGame := rw.ReadConfig(configIniPath, "Version", "AdditionalGame")
	if additionalGame != "" {
		option.SetAdditionalGame(additionalGame)
	}
	err = mmcll.LaunchGame(*option, true, func(back []string) {
		runtime.EventsEmit(lm.Ctx, "launch_success")
		cmd := CMD(back[0], back[1:]...)
		e := cmd.Run()
		if e != nil {
			panic(e)
		}
		//stdout, _ := cmd.StdoutPipe()
		//e := cmd.Start()
		//if e != nil {
		//	panic(e)
		//	return
		//}
		//scanner := bufio.NewScanner(stdout)
		//for scanner.Scan() {
		//	fmt.Println(scanner.Text())
		//	runtime.EventsEmit(lm.Ctx, "game_log", scanner.Text())
		//}
		//e = cmd.Wait()
		//if e != nil {
		//	panic(e)
		//	return
		//}
	})
	if err != nil {
		return "You have some error! please try to view it: " + err.Error()
	}
	return ""
}
