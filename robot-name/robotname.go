package robotname

import (
	"math/rand"
	"time"
)

type Robot struct {
	ID string
}

func getRandChar() byte {
	rand.Seed(time.Now().UnixNano())
	return byte(rand.Intn(int('Z'-'A'+1)) + 'A')
}

func getRandNumb() byte {
	rand.Seed(time.Now().UnixNano())
	return byte(rand.Intn(int('9'-'0'+1)) + '0')
}

func getRandBytes() []byte {
	bytes := make([]byte, 5)
	for i, _ := range bytes {
		switch {
		case i < 2:
			bytes[i] = getRandChar()
		default:
			bytes[i] = getRandNumb()
		}
	}
	return bytes
}

func (r *Robot) Reset() {
	d := NewData()
	r.ID = ""
	d.DelKey(r.ID)
}

func (r *Robot) Name() string {
	d := NewData()
	if r.ID == "" {
		for {
			r.ID = string(getRandBytes())
			if !d.ExistKey(r.ID) {
				d.AddKey(r.ID)
				break
			}
		}
	}
	return r.ID
}
