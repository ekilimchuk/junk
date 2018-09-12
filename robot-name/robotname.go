package robotname

import (
	"fmt"
	"math/rand"
	"sync"
)

// MAX is a maximum count of IDs.
const MAX = 676000 //int(math.Pow(26.0, 2.0) * math.Pow(10.0, 3.0))

type Robot struct {
	l         sync.Mutex
	currentID string
	oldIDs    map[string]bool
}

func getRandChar() rune {
	return rune(rand.Intn(int('Z'-'A'+1)) + 'A')
}

func getRandNumb() rune {
	return rune(rand.Intn(int('9'-'0'+1)) + '0')
}

func getRandString() string {
	return fmt.Sprintf("%c%c%c%c%c", getRandChar(), getRandChar(), getRandNumb(), getRandNumb(), getRandNumb())
}

func (r *Robot) Reset() {
	r.l.Lock()
	d := NewData()
	d.DelKey(r.currentID)
	if r.oldIDs == nil {
		r.oldIDs = map[string]bool{}
	}
	r.oldIDs[r.currentID] = true
	r.currentID = ""
	r.l.Unlock()
}

func (r *Robot) Name() string {
	r.l.Lock()
	globalIDs := NewData()
	if r.oldIDs == nil {
		r.oldIDs = map[string]bool{}
	}
	if r.currentID == "" {
		collision := 0
		for {
			r.currentID = getRandString()
			if _, ok := r.oldIDs[r.currentID]; !globalIDs.ExistKey(r.currentID) && !ok {
				globalIDs.AddKey(r.currentID)
				break
			}
			collision++
			if collision >= MAX {
				panic("IDs overflow")
			}
		}
	}
	r.l.Unlock()
	return r.currentID
}
