package main

import (
	"fmt"
	"time"

	"github.com/tomygin/postal"
)

func main() {

	t := time.Now()
	//初始化发射站
	p := postal.NewPostal(
		&postal.Dida{
			Account:  "xxx",
			Password: "xxx",
		}, &postal.Xtuis{
			Token: "jTtG51Dicl",
		},
	)

	//起草一个信息，如果成功起草将结束之前的发射
	//如不指定，默认5秒超时
	s := p.Draft("tomygin", "welcome", 15*time.Second)
	fmt.Println("草稿完成", time.Since(t))
	//发射
	go s.Send()

	//手动控制取消
	go func() {
		time.Sleep(10 * time.Second)
		s.Cancel()
	}()

	//等待完成或者取消，否则会阻塞在这里
	s.Wait()

	//查看发送情况
	fmt.Println(p.Status())

	//添加
	p.AddMsger(&postal.QQMail{
		SendAddr:    "xxx",
		ReceiveAddr: []string{"xxx"},
		AuthCode:    "xxx",
	})

	//退出,传入true就强制退出之前的发送，false就等待后再退出
	s.SignOut(true)
}
