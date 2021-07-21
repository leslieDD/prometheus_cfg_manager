package utils

type NoBlockLock struct {
	ask  chan struct{}
	done chan struct{}
}

func NewNoBlockLock() *NoBlockLock {
	nb := NoBlockLock{
		ask:  make(chan struct{}),
		done: nil,
	}
	return &nb
}

func (n *NoBlockLock) running() {
	for {
		select {
		case n.ask <- struct{}{}:
		case <-n.done:
			return
		}
	}
}

// 返回true，说明已经有人在，返回false，则表示没有人应答
func (n *NoBlockLock) AnyOne() bool {
	select {
	case <-n.ask:
		return true
	default:
		n.done = make(chan struct{})
		go n.running()
		return false
	}
}

// 把自己清理
func (n *NoBlockLock) Done() {
	close(n.done)
}
