package launcher

import (
	"NovaPlus/module/mmcll"
	"NovaPlus/module/mmcll_utils"
	"encoding/base64"
	"fmt"
	"io/fs"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type ReaderWriter struct{}
type MainMethod struct{}

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
	home, err := mmcll_utils.GetHomeDir()
	if err != nil {
		exePath, err := os.Executable()
		if err != nil {
			panic(err)
		}
		home = filepath.Join(filepath.Dir(exePath), "PCL.Nova", "config")
	}
	res := filepath.Join(home, "Other.ini")
	err = EnsureConfigFile(res)
	if err != nil {
		panic(err)
	}
	return res
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
func (rw *ReaderWriter) GetConfigIniPath() string {
	res := filepath.Join(GetCurrentExeDir(), "PCL.Nova", "config", "PCL2.Nova.ini")
	err := EnsureConfigFile(res)
	if err != nil {
		panic(err)
	}
	return res
}

func (mm *MainMethod) GenerateBukkitUUID(username string) string {
	return mmcll.GenerateBukkitUUID(username)
}

func (mm *MainMethod) UUIDToAvatar(uuid string) int64 {
	regex, _ := regexp.Compile("^[a-z0-9]{32}$")
	if regex.MatchString(uuid) {
		bin := ""
		for _, char := range uuid {
			switch char {
			case '0':
				bin += "0000"
				break
			case '1':
				bin += "0001"
				break
			case '2':
				bin += "0010"
				break
			case '3':
				bin += "0011"
				break
			case '4':
				bin += "0100"
				break
			case '5':
				bin += "0101"
				break
			case '6':
				bin += "0110"
				break
			case '7':
				bin += "0111"
				break
			case '8':
				bin += "1000"
				break
			case '9':
				bin += "1001"
				break
			case 'a':
				bin += "1010"
				break
			case 'b':
				bin += "1011"
				break
			case 'c':
				bin += "1100"
				break
			case 'd':
				bin += "1101"
				break
			case 'e':
				bin += "1110"
				break
			case 'f':
				bin += "1111"
				break
			default:
				panic("[Invalid Input!]")
			}
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
		// 跳过目录，只处理文件
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		ext := strings.ToLower(filepath.Ext(fileName)) // 获取小写扩展名

		// 检查是否为图片文件
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
