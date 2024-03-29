package postal

import (
	"fmt"
	"net/http"
)

type Xtuis struct {
	_     [0]int //让初始化的时候必须指明字段
	Token string
	url   string
}

func (x *Xtuis) Init() bool {
	x.url = "http://wx.xtuis.cn/" + x.Token + ".send?text=%s" + "&desp=%s"
	return true
}

func (x *Xtuis) Msg(title, msg string) chan struct{} {
	url := fmt.Sprintf(x.url, title, msg)
	http.Get(url)
	c := make(chan struct{}, 1)
	c <- struct{}{}
	return c
}

var _ Msger = (*Xtuis)(nil)
