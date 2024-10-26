package push

import (
	p "push/instructions"
)

func Checker(s1 []int) bool {
	for i := 0; i < len(s1); i++ {
		for j := i; j < len(s1); j++ {
			if i != j && s1[i] > s1[j] {
				return false
			}
		}
	}
	return true
}

func CaseThree(s1 []int, s2 []int) string {
	result := ""
	if Checker(s1) {
		return ""
	}
	if s1[0] > s1[1] && s1[1] < s1[2] {
		if s1[0] < s1[2] {
			s1, s2 = p.Sa(s1, s2)
			result += "sa" + "\n"
		} else {
			s1, s2 = p.Ra(s1, s2)
			result += "ra" + "\n"
		}
	} else if s1[0] > s1[1] && s1[1] > s1[2] {
		s1, s2 = p.Ra(s1, s2)
		result += "ra" + "\n"
		s1, s2 = p.Sa(s1, s2)
		result += "sa" + "\n"
	} else if s1[0] < s1[1] && s1[1] > s1[2] && s1[0] < s1[2] {
		s1, s2 = p.Sa(s1, s2)
		result += "sa" + "\n"
		s1, s2 = p.Ra(s1, s2)
		result += "ra" + "\n"
	} else if s1[0] < s1[1] && s1[1] > s1[2] && s1[0] > s1[2] {
		s1, s2 = p.Rra(s1, s2)
		result += "rra" + "\n"

	}
	return result
}

func Push_Swap(s1 []int, s2 []int) string {
	result := ""
	index := 0
	for len(s1) != 3 {
		smaller := s1[0]
		if !Checker(s1) {
			for i := 1; i < len(s1); i++ {
				if s1[i] < smaller {
					smaller = s1[i]
					index = i
				}
			}
			if (index != 0 && index != 1) && (len(s1) != 3) {
				if index > len(s1)/2 {
					for i := index; i < len(s1); i++ {
						s1, s2 = p.Rra(s1, s2)
						result += "rra" + "\n"
					}
					s1, s2 = p.Pb(s1, s2)
					result += "pb" + "\n"
				} else if index <= len(s1)/2 {
					for i := index; i > 0; i-- {
						s1, s2 = p.Ra(s1, s2)
						result += "ra" + "\n"
					}
					s1, s2 = p.Pb(s1, s2)
					result += "pb" + "\n"
				}
			} else {
				if index == 1 {
					s1, s2 = p.Sa(s1, s2)
					result += "sa" + "\n"
				}
				s1, s2 = p.Pb(s1, s2)
				result += "pb" + "\n"
			}
		} else {
			result += ""
			break
		}

	}
	if len(s1) == 3 {
		result += CaseThree(s1, s2)
	}
	for i := len(s2); i > 0; i-- {
		s1, s2 = p.Pa(s1, s2)
		result += "pa" + "\n"
	}
	return result
}
