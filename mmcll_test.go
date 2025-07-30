package main

import (
	"NovaPlus/module/launcher"
	"NovaPlus/module/mmcll"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

// 测试 MMCLL 启动游戏~
func TestMMCLL(t *testing.T) {
	account := mmcll.NewLaunchAccountOffline("aooooo", "1234567980abcdef1234567890abcdef")
	gp := "D:/mc/testmcinheritsfrom/.minecraft/versions/1.21.1-52.1.1"
	options := mmcll.NewLaunchOption(account, "D:/Languages/Java/jdk-21.0.7/bin/java.exe", "D:/mc/testmcinheritsfrom/.minecraft", gp, gp)
	err := mmcll.LaunchGame(*options, true, func(back []string) {
		//for _, b := range back {
		//	fmt.Println(b)
		//}
		cmd := exec.Command(back[0], back[1:]...)
		stdout, _ := cmd.StdoutPipe()
		err := cmd.Start()
		if err != nil {
			t.Fatal(err)
			return
		}
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		err = cmd.Wait()
		if err != nil {
			t.Fatal(err)
			return
		}
	})
	if err != nil {
		t.Error(err)
	}
}

func TestSafe(t *testing.T) {
	j := `{
	"Hello": {
		"World": [
			{},
			"He",
			{
				"Goji": 7879
			}
		]
	}
}`
	var c map[string]any
	_ = json.Unmarshal([]byte(j), &c)
	op := mmcll.Safe(c, nil, "Hello", "World", 2, "Goji")
	fmt.Println(op) // 输出 nil
	op = mmcll.Safe(c, 5.21, "Hello", "World", 2, "Goji")
	fmt.Println(op) // 输出 7879
	op = mmcll.Safe(c, "I love you", "Hello", "World", 2, "Goji")
	fmt.Println(op) // 输出 I love you
	op = mmcll.Safe(c, []any{}, "Hello", "World")
	fmt.Println(op) // 输出 [map[] He map[Goji:7879]]
	op = mmcll.Safe(c, []any{}, "Hello", "Hello")
	fmt.Println(op)
}

// TestMicrosoftLogin 测试微软登录，测试通过！！
func TestMicrosoftLogin(t *testing.T) {
	// 下列字段将文件被忽略，你需要填入自己的 Client ID！
	clientId := launcher.ClientId()
	login := mmcll.NewAccountLogin(clientId)
	uc, dc, err := login.GetUserCode()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(uc)
	for {
		time.Sleep(5 * time.Second)
		s, err2 := login.LoginMicrosoft(dc)
		if err2 != nil {
			var err3 mmcll.ErrorMMCLL
			errors.As(err2, &err3)
			if err3.Code() != -6 {
				break
			}
		} else {
			fmt.Println(s.Name())
			fmt.Println(s.Uuid())
			fmt.Println(s.AccessToken())
			t.Log("Login success!")
			break
		}
	}
}
func TestDir(t *testing.T) {
	v := "/Workspace/GoWork"
	f1 := filepath.Dir(v)
	f2 := filepath.Dir(f1)
	f3 := filepath.Dir(f2)
	f4 := filepath.Dir(f3)
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	fmt.Println(f4)
}
func TestConvertName(t *testing.T) {
	s1 := "a.b.c.d:e:f:g@kkk"
	s2 := "a.b.c.d:e:f@kkk"
	s3 := "a.b.c.d:e:f:g"
	s4 := "a.b.c.d:e:f"
	fmt.Println(mmcll.ConvNameToPath(s1)) // a/b/c/d/e/f/e-f-g.kkk
	fmt.Println(mmcll.ConvNameToPath(s2)) // a/b/c/d/e/f/e-f.kkk
	fmt.Println(mmcll.ConvNameToPath(s3)) // a/b/c/d/e/f/e-f-g.jar
	fmt.Println(mmcll.ConvNameToPath(s4)) // a/b/c/d/e/f/e-f.jar
}

func TestOther(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	a = append(a[:5], a[7:]...)
	fmt.Println(a)
}

func TestIPv6(t *testing.T) {
	var sw launcher.MainMethod
	a := sw.GetAllIPv6()
	fmt.Println(a)
}
func TestCMD(t *testing.T) {
	var lm launcher.LaunchMethod
	java1 := lm.GetJavaInfo("D:/Languages/Java/jdk-21.0.8/bin/java.exe")
	t.Logf("%s\n%s\n%s\n%s", java1.Path, java1.Arch, java1.Vendor, java1.Version)
	java2 := lm.GetJavaInfo("D:/Languages/Java/jdk-17.0.12/bin/java.exe")
	t.Logf("%s\n%s\n%s\n%s", java2.Path, java2.Arch, java2.Vendor, java2.Version)
	java3 := lm.GetJavaInfo("D:/Languages/Java/jdk1.8.0_311/bin/java.exe")
	t.Logf("%s\n%s\n%s\n%s", java3.Path, java3.Arch, java3.Vendor, java3.Version)
}
func TestMemory(t *testing.T) {
	var mm launcher.MainMethod
	for _ = range 100 {
		fmt.Println(mm.GetTotalMemory())
		fmt.Println(mm.GetAvailableMemory())
		time.Sleep(100 * time.Millisecond)
	}
}
func TestLaunchGame(t *testing.T) {
	var lm launcher.LaunchMethod
	fmt.Println(lm.LaunchGame())
}
