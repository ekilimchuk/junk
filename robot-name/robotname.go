package robotname

import (
	"fmt"
	"math/rand"
)

const MAX = 676000 //int(math.Pow(26.0, 2.0) * math.Pow(10.0, 3.0))

type Robot struct {
	ID string
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
	d := NewData()
	r.ID = ""
	d.DelKey(r.ID)
}

func (r *Robot) Name() string {
	d := NewData()
	if r.ID == "" {
		collision := 0
		for {
			r.ID = getRandString()
			if !d.ExistKey(r.ID) {
				d.AddKey(r.ID)
				break
			}
			collision++
			if collision >= MAX {
				panic("IDs overflow")
			}
		}
	}
	return r.ID
}
