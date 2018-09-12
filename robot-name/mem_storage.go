package robotname

import (
	"sync"
)

type data struct {
	l sync.Mutex
	m map[string]bool
}

var gd data

func shareData() *data {
	gd.l.Lock()
	defer gd.l.Unlock()
	if gd.m == nil {
		gd.m = map[string]bool{}
	}
	return &gd
}

func (d *data) key(s string, b bool) {
	d.l.Lock()
	defer d.l.Unlock()
	d.m[s] = b
}

func (d *data) addKey(s string) {
	d.key(s, true)
}

func (d *data) delKey(s string) {
	d.key(s, false)
}

func (d *data) existKey(s string) bool {
	d.l.Lock()
	defer d.l.Unlock()
	if _, ok := d.m[s]; !ok {
		return false
	}
	return d.m[s]
}
