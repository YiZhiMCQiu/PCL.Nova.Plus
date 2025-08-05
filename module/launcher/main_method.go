package launcher

import (
	"NovaPlus/module/mmcll"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/shirou/gopsutil/mem"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io/fs"
	"math"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"regexp"
	runtime2 "runtime"
	"strings"
	"syscall"
	"time"
)

type ReaderWriter struct {
	Ctx context.Context
}
type MainMethod struct {
	Ctx context.Context
}

type IPv6Struct struct {
	Ip      string `json:"ip"`
	Success bool   `json:"success"`
}

// Ping 对本机 IPv6 进行一次 ping 操作，也就是仅进行一次 Dial 连接。
// 使用 Golang 内置库达到跨平台效果！
func Ping(host string, timeout time.Duration) error {
	cmd := PingCMD(host, timeout)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
	return cmd.Run()
}

// EnsureConfigFile 生成所有父文件夹，以及在此处生成一个文件
// 例如 D:/aa/bb/cc.json，该函数会直接生成一个空的 cc.json 出来（
func EnsureConfigFile(path string) error {
	if _, err := os.Stat(path); err == nil {
		return nil // 文件已存在
	}
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, fs.ModePerm); err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

// GetAllIPv6 获取所有本机 IPv6
// 使用 Golang 内置库达到跨平台效果！
func (mm *MainMethod) GetAllIPv6() []IPv6Struct {
	s, err := net.InterfaceAddrs()
	if err != nil {
		return []IPv6Struct{}
	}
	var result []IPv6Struct
	for _, a := range s {
		ipNet, ok := a.(*net.IPNet)
		if !ok {
			continue
		}
		ip := ipNet.IP
		if ip.To4() == nil && ip.To16() != nil && !ip.IsLoopback() && !ip.IsLinkLocalUnicast() {
			if err = Ping(ip.String(), time.Second); err == nil {
				result = append(result, IPv6Struct{Ip: ip.String(), Success: true})
			} else {
				result = append(result, IPv6Struct{Ip: ip.String(), Success: false})
			}
		}
	}
	return result
}
func (rw *ReaderWriter) WriteConfig(path, section, key, value string) {
	conf := NewConfig(path)
	err := conf.Write(section, key, value)
	if err != nil {
		panic(err)
	}
}
func (rw *ReaderWriter) ReadConfig(path, section, key string) string {
	conf := NewConfig(path)
	if v, err := conf.Read(section, key); err == nil {
		return v
	}
	return ""
}
func (rw *ReaderWriter) GetOtherIniPath() string {
	home, err := GetHomeDir()
	if err != nil {
		home = filepath.Join(GetCurrentExeDir(), "PCL.Nova", "config")
	}
	res := filepath.Join(home, "Other.ini")
	err = EnsureConfigFile(res)
	if err != nil {
		panic(err)
	}
	return res
}
func (rw *ReaderWriter) GetConfigIniPath() string {
	res := filepath.Join(GetCurrentExeDir(), "PCL.Nova", "config", "PCL2.Nova.ini")
	err := EnsureConfigFile(res)
	if err != nil {
		panic(err)
	}
	return res
}
func (rw *ReaderWriter) GetCurrentExeDir() string {
	return GetCurrentExeDir()
}
func (rw *ReaderWriter) OpenDirectoryDialog(title string) string {
	dialog, err := runtime.OpenDirectoryDialog(rw.Ctx, runtime.OpenDialogOptions{
		Title:                title,
		CanCreateDirectories: true,
	})
	if err != nil {
		return ""
	}
	return dialog
}
func (rw *ReaderWriter) OpenFileDialog(title string, filter []string) string {
	var fileFilter []runtime.FileFilter
	for _, f := range filter {
		fileFilter = append(fileFilter, runtime.FileFilter{
			DisplayName: f,
			Pattern:     f,
		})
	}
	dialog, err := runtime.OpenFileDialog(rw.Ctx, runtime.OpenDialogOptions{
		Title:                title,
		CanCreateDirectories: true,
		Filters:              fileFilter,
	})
	if err != nil {
		return ""
	}
	return dialog
}
func GetCurrentExeDir() string {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(exePath)
}
func GetHash(str string) uint64 {
	result := uint64(5381)
	for _, r := range str {
		result = (result << 5) ^ result ^ uint64(r)
	}
	return result ^ uint64(0xA98F501BC684032F)
}
func StrFill(str, code string, length int) string {
	if len(str) > length {
		return str[:length]
	}
	fillCount := length - len(str)
	var builder strings.Builder
	for i := 0; i < fillCount; i++ {
		builder.WriteString(code)
	}
	builder.WriteString(str)
	return builder.String()
}
func Mid(source string, start, length int) string {
	if source == "" || start <= 0 || length <= 0 {
		return ""
	}
	start0 := start - 1
	if start0 >= len(source) {
		return ""
	}
	available := len(source) - start0
	actualLength := length
	if actualLength > available {
		actualLength = available
	}
	return source[start0 : start0+actualLength]
}

