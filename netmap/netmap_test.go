package netmap

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func BenchmarkXxx(b *testing.B) {

	var d byte

	var r = make([]byte, 4)
	rand.Read(r)
	for k := 0; k < b.N; k++ {
		for range r {
			for i := uint8(0); i < 4; i++ {
				// d = v & (0b11 << i)
				d = d + i
			}
			// d = d + 1
			// d = d + 2
			// d = d + 3
			// d = d + 4

			// d = v & (0b11 << 0)
			// d = v&(0b11<<1) + d
			// d = v&(0b11<<2) - d
			// d = v&(0b11<<3) + d
		}
	}

	if false {
		b.Log(d)
	}
}

func TestXxx(t *testing.T) {

	m := NewMap()

	var key = []byte{172, 17, 16, 1}
	var val = uint32(1998)

	m.Set(key, val)

	r := m.Get(key)
	t.Log(r)
}

func Test_Size(t *testing.T) {

	// n bits
	// step x
	//
	// 存一个val需要B:
	// n/x * 2^x^
	//

	m := NewMap()

	for i, ip := range ips {
		m.Set(ip, uint32(i+1))
	}

	t.Log(len(m.list))

	var tt, u uint32
	for i, v := range m.list {

		var null = true
		for _, uv := range v {
			tt = tt + 4
			if uv != 0 {
				u = u + 4
				null = false
			}
		}
		if null {
			fmt.Println("null: ", i)
		}

	}

	t.Log(float64(u) / float64(tt))
	t.Log(float64(len(ips)*(4+4)) / float64(tt))
}
