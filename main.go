package main

import (
	"NovaPlus/module/launcher"
	"NovaPlus/module/utils"
	"context"
	"embed"
	"fmt"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"math"
	"strings"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) StartDownload() {
	for progress := 10; progress <= 100; progress += 10 {
		time.Sleep(time.Second)
		runtime.EventsEmit(a.ctx, "download_progress", progress)
	}
	runtime.EventsEmit(a.ctx, "download_success")
}

// GetMachineCode 在Windows上尝试生成PCL2原版的识别码~
func (a *App) GetMachineCode() string {
	return strings.ToUpper(utils.GetMachineCode())
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

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "PCL2.Nova.App",
		Width:     1024,
		Height:    614,
		MinWidth:  1024,
		MinHeight: 614,
		Frameless: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		//BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
			&launcher.ReaderWriter{},
			&launcher.MainMethod{},
			&launcher.Account{},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
