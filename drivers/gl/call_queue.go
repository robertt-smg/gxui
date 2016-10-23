package gl

import "sync"

type CallQueue struct {
	mu     sync.Mutex
	onDeck chan *callNode
	head   *callNode
	tail   *callNode
}

func NewCallQueue() *CallQueue {
	return &CallQueue{
		onDeck: make(chan *callNode, 1),
	}
}

func (c *CallQueue) Inject(call func()) {
	node := &callNode{v: call}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.head == nil {
		c.head = node
		c.tail = node
		c.onDeck <- c.head
		return
	}
	c.tail.next = node
	c.tail = node
}

func (c *CallQueue) Pop() (func(), bool) {
	select {
	case node, ok := <-c.onDeck:
		if !ok {
			return nil, false
		}
		c.shift()
		return node.v, true
	default:
		return nil, true
	}
}

func (c *CallQueue) PopWhenReady() (func(), bool) {
	node, ok := <-c.onDeck
	if !ok {
		return nil, false
	}
	c.shift()
	return node.v, true
}

func (c *CallQueue) Close() {
	close(c.onDeck)
}

func (c *CallQueue) shift() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.head = c.head.next
	if c.head == nil {
		c.tail = nil
		return
	}
	c.onDeck <- c.head
}

type callNode struct {
	v    func()
	next *callNode
}
