/*
postal 初始化所有的msger后会获得两个个方法

	Send	以协程向所有注册成功的平台发送消息
	Logout	最后处理退出，复用平台
*/
package postal

import (
	"fmt"
	"sync"
	"time"
)

type Msger interface {
	Init() bool
	Send(title, msg string) bool
	WaitTime() time.Duration
	Logout()
}

// psotal保存所有告警消息的客户端
type postal struct {
	msgers   []Msger
	badMsers []Msger
	done     chan struct{}
}

//TODO 保存登录的信息，然后下次直接加载

// NewPostal会根据当前给的配置信息去初始化每个告警客户端
// 如果满足Msger接口并且初始化成功才会添加到msgers
func NewPostal(configMsgers ...Msger) *postal {
	postal := &postal{msgers: make([]Msger, 0, 3), badMsers: make([]Msger, 0, 3), done: make(chan struct{}, 1)}
	for _, msger := range configMsgers {

		if msger.Init() {
			postal.msgers = append(postal.msgers, msger)
		} else {
			postal.badMsers = append(postal.badMsers, msger)
		}

	}
	return postal
}

// Send控制所有告警客户端发送告警信息
func (p *postal) Send(title, msg string) {

	var allTask sync.WaitGroup
	taskNums := len(p.msgers)
	allTask.Add(taskNums)

	var timeout time.Duration

	//检查是否完成
	go func(done chan<- struct{}, w *sync.WaitGroup) {

		//开始等待
		w.Wait()
		p.done <- struct{}{}
	}(p.done, &allTask)

	for _, msger := range p.msgers {
		timeout += msger.WaitTime()

		go func(m Msger, w *sync.WaitGroup) {
			m.Send(title, msg)
			w.Done()
		}(msger, &allTask)
	}

	// 等待完成
	select {
	case <-p.done:
		break
	case <-time.After(timeout):
		//发送超时要更新 done chan bool
		//原来的就交给GC回收
		p.done = make(chan struct{}, 1)
		fmt.Println("msger send msg timeout")
	}

}

// Logout 退出所有平台
func (p *postal) Logout() {
	for i := range p.msgers {
		p.msgers[i].Logout()
	}
}

// Shutdown会等待所有的告警信息发送完成后再退出
// timeout是最长等待告警信息发送完成的时间
// func (p *postal) Shutdown(timeout time.Duration) {

// 	//延时等待完成
// 	select {
// 	case <-p.done:
// 		break
// 	case <-time.After(timeout):
// 		if timeout == 0 {
// 			log.Info("msger force quit suceess")
// 		} else {
// 			log.Error("msger send msg Timeout")
// 		}
// 	}

// 	for _, msger := range p.msgers {
// 		msger.Shutdown()
// 	}
// }
