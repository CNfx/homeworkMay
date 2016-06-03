package pool

import (
	"fmt"
	"time"
)

type StringPool struct {
	idle   []string
	active []string
}

func (p *StringPool) Get() string {

	var d string

	if p.IsNotIdle() {
		d = genString()
	} else {
		// Pop one from idle array
		d = p.idle[0]
		p.idle = p.idle[1:len(p.idle)]
	}

	// Mark active
	p.active = append(p.active, d)
	return d
}

func (p *StringPool) Back(data string) {
	p.idle = append(p.idle, data)
	p.active = p.remove(p.active, data)
}

func (p *StringPool) remove(l []string, d string) []string {

	for k, v := range l {
		if v == d {
			return append(l[:k], l[k+1:]...)
		}
	}
	return l
}

func (p *StringPool) IsNotIdle() bool {
	return len(p.idle) == 0
}

func genString() string {
	return fmt.Sprintf("%d", time.Time{}.UnixNano())
}

func NewStringPool(length int) *StringPool {

	idle := []string{}
	for i := length; i > 0; i-- {
		idle = append(idle, genString())
	}

	return &StringPool{
		idle: idle,
	}

}
