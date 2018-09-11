package robotname

import (
	"math/rand"
	"time"
)

type Robot struct {
	ID string
}

func getRandBytes() []byte {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, 5)
	for i, _ := range bytes {
		switch {
		case i < 2:
			bytes[i] = byte(rand.Intn(int('Z'-'A'+1)) + 'A')
		default:
			bytes[i] = byte(rand.Intn(int('9'-'0'+1)) + '0')
		}
	}
	return bytes
}

func (r *Robot) Reset() {
	r.ID = ""
}

func (r *Robot) Name() string {
	if r.ID == "" {
		r.ID = string(getRandBytes())
	}
	return r.ID
}
