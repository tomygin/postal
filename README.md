

## 消息提醒

尽量用最少的依赖，实现的多平台消息提醒，一键推送，总有一个是你需要的😋

```go
package main

import (
	"time"

	"github.com/tomygin/postal"
)

func main() {

	//初始化发射站
    //大写的字段都必须填写
	p := postal.NewPostal(
		&postal.Xtuis{
			Token: "token",
		},
		&postal.QQMail{
			SendAddr:    "xxxx@xx.com",
			ReceiveAddr: []string{"xxxx@xx.com"},
			AuthCode:    "codexxx"},
	)

	//起草一个信息，如果成功起草将结束之前的发射
	//最大超时默认5秒
	s := p.Draft("tomygin", "welcome", 10*time.Second)
	//发射
	go s.Send()

	//手动控制取消
	go func() {
		time.Sleep(7 * time.Second)
		s.Cancel()
	}()

	//等待取消，否则会阻塞在这里
	s.Wait()
}

```

### 目前接入的平台

- [虾推啥](https://xtuis.cn/)
- [滴答清单](https://www.dida365.com/)
- [QQ邮箱](https://mail.qq.com/)

### 必要信息
- 虾推啥，滴答清单使用的是80端口，所以尽量不要服务也使用80，容易端口冲突