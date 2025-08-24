package launcher

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"time"
)

func GetMachineCode() string {
	return "DNS"
}

func GetHomeDir() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(currentUser.HomeDir, ".PCL.Nova"), nil
}
func PingCMD(ip string, timeout time.Duration) *exec.Cmd {
	return CMD("ping", "-c", "1", "-W", fmt.Sprintf("%.0f", timeout.Seconds()), ip)
}

func CMD(name string, args ...string) *exec.Cmd {
	return exec.Command(name, args...)
}

func openExp(fpath string) error {
	// 尝试检测桌面环境并打开文件管理器
	desktopEnv := os.Getenv("XDG_CURRENT_DESKTOP")
	desktopEnv = strings.ToLower(desktopEnv)

	// 根据桌面环境选择相应的文件管理器
	switch {
	case strings.Contains(desktopEnv, "gnome"):
		return exec.Command("nautilus", path).Start()
	case strings.Contains(desktopEnv, "kde"):
		return exec.Command("dolphin", path).Start()
	case strings.Contains(desktopEnv, "xfce"):
		return exec.Command("thunar", path).Start()
	case strings.Contains(desktopEnv, "mate"):
		return exec.Command("caja", path).Start()
	case strings.Contains(desktopEnv, "lxde"):
		return exec.Command("pcmanfm", path).Start()
	case strings.Contains(desktopEnv, "cinnamon"):
		return exec.Command("nemo", path).Start()
	case strings.Contains(desktopEnv, "pantheon"):
		return exec.Command("io.elementary.files", path).Start()
	case strings.Contains(desktopEnv, "budgie"):
		return exec.Command("nautilus", path).Start()
	default:
		// 尝试使用xdg-open作为后备方案
		if err := exec.Command("xdg-open", path).Start(); err != nil {
			return fmt.Errorf("无法识别桌面环境或缺少文件管理器: %s", desktopEnv)
		}
		return nil
	}
}
