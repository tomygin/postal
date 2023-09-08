package main

import (
	"fmt"
	"time"

	"github.com/tomygin/postal"
)

func main() {

	//初始化发射站
	p := postal.NewPostal(
		&postal.Dida{
			Account:  "xxxx",
			Password: "xxxx",
		}, &postal.Xtuis{
			Token: "xxxx",
		},
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

	//查看发送情况
	fmt.Println(p.Status())

	//添加
	p.AddMsger(&postal.QQMail{
		SendAddr:    "xxx",
		ReceiveAddr: []string{"xxx"},
		AuthCode:    "xxx",
	})

}
