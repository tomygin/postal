

## 消息提醒

尽量用最少的依赖，实现的多平台消息提醒，一键推送，总有一个是你需要的😋

```go
package main

import (
	"time"

	"github.com/tomygin/postal"
)

func main() {
	// 注册推送平台
	p := postal.NewPostal(
        //大写的字段都要填写
		&postal.Xtuis{Token: "your token"},
		&postal.Dida{Account: "xxx@outlook.com", Password: "xxx"})
	// 以协程向所有成功注册的平台发送消息
	// 发送完毕就退出单个send最大阻塞时间为平台数*1s
	p.Send("tomygin", "nice!")
	p.Send("第二波", "ok")
	p.Send("第三波", "ok")
	p.Send("第四波", "ok")


}


```

### 目前接入的平台

- [虾推啥](https://xtuis.cn/)
- [滴答清单](https://www.dida365.com/)