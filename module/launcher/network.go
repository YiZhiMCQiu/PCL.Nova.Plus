package launcher

import (
	"NovaPlus/module/mmcll"
	"context"
)

type Network struct {
	Ctx context.Context
}

func (n *Network) HttpGet(url, authorization string) string {
	m := mmcll.NewUrlMethod(url)
	if authorization == "" {
		getDefault, err := m.GetDefault()
		return mmcll.If(err != nil, "", string(getDefault)).(string)
	} else {
		get, err := m.Get(authorization)
		return mmcll.If(err != nil, "", get).(string)
	}
}

func (n *Network) HttpPost(url, key string, isJSON bool) string {
	m := mmcll.NewUrlMethod(url)
	post, err := m.Post(key, isJSON)
	return mmcll.If(err != nil, "", post).(string)
}

func (n *Network) HttpCurseForge(url string) string {
	m := mmcll.NewUrlMethod(url)
	getCurseForge, err := m.GetCurseForge(CF_KEY)
	return mmcll.If(err != nil, "", getCurseForge).(string)
}
