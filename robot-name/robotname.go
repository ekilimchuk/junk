package robotname

import (
	"fmt"
	"math/rand"
	"sync"
)

// MAX is a maximum count of IDs.
const MAX = 676000 // int(math.Pow(26.0, 2.0) * math.Pow(10.0, 3.0))

// Robot is a safe struct which stores old IDs an a current ID.
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

// Reset stores old IDs and resets a current ID.
func (r *Robot) Reset() {
	r.l.Lock()
	defer r.l.Unlock()
	globalIDs := shareData()
	globalIDs.delKey(r.currentID)
	if r.oldIDs == nil {
		r.oldIDs = map[string]bool{}
	}
	r.oldIDs[r.currentID] = true
	r.currentID = ""
}

// Name generates a new ID.
func (r *Robot) Name() string {
	r.l.Lock()
	defer r.l.Unlock()
	globalIDs := shareData()
	if r.oldIDs == nil {
		r.oldIDs = map[string]bool{}
	}
	if r.currentID != "" {
		return r.currentID
	}
	collision := 0
	// Look for a free ID in a global struct and an object struct (a local).
	for {
		r.currentID = getRandString()
		if _, ok := r.oldIDs[r.currentID]; !globalIDs.existKey(r.currentID) && !ok {
			globalIDs.addKey(r.currentID)
			break
		}
		collision++
		// Check collision and return panic on overflow.
		if collision >= MAX {
			panic("IDs overflow")
		}
	}
	return r.currentID
}
