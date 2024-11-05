package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	p "push/instructions"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	splitted_arg := strings.Split(os.Args[1], " ")
	convert := Converting(splitted_arg)
	if len(convert) == 0 {
		fmt.Println("Error")
		return
	}
	s1 := convert
	s2 := []int{}
	if !Duplicates(s1) {
		fmt.Println("Error")
		return
	}
	finalResult := Push_Swap(s1, s2)
	fmt.Print(finalResult)
}

func Converting(slice []string) []int {
	converted := []int{}
	for i := 0; i < len(slice); i++ {
		if slice[i] == "" {
			continue
		}
		holder, err := strconv.Atoi(slice[i])
		if err != nil {
			return []int{}
		}
		converted = append(converted, holder)
	}
	return converted
}

func Duplicates(s1 []int) bool {
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s1); j++ {
			if i != j && s1[i] == s1[j] {
				return false
			}
		}
	}
	return true
}

func Sorted(s1 []int) bool {
	for i := 0; i < len(s1)-1; i++ {
		if s1[i] > s1[i+1] {
			return false
		}
	}
	return true
}

func HandleThree(s1 []int, s2 []int) ([]int, []int) {
	if Sorted(s1) {
		return s1, s2
	}
	if s1[0] > s1[1] && s1[1] < s1[2] {
		if s1[0] < s1[2] {
			s1, s2 = p.Sa(s1, s2)
		} else {
			s1, s2 = p.Ra(s1, s2)
		}
	} else if s1[0] > s1[1] && s1[1] > s1[2] {
		s1, s2 = p.Ra(s1, s2)
		s1, s2 = p.Sa(s1, s2)
	} else if s1[0] < s1[1] && s1[1] > s1[2] && s1[0] < s1[2] {
		s1, s2 = p.Sa(s1, s2)
		s1, s2 = p.Ra(s1, s2)
	} else if s1[0] < s1[1] && s1[1] > s1[2] && s1[0] > s1[2] {
		s1, s2 = p.Rra(s1, s2)
	}
	return s1, s2
}

func CaseThree(s1 []int, s2 []int) string {
	result := ""
	if Sorted(s1) {
		return ""
	}
	if s1[0] > s1[1] && s1[1] < s1[2] {
		if s1[0] < s1[2] {
			result += "sa" + "\n"
		} else {
			result += "ra" + "\n"
		}
	} else if s1[0] > s1[1] && s1[1] > s1[2] {
		result += "ra" + "\n"
		result += "sa" + "\n"
	} else if s1[0] < s1[1] && s1[1] > s1[2] && s1[0] < s1[2] {
		result += "sa" + "\n"
		result += "ra" + "\n"
	} else if s1[0] < s1[1] && s1[1] > s1[2] && s1[0] > s1[2] {
		result += "rra" + "\n"
	}
	return result
}

func Smaller(s1 []int) (int, int) {
	smaller := s1[0]
	index := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] < smaller {
			smaller = s1[i]
			index = i
		}
	}
	return smaller, index
}

func Push_Swap(s1 []int, s2 []int) string {
	result := ""
	for len(s1) != 3 {
		if !Sorted(s1) {
			_, index := Smaller(s1)
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
				if index == 1 && s1[0] > s1[1] {
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
		s1, s2 = HandleThree(s1, s2)
	}
	for i := len(s2); i > 0; i-- {
		s1, s2 = p.Pa(s1, s2)
		result += "pa" + "\n"
	}
	fmt.Println(s1)
	return result
}
