package launcher

import (
	"os/user"
	"path/filepath"
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
	cmd := exec.Command("ping", "-c", "1", "-W", fmt.Sprintf("%d", timeout.Seconds()), ip)
	return cmd
}
