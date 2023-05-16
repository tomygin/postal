package postal

import (
	"fmt"
	"net/http"
	"time"
)

type Xtuis struct {
	Token string
	url   string
}

func (x *Xtuis) Init() bool {
	x.url = "http://wx.xtuis.cn/" + x.Token + ".send?text=%s" + "&desp=%s"
	return true
}

func (x *Xtuis) Send(title, msg string) bool {
	url := fmt.Sprintf(x.url, title, msg)
	resp, ok := http.Get(url)
	return resp.StatusCode == 200 && ok == nil
}
func (x *Xtuis) WaitTime() time.Duration {
	return time.Duration(500 * time.Millisecond)
}
func (x *Xtuis) Logout() {}

var _ Msger = (*Xtuis)(nil)
