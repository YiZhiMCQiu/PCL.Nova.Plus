package launcher

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func GetHomeDir() (string, error) {
	// 获取当前用户
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(currentUser.HomeDir, "AppData", "Roaming", "PCL.Nova"), nil
}

func GetMachineCode() string {
	key, _ := registry.OpenKey(registry.LOCAL_MACHINE, "SYSTEM\\HardwareConfig", registry.READ)
	defer key.Close()
	str, _, err := key.GetStringValue("LastConfig")
	if err != nil {
		return "ERR_1"
	}
	str = strings.Trim(strings.ToUpper(str), "{}")
	key2, _ := registry.OpenKey(registry.CURRENT_USER, "Software\\PCL", registry.READ)
	defer key2.Close()
	str2, _, err := key2.GetStringValue("Identify")
	if err != nil {
		return "ERR_2"
	}
	str3 := StrFill(strconv.FormatUint(GetHash(str+str2), 16), "7", 16)
	return Mid(str3, 5, 4) + "-" + Mid(str3, 13, 4) + "-" + Mid(str3, 1, 4) + "-" + Mid(str3, 9, 4)
}

func PingCMD(ip string, timeout time.Duration) *exec.Cmd {
	cmd := exec.Command("ping", "-n", "1", "-w", fmt.Sprintf("%d", timeout.Milliseconds()), ip)
	return cmd
}
