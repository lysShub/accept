package netmap

type Map struct {

	// list[0]是entry, 任何idx==0表示不存在
	list []block

	idles []uint32
}

type block [4]uint32

const s = 0b11

func NewMap() *Map {
	return &Map{
		list:  make([]block, 1, 64),
		idles: make([]uint32, 0, 16),
	}
}

func (m *Map) Set(k []byte, id uint32) {
	var idx uint32
	for i, b := range k {

		if m.list[idx][b&s] == 0 {
			m.list[idx][b&s] = m.get()
		}
		idx = m.list[idx][b&s]

		if m.list[idx][b>>2&s] == 0 {
			m.list[idx][b>>2&s] = m.get()
		}
		idx = m.list[idx][b>>2&s]

		if m.list[idx][b>>4&s] == 0 {
			m.list[idx][b>>4&s] = m.get()
		}
		idx = m.list[idx][b>>4&s]

		if i == len(k)-1 {
			m.list[idx][b>>6&s] = id
		} else {
			if m.list[idx][b>>6&s] == 0 {
				m.list[idx][b>>6&s] = m.get()
			}
			idx = m.list[idx][b>>6&s]
		}
	}
}

func (m *Map) get() (idx uint32) {
	if cap(m.list) > len(m.list) {
		idx = uint32(len(m.list))
		m.list = m.list[:idx+1]
		return idx
	} else if len(m.idles) > 0 {
		idx = m.idles[len(m.idles)-1]
		m.idles = m.idles[:len(m.idles)-1]
		return idx
	} else {
		m.list = append(m.list, block{})
		return uint32(len(m.list) - 1)
	}
}

func (m *Map) Get(k []byte) (id uint32) {

	n := len(k)
	b := byte(0)
	for i := 0; i < n; i++ {
		// TODO: optimize
		b = k[i]

		id = m.list[id][b&s]
		if id == 0 {
			return
		}

		id = m.list[id][b>>2&s]
		if id == 0 {
			return
		}

		id = m.list[id][b>>4&s]
		if id == 0 {
			return
		}

		id = m.list[id][b>>6&s]
		if id == 0 {
			return
		}
	}

	return
}
