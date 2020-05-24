package matrix

//import "fmt"

type Matrix struct {
	rc, cc int
	data   [][]int
}

func NewMatrix(data [][]int) (bool, Matrix) {
	if len(data) == 0 || len(data[0]) == 0 {
		return false, Matrix{}
	}

	m := NewEmptyMatrix(len(data), len(data[0]))

	if !m.InitData(data) {
		return false, Matrix{}
	}

	return true, m
}

func NewEmptyMatrix(rc, cc int) Matrix {
	m := Matrix{
		rc,
		cc,
		make([][]int, rc),
	}
	for i := 0; i < m.GetRowCount(); i++ {
		m.data[i] = make([]int, cc)
	}
	return m
}

func (m *Matrix) GetRowCount() int {
	return m.rc
}

func (m *Matrix) GetColCount() int {
	return m.cc
}

func (m *Matrix) InitData(data [][]int) bool {
	rc := m.GetRowCount()
	cc := m.GetColCount()
	if len(m.data) != rc {
		return false
	}

	for i := 0; i < rc; i++ {
		if len(data[i]) != cc {
			return false
		}
	}
	m.data = data
	return true
}

func (m *Matrix) Get(i, j int) (bool, int) {
	if i >= m.GetRowCount() || j >= m.GetColCount() {
		return false, 0
	}
	return true, m.data[i][j]
}

func colSelected(c int, col []int) bool {
	for _, v := range col {
		if v == c {
			return true
		}
	}
	return false
}

func findV(s []int, val int) int {
	for i, v := range s {
		if v == val {
			return i
		}
	}
	return -1
}

func swap(s []int, i, j int) {
	tmp := s[j]
	s[j] = s[i]
	s[i] = tmp
}

func sign(s []int) int {
	ret := 1
	for i, v := range s {
		if i != v {
			idxv := findV(s, i)
			swap(s, i, idxv)
			ret *= -1
		}
	}
	return ret
}

func (m *Matrix) det(col []int) int {
	ret := 1
	// Terminal recursion if one value per row/col is selected
	// => return the build multiplication times the correction factor
	if len(col) >= m.GetRowCount() {
		for i, v := range col {
			ret *= m.data[i][v]
			// fmt.Printf("a(%v, %v) * ", i+1, v+1)
		}
		s := sign(col)
		// fmt.Printf("%2d\n", s)
		return ret * s
	}

	ret = 0
	for i := 0; i < m.GetColCount(); i++ {
		if colSelected(i, col) {
			continue
		}
		ret += m.det(append(append([]int(nil), col...), i))
	}
	return ret
}

func (m *Matrix) Det() (bool, int) {
	if m.GetColCount() != m.GetRowCount() {
		return false, 0
	}

	return true, m.det(nil)
}