func (mm *MainMethod) GenerateBukkitUUID(username string) string {
	return mmcll.GenerateBukkitUUID(username)
}

func (mm *MainMethod) UUIDToAvatar(uuid string) int64 {
	chars := map[rune]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'a': "1010",
		'b': "1011",
		'c': "1100",
		'd': "1101",
		'e': "1110",
		'f': "1111",
	}
	regex, _ := regexp.Compile("^[a-z0-9]{32}$")
	if regex.MatchString(uuid) {
		bin := ""
		for _, char := range uuid {
			bin += chars[char]
		}
		most1 := bin[0:64]
		least1 := bin[64:128]
		xor1 := ""
		for index := 0; index < 64; index++ {
			if most1[index] == least1[index] {
				xor1 += "0"
			} else {
				xor1 += "1"
			}
		}
		most2 := xor1[0:32]
		least2 := xor1[32:64]
		xor2 := ""
		for index := 0; index < 32; index++ {
			if most2[index] == least2[index] {
				xor2 += "0"
			} else {
				xor2 += "1"
			}
		}
		var ten int64
		for index := 0; index < 32; index++ {
			if xor2[index] == '1' {
				ten += int64(math.Trunc(math.Pow(float64(len(xor2)-index), 2.0)))
			}
		}
		return ten
	}
	return -1
}

// GetBackgroundImage 获取一张附带索引的背景图片
func (mm *MainMethod) GetBackgroundImage(index int) []string {
	res := filepath.Join(GetCurrentExeDir(), "PCL.Nova", "BackgroundImage")
	if err := os.MkdirAll(res, fs.ModePerm); err != nil {
		panic(err)
	}
	files, err := os.ReadDir(res)
	if err != nil {
		panic(err)
	}
	var ind int
	if len(files) > 0 {
		if index < 0 || index > len(files)-1 {
			ind = rand.Intn(len(files))
		} else {
			ind = index
		}
	} else {
		return []string{}
	}
	i := 0
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		ext := strings.ToLower(filepath.Ext(fileName))
		if ext != ".png" && ext != ".jpg" {
			continue
		}
		if i == ind {
			fullPath := filepath.Join(res, fileName)
			fileData, err := os.ReadFile(fullPath)
			if err != nil {
				fmt.Printf("读取文件 %s 失败: %v\n", fileName, err)
				continue
			}
			base64Str := base64.StdEncoding.EncodeToString(fileData)
			return []string{
				base64Str,
				ext,
			}
		}
		i += 1
	}
	return []string{}
}

// GetTotalMemory 获取内存总量
func (mm *MainMethod) GetTotalMemory() uint64 {
	v, _ := mem.VirtualMemory()
	return v.Total / 1024 / 1024
}

// GetAvailableMemory 获取内存余量
func (mm *MainMethod) GetAvailableMemory() uint64 {
	v, _ := mem.VirtualMemory()
	return v.Available / 1024 / 1024
}

// GetJavaExecutableFileName 可以跨平台获取到 Java 的可执行文件，并且 O(1) 的复杂度，在前端想怎么调用就怎么调用。。
// 永远返回一个列表，第一个元素是 java 第二个元素是 javaw
func (mm *MainMethod) GetJavaExecutableFileName() []string {
	if runtime2.GOOS == "windows" {
		return []string{"java.exe", "javaw.exe"}
	} else {
		return []string{"java", "javaw"}
	}
}
