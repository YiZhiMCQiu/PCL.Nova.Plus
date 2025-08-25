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
	return PingCMD(host, timeout).Run()
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
func (mm *MainMethod) GetAllIPv6() ExceptionHandler[[]IPv6Struct] {
	s, err := net.InterfaceAddrs()
	if err != nil {
		return NewExceptionHandler(500, false, err.Error(), make([]IPv6Struct, 0))
	}
	result := make([]IPv6Struct, 0)
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
	return NewExceptionHandler(200, true, "Ok!", result)
}
func (rw *ReaderWriter) WriteConfig(path, section, key, value string) {
	conf := NewConfig(path)
	_ = conf.Write(section, key, value)
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
	return mmcll.If(err != nil, "", res).(string)
}
func (rw *ReaderWriter) GetConfigIniPath() string {
	res := filepath.Join(GetCurrentExeDir(), "PCL.Nova", "config", "PCL2.Nova.ini")
	err := EnsureConfigFile(res)
	return mmcll.If(err != nil, "", res).(string)
}
func (rw *ReaderWriter) GetCurrentExeDir() string {
	return GetCurrentExeDir()
}
func (rw *ReaderWriter) OpenDirectoryDialog(title string) string {
	dialog, err := runtime.OpenDirectoryDialog(rw.Ctx, runtime.OpenDialogOptions{
		Title:                title,
		CanCreateDirectories: true,
	})
	return mmcll.If(err != nil, "", dialog).(string)
}
func (rw *ReaderWriter) OpenFileDialog(title string, filter []string) string {
	var fileFilter []runtime.FileFilter
	for _, f := range filter {
		fileFilter = append(fileFilter, runtime.FileFilter{
			DisplayName: f,
			Pattern:     f,
		})
	}
	if len(filter) == 0 {
		dialog, err := runtime.OpenFileDialog(rw.Ctx, runtime.OpenDialogOptions{
			Title:                title,
			CanCreateDirectories: true,
		})
		return mmcll.If(err != nil, "", dialog).(string)
	} else {
		dialog, err := runtime.OpenFileDialog(rw.Ctx, runtime.OpenDialogOptions{
			Title:                title,
			CanCreateDirectories: true,
			Filters:              fileFilter,
		})
		return mmcll.If(err != nil, "", dialog).(string)
	}
}
func (rw *ReaderWriter) OpenExplorer(fpath string) bool {
	err := openExp(fpath)
	return err == nil
}

func GetCurrentExeDir() string {
	exePath, err := os.Executable()
	return mmcll.If(err != nil, "", filepath.Dir(exePath)).(string)
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
		return []string{}
	}
	files, err := os.ReadDir(res)
	if err != nil {
		return []string{}
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
		return []string{}
	}
}

