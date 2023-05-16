

## 消息提醒

尽量用最少的依赖，实现的多平台消息提醒，一键推送，总有一个是你需要的😋

```go
package main

import "github.com/tomygin/postal"

func main() {
	// 注册推送平台
	p := postal.NewPostal(
		//大写的字段都要填写
		&postal.Xtuis{Token: "jTR6xWspfVyxxxasadsadHuisa"},
		&postal.Dida{Account: "xxx@outlook.com", Password: "xxx"},

		&postal.QQMail{SendAddr: "xxxx@foxmail.com",
			AuthCode:    "my smtp code",
			ReceiveAddr: []string{"xxxx@outlook.com", "xxxx@qq.com"}},
	)

	// 最后退出所有平台
	defer p.Logout()

	// 以协程向所有成功注册的平台发送消息
	p.Send("tomygin", "nice!")
	p.Send("第1波", "ok")

}


```

### 目前接入的平台

- [虾推啥](https://xtuis.cn/)
- [滴答清单](https://www.dida365.com/)
- [QQ邮箱](https://mail.qq.com/)
