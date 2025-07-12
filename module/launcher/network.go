package launcher

import "NovaPlus/module/mmcll"

type Network struct{}

func (n *Network) HttpGet(url, authorization string) string {
	m := mmcll.NewUrlMethod(url)
	if authorization == "" {
		getDefault, err := m.GetDefault()
		return mmcll.If(err != nil, "", getDefault).(string)
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