type HomePageStruct struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (mm *MainMethod) ReadFile(path string) string {
	data, err := mmcll.GetFile(path)
	return mmcll.If(err != nil, "", data).(string)
}
func (mm *MainMethod) GetAllHomePage() ExceptionHandler[[]HomePageStruct] {
	res := make([]HomePageStruct, 0)
	dir := filepath.Join(GetCurrentExeDir(), "PCL.Nova", "HomePage")
	if err := os.MkdirAll(dir, fs.ModePerm); err != nil {
		return NewExceptionHandler(400, false, "File System Error: "+err.Error(), []HomePageStruct{})
	}
	files, err := os.ReadDir(dir)
	if err != nil {
		return NewExceptionHandler(400, false, "File System Error: "+err.Error(), []HomePageStruct{})
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		name := file.Name()
		if !strings.HasSuffix(name, ".nxml") {
			continue
		}
		res = append(res, HomePageStruct{
			Path: filepath.Join(dir, name),
			Name: name,
		})
	}
	return NewExceptionHandler(200, true, "Ok!", res)
}
func (mm *MainMethod) GenerateTutorialHomePage() {
	HPPath := filepath.Join(GetCurrentExeDir(), "PCL.Nova", "HomePage", "Custom.nxml")
	err := EnsureConfigFile(HPPath)
	if err != nil {
		return
	}
	_ = mmcll.SetFile(HPPath, `<MyCard title="有标题，但是折叠" isExpand="True" canExpand="True">
    <MySpan>有标题，但是已折叠，所有的 isExpand 和 canExpand 属性均默认为 False，因此如果你想折叠卡片，你需要显式声明两个属性！</MySpan>
</MyCard>
<MyCard title="有标题，但是已展开" isExpand="False" canExpand="True">
    <MySpan>有标题，但是已展开</MySpan>
</MyCard>
<MyCard title="有标题，但是无法折叠1" isExpand="True" canExpand="False">
    <MySpan>有标题，但是无法折叠1</MySpan>
</MyCard>
<MyCard title="有标题，但是无法折叠2" isExpand="False" canExpand="False">
    <MySpan>有标题，但是无法折叠2，当 canExpand 设置为 False 的时候，将会自动忽略 isExpand 属性，永久性无法被折叠。</MySpan>
</MyCard>
<MyCard title="" isExpand="False" canExpand="False">
    <MySpan>无标题，无法折叠，会忽略 isExpand 和 canExpand 两个属性</MySpan>
</MyCard>

<MyCard title="MyDiv 测试用例" isExpand="True" canExpand="True">
    <MyDiv style="display: flex; flex-direction: column;">
        <MySpan>当设计了 父组件 MyDiv 的 display: flex; flex-direction: column 之后，下面所有的 MySpan 均为垂直排列。</MySpan>
        <MySpan>颜色为<MySpan style="color: white;">纯白色</MySpan>，但是却很大<MySpan style="font-size: 50px;">50px</MySpan>的文字~</MySpan>
        <MySpan>请记住，无论你使不使用 MySpan 进行布局书写，只要你没单独为 MySpan 设置 color 的 style 属性，文字均会按照【暗色模式】的样式写字体模式。</MySpan>
    </MyDiv>
</MyCard>
<MyCard title="MyButton 测试用例 与 类 PCL Grid 布局~" isExpand="True" canExpand="True">
    <MyDiv style="width: 100%; height: 60px"><MyButton width="100%" height="60px">MyButton 是行内样式，默认不会换行。你可以使用空的 MyDiv
        作为父样式来换行（自定义主页不支持 br 标签）</MyButton></MyDiv>
    <MySpan>这是高级样式，父组件设置成这种布局之后，里面所有按钮都会扩宽等距排列。你也可以设置 justify-content 属性为 space-around，这样会更好看。请注意，你一定要使用 width: 100%，否则可能会导致扩宽不成功（</MySpan>
    <MyDiv style="width: 100%; display: flex; align-items: center; justify-content: space-between">
        <MyButton width="30%" height="50px">等距排列按钮1</MyButton>
        <MyButton width="30%" height="50px">等距排列按钮2</MyButton>
        <MyButton width="30%" height="50px">等距排列按钮3</MyButton>
    </MyDiv>
</MyCard>
<MyCard title="MyDiv 高级样式之自制仿 PCL2 的 MyHint" isExpand="True" canExpand="True">
    <MyDiv style="border-radius: 4px; font-size: 14px; padding: 8px 14px; border-left: 3px solid rgb(216, 41, 41); background-color: rgb(255, 221, 223); color: rgb(216, 41, 41);">这也是一个模仿 PCL2 的 MyHint 的控件，但是颜色为红色~</MyDiv>
    <MyDiv style="border-radius: 4px; font-size: 14px; padding: 8px 14px; border-left: 3px solid rgb(245, 122, 0); background-color: rgb(255, 235, 215); color: rgb(245, 122, 0);">请记住，在这里面的所有 MyDiv 控件布局，一定要加 margin-top，不然所有控件杂糅在一起就会很难看（默认的 margin-top 是 5px，你可以往上加很多。</MyDiv>
    <MyDiv style="border-radius: 4px; font-size: 14px; padding: 8px 14px; border-left: 3px solid rgb(17, 114, 212); background-color: rgb(217, 236, 255); color: rgb(17, 114, 212);">具体所有的 CSS 值，你们完全可以去 菜鸟教程 官网随意查阅嗷~</MyDiv>
    <MyDiv style="position: relative; height: 50px;">
        <MyDiv style="position: absolute; right: 0; width: 100px; height: 50px; background-color: orange; color: black;">这是文字处在右边~</MyDiv>
    </MyDiv>
</MyCard>
<MyCard title="MySpan 高级样式" isExpand="True" canExpand="True">
    <MyDiv style="margin: 10px 0">
        下列是一堆 MySpan 的示例！MySpan 是行内样式，是可以嵌入到任何<MySpan style="font-size: 50px;">50px</MySpan>的文字里面的！
    </MyDiv>
    <MySpan style="user-select: text">这串文字是可以被选择的喵~</MySpan>
    <MySpan style="font-weight: bold">加粗的文字喵~</MySpan>
    <MySpan style="font-style: oblique">斜体的文字喵~</MySpan>
    <MySpan style="text-decoration: line-through #f00">删除的文字喵，线也是红色的喵~</MySpan>
    <MySpan style="color: #00ffff">这是水蓝色的文字喵~</MySpan>
    甚至还有
    <MySpan style="background-color: lightgray; border: 1px solid gray; border-radius: 3px; font-size: 16px; vertical-align: middle; padding: 1px 4px; color: black">code</MySpan>
    代码块捏~
</MyCard>
<MyCard title="MyButton 支持的所有样式" isExpand="True" canExpand="True">
    <MyDiv>这里介绍一下 MyButton 的几种目前支持的样式</MyDiv>
    <MyDiv>首先是正常的 Button，这个 Button 是一个行内元素，你可以在文字之间嵌入 Button，就像以下：</MyDiv>
    <MyDiv>在文字之间<MyButton>嵌入一个按钮</MyButton>捏~</MyDiv>
    <MyDiv>按钮支持的属性有：width、height、color、event、type，其中 event 属性在下一个卡片会讲解。type 总共有 2 个值。如果什么也不填的话，就是 default。</MyDiv>
    <MyDiv>以下按钮为 设置了 color 的按钮。color 仅针对边框，如果你需要同时连文字一同设置颜色，你需要在里面再设置一层 MySpan（或者 MyDiv 都行。。）</MyDiv>
    <MyDiv>
        <MyButton type="default" width="200px" height="100px" color="#FF0000">
            <MyDiv style="font-size: 14px;">没有人规定不能在 MyButton 里面写标签吧~</MyDiv>
            <MySpan style="color: red; font-size: 14px;">严重警告的按钮（红色）</MySpan>
        </MyButton>
    </MyDiv>
    <MyDiv>你可能会问，为什么按钮不支持 style，原因其实很简单，因为按钮是我自制组件的按钮，而上述 MySpan 和 MyDiv 都是 HTML 自带的！你一旦修改了 style，很可能会破坏按钮原本的 css 样式！</MyDiv>
    <MyDiv>我们还有一种非常特殊的按钮，那就是【文本按钮】！其中，如果将 MyButton 的 type 属性填成了【label】，就说明这个按钮是文本按钮！文本按钮有以下属性：style、hov-style、event</MyDiv>
    <MyDiv style="display: flex; flex-direction: column; align-items: center">
        <MyButton type="label" hov-style="color: skyblue;">这是一个文本按钮</MyButton>
        <MyButton type="label" style="border-radius: 10px; padding: 4px 6px;" hov-style="border-radius: 10px; padding: 4px 6px; background-color: rgba(0, 0, 255, 0.2); color: #0000FF">类似 PCL 的 MyTextButton 那样的按钮！</MyButton>
    </MyDiv>
    <MyDiv>好了，在 type 不一致时，你们不允许设置一些别的那种属性，就好比如说当 type 为 default 或者不填时，无法拥有 style 属性，而 type 为 label 时，就有 2 个新增的属性啦！</MyDiv>
</MyCard>
<MyCard title="图标与图片讲解" isExpand="True" canExpand="True">
    <MyDiv>接下来是图标详解！我们可以通过声明一个新的元素：MySvg 去使用我们的 SVG 矢量图标~具体参见以下示例：</MyDiv>
    <MyDiv>首先，MySvg 的样式与其他的可能会稍微有点不同，他自动声明了 role="img" 以及 xmlns="http://www.w3.org/2000/svg"，你需要做的只是声明它的 stroke、icon 等一系列的属性即可！</MyDiv>
    <MyDiv>哦对了！还有一点！MySvg 是一个单标签，你无需声明它的子标签！而在这个 Svg，你需要填入一个必填项，那就是 path！这个相当于矢量图里的 path 路径！类似通过 path d="xxx" 去绘制矢量图！</MyDiv>
    <MyDiv style="display: flex; align-items: center; justify-content: center;">
        <MySvg viewbox="0 0 24 24" style="width: 32px; height: 32px; fill: red;" path="M12 17q.425 0 .713-.288T13 16v-4q0-.425-.288-.712T12 11t-.712.288T11 12v4q0 .425.288.713T12 17m0-8q.425 0 .713-.288T13 8t-.288-.712T12 7t-.712.288T11 8t.288.713T12 9m0 13q-2.075 0-3.9-.788t-3.175-2.137T2.788 15.9T2 12t.788-3.9t2.137-3.175T8.1 2.788T12 2t3.9.788t3.175 2.137T21.213 8.1T22 12t-.788 3.9t-2.137 3.175t-3.175 2.138T12 22m0-2q3.35 0 5.675-2.325T20 12t-2.325-5.675T12 4T6.325 6.325T4 12t2.325 5.675T12 20m0-8" />
    </MyDiv>
    <MyDiv>上述，我画了一个 圆圈，中间是一个 i 字符的一个图标~</MyDiv>
    <MyDiv>与此同时，我们不仅可以用 fill 去绘制，我们还可以用 stroke 去绘制图标！例如以下：</MyDiv>
    <MyDiv style="display: flex; align-items: center; justify-content: center;">
        <MySvg viewbox="0 0 16 16" style="width: 32px; height: 32px; stroke: red; stroke-width: 2px; stroke-linecap: round; stroke-linejoin: round; fill: none;"
               path="M3 15.5h3a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 .5.5v4a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 .5-.5V8.5h1.453a.497.497 0 0 0 .404-.836L8.354.654a.5.5 0 0 0-.708 0L.643 7.664a.497.497 0 0 0 .404.836H2.5V15a.5.5 0 0 0 .5.5z" />
    </MyDiv>
    <MyDiv>由于 MySvg 不支持 hov-style 事件，因此你暂时无法对其进行 鼠标悬浮上去的提示（也许将来会支持呢！</MyDiv>
    <MyDiv>好，接下来说一下图片吧！图片其实非常简单！只需要输入 MyImg 标签即可！里面有三个属性，分别是：src、alt、style。。是的！这是整个 Nova 控件库里面第四个支持 style 的控件！请看以下示例！</MyDiv>
    <MyDiv>例如，加载一个苹果~</MyDiv>
    <MyDiv style="display: flex; align-items: center; justify-content: center;">
        <MyImg src="https://q8.itc.cn/images01/20250816/0626eec82e4d4c10800940ba9c3954bb.jpeg" alt="苹果" style="width: 100%;" />
    </MyDiv>
    <MyDiv>哦对了，无论是 MyImg 还是 MySvg 都是行内元素，各位可以随时的插入到 MyButton 里面！</MyDiv>
</MyCard>
<MyCard title="MyButton 高级使用案例" isExpand="True" canExpand="True">
    <MyDiv>这一张卡片，我将教会各位如何使用 MyButton 写出那种类似 PCL 的 MyListItem！在我们学习了上述 MyImg 和 MySvg 以及 MySpan 之后，各位应该对于盒子布局有所了解了！请看下面示例！</MyDiv>
    <MyDiv style="width: 100%; height: 50px;">
        <MyButton type="label" style="display: block; width: 100%; height: 40px; border-radius: 10px;" hov-style="display: block; width: 100%; height: 40px; border-radius: 10px; cursor: pointer; background-color: rgba(0, 0, 255, 0.2);">
            <MyDiv style="display: flex; align-items: center; width: 100%; height: 40px;">
                <MyImg src="https://www.baidu.com/favicon.ico" alt="百度" style="margin-left: 10px; height: 30px; width: 30px" />
                <MyDiv style="flex: 1; margin-left: 10px; margin-top: 0">
                    <MyDiv style="font-size: 14px; margin-top: 0;">这是 MyListItem 的主标题</MyDiv>
                    <MyDiv style="font-size: 10px; color: gray; margin-top: 0;">这是 MyListItem 的副标题</MyDiv>
                </MyDiv>
            </MyDiv>
        </MyButton>
    </MyDiv>
    <MyDiv>我知道上述代码可能会有点小多，不过如果你但凡学过一点点的前端知识，应该能懂上面写了什么，不过不懂也没有关系~完全可以 copy 上述的写法~</MyDiv>
</MyCard>
<MyCard title="NovaPlus 的替换标记与内置图片~" isExpand="True" canExpand="True">
    <MyDiv>在 Nova Plus 里，有着很多的内置图片与替换标记，可以参考以下示例：</MyDiv>
    <MyDiv>首先，说的便是我们的 path 标记啦！和 PCL 一致，只需要用花括号括起来，就知道了替换标记是什么了~就像这样：（{path}），可以用于执行某个事件时需要用到！</MyDiv>
    <MyDiv>目前 Nova Plus 支持以下 几种 替换标记，分别是：</MyDiv>
    <MyDiv style="display: table; width: 100%;">
        <MyDiv style="display: table-header-group">
            <MyDiv style="display: table-row; height: 50px;">
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; font-weight: bold; vertical-align: middle;">替换标记</MyDiv>
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; font-weight: bold; vertical-align: middle;">实际内容</MyDiv>
            </MyDiv>
        </MyDiv>
        <MyDiv style="display: table-row-group">
            <MyDiv style="display: table-row; height: 50px;">
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">path</MyDiv>
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">{path}</MyDiv>
            </MyDiv>
        </MyDiv>
    </MyDiv>
    <MyDiv>啊哈！又让你们学到了一招！表格！！表格见上方呀~</MyDiv>
    <MyDiv>下面是 内置图片~内置图片见下方：</MyDiv>
    <MyDiv>
        <MyImg src="CommandBlock" style="width: 32px; height: 32px; margin-left: 5px;" />
        <MyImg src="Cobblestone" style="width: 32px; height: 32px; margin-left: 5px;" />
        <MyImg src="GoldBlock" style="width: 32px; height: 32px; margin-left: 5px;" />
        <MyImg src="Grass" style="width: 32px; height: 32px; margin-left: 5px;" />
        <MyImg src="GrassPath" style="width: 32px; height: 32px; margin-left: 5px;" />
        <MyImg src="Anvil" style="width: 32px; height: 32px; margin-left: 5px;" />
        <MyImg src="RedstoneBlock" style="width: 32px; height: 32px; margin-left: 5px;" />
        <MyImg src="RedstoneLampOn" style="width: 32px; height: 32px; margin-left: 5px;" />
        <MyImg src="RedstoneLampOff" style="width: 32px; height: 32px; margin-left: 5px;" />
        <MyImg src="Egg" style="width: 32px; height: 32px; margin-left: 5px;" />
        <MyImg src="Fabric" style="width: 32px; height: 32px; margin-left: 5px;" />
        <MyImg src="NeoForge" style="width: 32px; height: 32px; margin-left: 5px;" />
    </MyDiv>
    <MyDiv>因此，各位可以使用上述图片来随时制作自己的图片按钮~就像以下：</MyDiv>
    <MyDiv style="width: 100%; height: 50px;">
        <MyButton type="label" style="display: block; width: 100%; height: 40px; border-radius: 10px;" hov-style="display: block; width: 100%; height: 40px; border-radius: 10px; cursor: pointer; background-color: rgba(0, 0, 255, 0.2);">
            <MyDiv style="display: flex; align-items: center; width: 100%; height: 40px;">
                <MyImg src="CommandBlock" alt="百度" style="margin-left: 10px; height: 32px; width: 32px" />
                <MyDiv style="flex: 1; margin-left: 10px; margin-top: 0">
                    <MyDiv style="font-size: 14px; margin-top: 0;">下载测试版</MyDiv>
                    <MyDiv style="font-size: 10px; color: gray; margin-top: 0;">23w07a</MyDiv>
                </MyDiv>
            </MyDiv>
        </MyButton>
    </MyDiv>
</MyCard>
<MyCard title="MyButton 的所有事件详解" isExpand="True" canExpand="True">
    <MyDiv>所有的事件以【event名|event参数1|event参数2|...】的样式进行~如果某一个事件没有参数，则不需要加【|】线，末尾同样也不需要|线</MyDiv>
    <MyDiv style="display: table; width: 100%;">
        <MyDiv style="display: table-header-group">
            <MyDiv style="display: table-row; height: 50px;">
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; font-weight: bold; vertical-align: middle;">事件名</MyDiv>
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; font-weight: bold; vertical-align: middle;">功能</MyDiv>
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; font-weight: bold; vertical-align: middle;">参数1</MyDiv>
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; font-weight: bold; vertical-align: middle;">参数2</MyDiv>
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; font-weight: bold; vertical-align: middle;">参数3</MyDiv>
            </MyDiv>
        </MyDiv>
        <MyDiv style="display: table-row-group">
            <MyDiv style="display: table-row; height: 50px;">
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">refresh</MyDiv>
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">刷新主页</MyDiv>
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">\</MyDiv>
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">\</MyDiv>
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">\</MyDiv>
            </MyDiv>
        </MyDiv>
        <MyDiv style="display: table-row-group">
            <MyDiv style="display: table-row; height: 50px;">
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">messagebox</MyDiv>
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">信息框</MyDiv>
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">标题</MyDiv>
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">内容</MyDiv>
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">颜色【0：蓝色、1：黄色、2：红色】</MyDiv>
            </MyDiv>
        </MyDiv>
        <MyDiv style="display: table-row-group">
            <MyDiv style="display: table-row; height: 50px;">
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">openurl</MyDiv>
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">打开网页</MyDiv>
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">url地址</MyDiv>
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">\</MyDiv>
                <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">\</MyDiv>
            </MyDiv>
        </MyDiv>
    </MyDiv>
    <MyDiv>举个例子~</MyDiv>
    <MyButton width="100px" height="50px" event="messagebox|这是标题|这是警告框的内容|1">触发警告信息框~</MyButton>
    <MyButton width="100px" height="50px" event="openurl|https://github.com/xphost008">打开作者的github~</MyButton>
</MyCard>`)
}
