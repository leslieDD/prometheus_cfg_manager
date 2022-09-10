package alertmgr

import (
	"fmt"
	"pro_cfg_manager/config"
	"sync"
	"time"

	"golang.org/x/net/proxy"
)

type ProxyManager struct {
	QueueConn    chan *proxy.Dialer
	ProxyConn    *map[string]*proxySaveInfo
	NewProxyConn *map[string]*proxySaveInfo
	lock         sync.RWMutex
	Quit         chan struct{}
	Last         time.Time
}

type proxySaveInfo struct {
	N          string
	D          *proxy.Dialer
	T          time.Time
	ExpireTime time.Time
}

var PrManager *ProxyManager

func init() {
	PrManager = NewProxyManager()
}

func NewProxyManager() *ProxyManager {
	pc := &ProxyManager{
		QueueConn:    make(chan *proxy.Dialer, 10),
		ProxyConn:    &map[string]*proxySaveInfo{},
		NewProxyConn: &map[string]*proxySaveInfo{},
		lock:         sync.RWMutex{},
		Quit:         make(chan struct{}),
		Last:         time.Now(),
	}
	go pc.Listen()
	return pc
}

func (p *ProxyManager) Listen() {
	for {
		select {
		case <-p.Quit:
			return
		default:
			p.lock.Lock()
			for _, newPConn := range *p.NewProxyConn {
				(*p.ProxyConn)[newPConn.N] = newPConn
			}
			for name, nowPConn := range *p.ProxyConn {
				if (!nowPConn.ExpireTime.IsZero()) && time.Now().Sub(nowPConn.ExpireTime).Seconds() > 0 {
					delete(*p.ProxyConn, name)
				}
			}
			p.NewProxyConn = &map[string]*proxySaveInfo{}
			p.lock.Unlock()
			for _, pc := range *p.ProxyConn {
				if (!pc.ExpireTime.IsZero()) && time.Now().Sub(pc.ExpireTime).Seconds() > 0 {
					continue
				}
				select {
				case <-p.Quit:
					return
				case p.QueueConn <- pc.D:
					pc.T = time.Now()
				}
			}
			if len(*p.ProxyConn) == 0 {
				time.Sleep(1 * time.Second)
			}
		}
	}
}

func (p *ProxyManager) GetDial(t *Task) *proxy.Dialer {
	go p.RecvNew(t)
	select {
	case <-p.Quit:
		return nil
	case d := <-p.QueueConn:
		return d
	case <-time.After(10 * time.Second):
		return nil
	}
}

func (p *ProxyManager) RecvNew(t *Task) {
	name := p.CreateName(t)
	var ok bool
	p.lock.Lock()
	defer p.lock.Unlock()
	_, ok = (*p.ProxyConn)[name]
	if ok {
		return
	}
	_, ok = (*p.NewProxyConn)[name]
	if ok {
		return
	}
	d2, e := p.CreateConn(t)
	if e != nil {
		config.Log.Error(e)
		return
	}
	(*p.NewProxyConn)[name] = &proxySaveInfo{D: d2, T: time.Now(), N: name, ExpireTime: t.ProxyExpireTime}
	return
}

func (p *ProxyManager) CreateName(t *Task) string {
	name := fmt.Sprintf("%v_%v_%v", t.ProxyHost, t.ProxyPort, t.ProxyAuthUser)
	//name := fmt.Sprintf("%v_%v_%v_%v", t.PlanID, t.ProxyHost, t.ProxyPort, t.ProxyAuthUser)
	return name
}

func (p *ProxyManager) CreateConn(t *Task) (*proxy.Dialer, error) {
	proxyAuth := &socket5{
		Auth:    nil,
		Network: "tcp",
		Host:    t.ProxyHost,
		Port:    t.ProxyPort,
	}
	if t.ProxyAuth {
		proxyAuth.Auth = &proxy.Auth{
			User:     t.ProxyAuthUser,
			Password: t.ProxyAuthPass,
		}
	}
	dialer, err := proxy.SOCKS5(
		proxyAuth.Network,
		proxyAuth.Host+":"+proxyAuth.Port,
		proxyAuth.Auth,
		proxy.Direct,
	)
	if err != nil {
		return nil, err
	}
	return &dialer, nil
}

func (p *ProxyManager) Stop() {

}
