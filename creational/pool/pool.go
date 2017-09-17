package pool

import (
	"container/list"
)

type PoolObject struct {
	id int
}

type Pool struct {
	idle   *list.List
	active *list.List
}

func InitPool() Pool {
	pool := &Pool{
		list.New(),
		list.New(),
	}
	return *pool
}

func (p *Pool) Loan() PoolObject {
	if p.idle.Len() > 0 {
		for e, i := p.idle.Front(), 0; e != nil; e, i = e.Next(), i+1 {
			if i == 0 {
				object := e.Value.(PoolObject)
				return object
			}
		}
	}

	object := PoolObject{p.NumberOfObjectsInPool() + 1}
	p.active.PushBack(object)
	return object
}

func (p *Pool) Receive(object PoolObject) {
	p.idle.PushBack(object)
	for e, i := p.active.Front(), 0; e != nil; e, i = e.Next(), i+1 {
		if object == e.Value.(PoolObject) {
			p.active.Remove(e)
			return
		}
	}
}

func (p *Pool) NumberOfObjectsInPool() int {
	return p.active.Len() + p.idle.Len()
}

func (p *Pool) NumberOfActiveObjects() int {
	return p.active.Len()
}

func (p *Pool) NumberOfIdleObjects() int {
	return p.idle.Len()
}
