/*
 *                        _oo0oo_
 *                       o8888888o
 *                       88" . "88
 *                       (| -_- |)
 *                       0\  =  /0
 *                     ___/`---'\___
 *                   .' \\|     |// '.
 *                  / \\|||  :  |||// \
 *                 / _||||| -:- |||||- \
 *                |   | \\\  - /// |   |
 *                | \_|  ''\---/''  |_/ |
 *                \  .-\__  '-'  ___/-. /
 *              ___'. .'  /--.--\  `. .'___
 *           ."" '<  `.___\_<|>_/___.' >' "".
 *          | | :  `- \`.;`\ _ /`;.`/ - ` : | |
 *          \  \ `_.   \_ __\ /__ _/   .-` /  /
 *      =====`-.____`.___ \_____/___.-`___.-'=====
 *                        `=---='
 *
 *
 *      ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *
 *            佛祖保佑       永不宕机     永无BUG
 */
package main

import (
	"NovaPlus/module/launcher"
	"NovaPlus/module/mmcll"
	"bufio"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"math"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// App struct
type App struct {
	Ctx context.Context
}

var app = &App{}
var readerWriter = &launcher.ReaderWriter{}
var mainMethod = &launcher.MainMethod{}
var accountMethod = &launcher.AccountMethod{}
var network = &launcher.Network{}
var launchMethod = &launcher.LaunchMethod{}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func startup(ctx context.Context) {
	app.Ctx = ctx
	readerWriter.Ctx = ctx
	mainMethod.Ctx = ctx
	accountMethod.Ctx = ctx
	network.Ctx = ctx
	launchMethod.Ctx = ctx
}

func (a *App) StartDownload() {
	for progress := 10; progress <= 100; progress += 10 {
		time.Sleep(time.Second)
		runtime.EventsEmit(a.Ctx, "download_progress", progress)
	}
	runtime.EventsEmit(a.Ctx, "download_success")
}

// GetMachineCode 在Windows上尝试生成PCL2原版的识别码~
func (a *App) GetMachineCode() string {
	return strings.ToUpper(launcher.GetMachineCode())
}

// GetTodayLucky 生成今日人品，有3个报错类型。-1：未能找到注册表中的机器码，你可能在虚拟机中运行Nova、-2：未能找到注册表中的PCL-Identify键值，你可能需要从原版PCL里主动生成一次识别码才能。-3：不支持的系统。
func (a *App) GetTodayLucky(code string) int {
	if code == "ERR_1" {
		return -1
	} else if code == "ERR_2" {
		return -2
	} else if code == "DNS" {
		return -3
	} else {
		// 以下为标准的PCL生成今日人品的代码【参数code为PCL原版识别码】~
		now := time.Now()
		str1 := fmt.Sprintf("asdfgbn%d12#3$45%dIUY", now.YearDay(), now.Year())
		str2 := fmt.Sprintf("QWERTY%s0*8&6%dkjhg", code, now.Day())
		hash1 := launcher.GetHash(str1)
		hash2 := launcher.GetHash(str2)
		combined := math.Abs(float64(hash1)/3.0 + float64(hash2)/3.0)
		num := int(math.Round(math.Mod(combined/527.0, 1001.0)))
		var lucky int
		if num >= 970 {
			lucky = 100
		} else {
			lucky = int(math.Round(float64(num) / 969.0 * 99.0))
		}
		return lucky
	}
}

//go:embed all:frontend/dist
var assets embed.FS

func runRawMain() {
	// Create an instance of the app structure
	// Create application with options
	err := wails.Run(&options.App{
		Title:     "PCL.Nova.Plus",
		Width:     1024,
		Height:    614,
		MinWidth:  1024,
		MinHeight: 614,
		Frameless: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		//BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		DragAndDrop: &options.DragAndDrop{
			EnableFileDrop: true,
		},
		OnStartup: startup,
		Bind: []any{
			app,
			readerWriter,
			accountMethod,
			mainMethod,
			network,
			launchMethod,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
func runMyTestMain() {
	defer func() {
		var m string
		_, _ = fmt.Scan(&m)
	}()
	if f, err := mmcll.GetFile(filepath.Join(launcher.GetCurrentExeDir(), "mmcll.json")); err == nil {
		var m map[string]any
		err := json.Unmarshal([]byte(f), &m)
		if err != nil {
			fmt.Println("JSON 格式有误，请重新写入文件！")
			return
		}
		name := mmcll.Safe(m, "", "name").(string)
		if name == "" {
			fmt.Println("名称没有输入，请重新输入！")
			return
		}
		uuid := mmcll.GenerateBukkitUUID(name)
		account := mmcll.NewLaunchAccountOffline(name, uuid)
		javaPath := mmcll.Safe(m, "", "java_path").(string)
		if javaPath == "" {
			fmt.Println("Java 路径没有输入，请重新输入！")
			return
		}
		rootPath := mmcll.Safe(m, "", "root_path").(string)
		if rootPath == "" {
			fmt.Println("版本根路径没有输入，请重新输入！")
			return
		}
		versPath := mmcll.Safe(m, "", "vers_path").(string)
		if versPath == "" {
			fmt.Println("版本文件路径没有输入，请重新输入！")
			return
		}
		isolation := mmcll.Safe(m, true, "isolation").(bool)
		gamePath := mmcll.If(isolation, versPath, rootPath).(string)
		option := mmcll.NewLaunchOption(account, javaPath, rootPath, versPath, gamePath)
		if height := mmcll.Safe(m, 480, "height").(int); height >= 480 {
			option.SetWindowHeight(uint32(height))
		}
		if width := mmcll.Safe(m, 854, "width").(int); width >= 854 {
			option.SetWindowWidth(uint32(width))
		}
		if memory := mmcll.Safe(m, 1024, "memory").(int); memory >= 1024 {
			option.SetMaxMemory(uint32(memory))
		}
		if info := mmcll.Safe(m, mmcll.LauncherName, "info").(string); info != mmcll.LauncherName {
			option.SetCustomInfo(info)
		}
		if addiJvm := mmcll.Safe(m, "", "addi_jvm").(string); addiJvm != "" {
			option.SetAdditionalJvm(addiJvm)
		}
		if addiGame := mmcll.Safe(m, "", "addi_game").(string); addiGame != "" {
			option.SetAdditionalGame(addiGame)
		}
		err = mmcll.LaunchGame(*option, true, func(back []string) {
			cmd := exec.Command(back[0], back[1:]...)
			stdout, _ := cmd.StdoutPipe()
			err := cmd.Start()
			if err != nil {
				fmt.Println("错误！无法运行 CMD!，错误信息：" + err.Error())
				return
			}
			scanner := bufio.NewScanner(stdout)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
			err = cmd.Wait()
			if err != nil {
				fmt.Println("错误！无法等待 CMD!，错误信息：" + err.Error())
				return
			}
		})
		if err != nil {
			fmt.Println("版本检查出错了！请将你的 MC 版本，Java 版本 以及 PCL2 能否正常启动的截图发送给作者！错误信息为：" + err.Error())
		}
	} else {
		fmt.Println("这里是 xphost 制作的 Golang 语言 MC 启动器 测试环境！请不要用其放入生产环境！本产品目前严禁制作整合包，仅限内部使用！如果你从外部获取了本产品，请于 24 小时内删除！")
		fmt.Println("目前仅限离线登录，无法使用正版，无法自主定义 UUID，仅能使用 Bukkit 方式生成 UUID")
		fmt.Println("请将 配置文件 保存到 mmcll.json 下")
		fmt.Println("配置格式为：")
		m := map[string]any{
			"name":      "[你的账户名称]",
			"java_path": "[你的 Java 路径（精确到 java.exe）]",
			"root_path": "[你的 游戏 根目录（精确到 .minecraft）]",
			"vers_path": "[你的 游戏 版本目录（精确到 .minecraft/versions/<版本名称>）]",
			"isolation": "[布尔值，是否版本隔离]",
			"height":    "[数字值，窗口高度]",
			"width":     "[数字值，窗口宽度]",
			"memory":    "[数字值，最大内存]",
			"info":      "[字符串，自定义信息，显示在游戏地图里 F3 后的左上角以及游戏主窗口的左下角]",
			"addi_jvm":  "[额外 JVM 参数，如果你不知道是什么，请不要设置此值]",
			"addi_game": "[额外 Game 参数，如果你想一进入游戏就全屏，可以设置 --fullScreen]",
		}
		res, _ := json.MarshalIndent(m, "", "\t")
		fmt.Println(string(res))
	}
}
func main() {
	runRawMain()
	//runMyTestMain()
}
