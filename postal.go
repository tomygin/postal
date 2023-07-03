/*
 */
package postal

import (
	"context"
	"time"
)

// Sender是用户调用接口
// Sender内部的Send方法用于统一发送
// Wait方法就会阻塞等待直到超时或者发送完成
type Sender interface {
	Send()
	Wait()
	Cancel()
}

// Msger是所有发送消息的平台应该满足的接口
type Msger interface {
	Init() bool
	Msg(title, msg string) chan struct{}
}

// psotal保存所有消息平台的客户端
type postal struct {
	msgers   []Msger
	badMsers []Msger
	ctx      context.Context
	cancel   context.CancelFunc

	title string
	msg   string
}

func NewPostal(msgers ...Msger) *postal {
	status := make([]chan bool, len(msgers))
	p := new(postal)
	for num, msger := range msgers {
		status[num] = make(chan bool, 1)
		go func(c chan bool, m Msger) {
			c <- m.Init()
		}(status[num], msger)
	}

	for num, msger := range msgers {
		if ok := <-status[num]; ok {
			p.msgers = append(p.msgers, msger)
		} else {
			p.badMsers = append(p.badMsers, msger)
		}
	}
	return p
}

func (p *postal) Send() {
	for num := range p.msgers {
		go func(p *postal, num int) {
			select {
			case <-p.ctx.Done():
			case <-p.msgers[num].Msg(p.title, p.msg):
			}
		}(p, num)
	}

}

func (p *postal) Wait() {
	<-p.ctx.Done()
}

func (p *postal) Cancel() {
	p.cancel()
}

func (p *postal) Draft(title, msg string, timeOut ...time.Duration) Sender {
	if timeOut == nil {
		timeOut = append(timeOut, 5*time.Second)
	}

	//结束之前的发送任务
	if p.cancel != nil {
		p.cancel()
	}

	p.ctx, p.cancel = context.WithTimeout(context.Background(), timeOut[0])

	p.msg = msg
	p.title = title

	return p
}
