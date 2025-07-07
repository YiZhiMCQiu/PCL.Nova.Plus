package main

import (
	"NovaPlus/module/mmcll"
	"fmt"
	"strings"
	"testing"
)

func TestMMCLL(t *testing.T) {
	account := mmcll.NewLaunchAccountOffline("aooooo", "1234567980abcdef1234567890abcdef")
	options := *mmcll.NewLaunchOption(account, "D:/Languages/Java/jdk-21.0.7/bin/java.exe", "D:/Workspace/GoWork/", "D:/Workspace/GoWork/", "D:/Workspace/GoWork/")
	err := mmcll.LaunchGame(options, true, func(back []string) {
		fmt.Println(strings.Join(back, "\n"))
	})
	if err != nil {
		t.Error(err)
	}
}

func TestMainMethod(t *testing.T) {
	a := App{}
	fmt.Println(a.GetTodayLucky("2E97-2E0C-4038-BE81"))
}
