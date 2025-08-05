package launcher

import (
	"fmt"
	"os/exec"
	"os/user"
	"path/filepath"
	"time"
)

func GetMachineCode() string {
	return "DNS"
}

func GetHomeDir() (string, error) {
	// 获取当前用户
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(currentUser.HomeDir, ".PCL.Nova"), nil
}

func PingCMD(ip string, timeout time.Duration) *exec.Cmd {
	return exec.Command("ping", "-c", "1", "-W", fmt.Sprintf("%d", timeout.Seconds()), ip)
}
