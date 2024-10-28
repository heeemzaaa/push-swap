package push

func Pa(s1 []int, s2 []int) ([]int, []int) {
	holder := s2[0]
	s1 = append([]int{holder}, s1...)
	s2 = s2[1:]
	return s1, s2
}

func Pb(s1 []int, s2 []int) ([]int, []int) {
	holder := s1[0]
	s2 = append([]int{holder}, s2...)
	s1 = s1[1:]
	return s1, s2
}

func Sa(s1 []int, s2 []int) ([]int, []int) {
	s1[0], s1[1] = s1[1], s1[0]
	return s1, s2
}

func Sb(s1 []int, s2 []int) ([]int, []int) {
	s2[0], s2[1] = s2[1], s2[0]
	return s1, s2
}

func Ss(s1 []int, s2 []int) ([]int, []int) {
	slice1, _ := Sa(s1, s2)
	_, slice2 := Sb(s1, s2)
	return slice1, slice2
}

func Ra(s1 []int, s2 []int) ([]int, []int) {
	holder := s1[0]
	s1 = s1[1:]
	s1 = append(s1, holder)
	return s1, s2
}

func Rb(s1 []int, s2 []int) ([]int, []int) {
	holder := s2[0]
	s2 = s2[1:]
	s2 = append(s2, holder)
	return s1, s2
}

func Rr(s1 []int, s2 []int) ([]int, []int) {
	slice1, _ := Ra(s1, s2)
	_, slice2 := Rb(s1, s2)
	return slice1, slice2
}

func Rra(s1 []int, s2 []int) ([]int, []int) {
	holder := s1[len(s1)-1]
	s1 = s1[:len(s1)-1]
	s1 = append([]int{holder}, s1...)
	return s1, s2
}

func Rrb(s1 []int, s2 []int) ([]int, []int) {
	holder := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	s2 = append([]int{holder}, s2...)
	return s1, s2
}

func Rrr(s1 []int, s2 []int) ([]int, []int) {
	slice1, _ := Rra(s1, s2)
	_, slice2 := Rrb(s1, s2)
	return slice1, slice2
}
