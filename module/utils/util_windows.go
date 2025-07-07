package utils

import (
	"NovaPlus/module/launcher"
	"golang.org/x/sys/windows/registry"
	"strconv"
	"strings"
)

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
	str3 := launcher.StrFill(strconv.FormatUint(launcher.GetHash(str+str2), 16), "7", 16)
	return launcher.Mid(str3, 5, 4) + "-" + launcher.Mid(str3, 13, 4) + "-" + launcher.Mid(str3, 1, 4) + "-" + launcher.Mid(str3, 9, 4)
}
