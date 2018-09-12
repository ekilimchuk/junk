package robotname

import (
	"sync"
)

type Data struct {
	l sync.Mutex
	m map[string]bool
}

var D Data

func NewData() *Data {
	D.l.Lock()
	if D.m == nil {
		D.m = map[string]bool{}
	}
	D.l.Unlock()
	return &D
}

func (d *Data) key(s string, b bool) {
	d.l.Lock()
	d.m[s] = b
	d.l.Unlock()
}

func (d *Data) AddKey(s string) {
	d.key(s, true)
}

func (d *Data) DelKey(s string) {
	d.key(s, false)
}

func (d *Data) ExistKey(s string) bool {
	d.l.Lock()
	defer d.l.Unlock()
	if _, ok := d.m[s]; !ok {
		return false
	}
	return d.m[s]
}
