package game

import (
	lua "github.com/yuin/gopher-lua"
)

func init() {
	L := lua.NewState()
	defer L.Close()
	L.OpenLibs()
}
