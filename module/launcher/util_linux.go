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
