package singleton

import "sync"

var (
	p *person
	once sync.Once
)

func GetInstance() *person  {
	once.Do(func() {
		p = &person{}
	})

	return p
}

type person struct {
	name string
	age int
	mux sync.Mutex
}

func (p *person) SetName(n string)  {
	p.name = n
}

func (p *person) GetName() string {
	return p.name
}

func (p *person) IncrementAge()  {
	p.mux.Lock()
	p.age++
	p.mux.Unlock()
}

func (p *person) GetAge() int  {
	p.mux.Lock()
	defer p.mux.Unlock()
	return  p.age
}




