package set

func NewSet() *Set {
	return &Set{
		isEmpty: true,
		size:    0,
		values:  make([]string, 10),
	}
}

func NewSetWithLimit(limit int) *Set {
	return &Set{
		isEmpty: true,
		size:    0,
		values:  make([]string, limit),
	}
}

type Set struct {
	isEmpty bool
	size    int
	values  []string
}

func (s *Set) IsEmpty() bool {
	return s.isEmpty
}

func (s *Set) Add(value string) {
	if len(s.values) == s.size {
		temp := make([]string, len(s.values)+1)
		copy(temp, s.values)
		s.values = temp
	}

	exists := s.Contains(value)
	if exists == false {
		s.values[s.size] = value
		s.isEmpty = false
		s.size++
	}

}

func (s *Set) Size() int {
	return s.size
}

func (s *Set) Contains(value string) bool {
	for i := 0; i < s.size; i++ {
		if s.values[i] == value {
			return true
		}
	}

	return false
}

func (s *Set) getPosition(value string) (int) {
	for i := 0; i < s.size; i++ {
		if s.values[i] == value {
			return i
		}
	}
	return 0
}
func (s *Set) Remove(value string) {
	exists:= s.Contains(value)
	if exists == true {
		position := s.getPosition(value)
		s.values[position] = s.values[s.size-1]
		s.values[s.size] = ""
		s.size--

		if s.size == 0 {
			s.isEmpty = true
		}

		return
	}
}
