package push

// push the top first element of stack b to stack a
func Pa(s1 []int, s2 []int) ([]int, []int) {
	holder := s2[0]
	s1 = append([]int{holder}, s1...)
	s2 = s2[1:]
	return s1, s2
}

// push the top first element of stack a to stack b
func Pb(s1 []int, s2 []int) ([]int, []int) {
	holder := s1[0]
	s2 = append([]int{holder}, s2...)
	s1 = s1[1:]
	return s1, s2
}

// swap first 2 elements of stack a
func Sa(s1 []int, s2 []int) ([]int, []int) {
	s1[0], s1[1] = s1[1], s1[0]
	return s1, s2
}

// swap first 2 elements of stack b
func Sb(s1 []int, s2 []int) ([]int, []int) {
	s2[0], s2[1] = s2[1], s2[0]
	return s1, s2
}

// execute sa and sb
func Ss(s1 []int, s2 []int) ([]int, []int) {
	slice1, _ := Sa(s1, s2)
	_, slice2 := Sb(s1, s2)
	return slice1, slice2
}

// rotate stack a (shift up all elements of stack a by 1, the first element becomes the last one)
func Ra(s1 []int, s2 []int) ([]int, []int) {
	holder := s1[0]
	s1 = s1[1:]
	s1 = append(s1, holder)
	return s1, s2
}

// rotate stack b (shift up all elements of stack a by 1, the first element becomes the last one)
func Rb(s1 []int, s2 []int) ([]int, []int) {
	holder := s2[0]
	s2 = s2[1:]
	s2 = append(s2, holder)
	return s1, s2
}

// execute ra and rb
func Rr(s1 []int, s2 []int) ([]int, []int) {
	slice1, _ := Ra(s1, s2)
	_, slice2 := Rb(s1, s2)
	return slice1, slice2
}

// reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
func Rra(s1 []int, s2 []int) ([]int, []int) {
	holder := s1[len(s1)-1]
	s1 = s1[:len(s1)-1]
	s1 = append([]int{holder}, s1...)
	return s1, s2
}

// reverse rotate b (shift down all elements of stack a by 1, the last element becomes the first one)
func Rrb(s1 []int, s2 []int) ([]int, []int) {
	holder := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	s2 = append([]int{holder}, s2...)
	return s1, s2
}

// execute rra and rrb
func Rrr(s1 []int, s2 []int) ([]int, []int) {
	slice1, _ := Rra(s1, s2)
	_, slice2 := Rrb(s1, s2)
	return slice1, slice2
}
