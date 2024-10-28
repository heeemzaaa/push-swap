package main

import (
	"bufio"
	"fmt"
	"os"
	p "push/instructions"
	"strconv"
	"strings"
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
	scanner := bufio.NewScanner(os.Stdin)
	slice := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		slice = append(slice, line)
	}
	s1 := convert
	s2 := []int{}
	s1, s2 = Apply(s1, s2, slice)
	if (len(s1) == 0 && len(s2) == 0) || Duplicates(s1) {
		fmt.Println("Error")
		return
	}
	finalResult := Checker(s1, s2)
	fmt.Println(finalResult)
}

func Duplicates(s1 []int) bool {
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s1); j++ {
			if i != j && s1[i] == s1[j] {
				return true
			}
		}
	}
	return false
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

func Apply(s1, s2 []int, slice []string) ([]int, []int) {
	for i := 0; i < len(slice); i++ {
		if slice[i] == "pa" {
			s1, s2 = p.Pa(s1, s2)
		} else if slice[i] == "pb" {
			s1, s2 = p.Pb(s1, s2)
		} else if slice[i] == "sa" {
			s1, s2 = p.Sa(s1, s2)
		} else if slice[i] == "sb" {
			s1, s2 = p.Sb(s1, s2)
		} else if slice[i] == "ss" {
			s1, s2 = p.Ss(s1, s2)
		} else if slice[i] == "ra" {
			s1, s2 = p.Ra(s1, s2)
		} else if slice[i] == "rb" {
			s1, s2 = p.Rb(s1, s2)
		} else if slice[i] == "rr" {
			s1, s2 = p.Rr(s1, s2)
		} else if slice[i] == "rra" {
			s1, s2 = p.Rra(s1, s2)
		} else if slice[i] == "rrb" {
			s1, s2 = p.Rrb(s1, s2)
		} else if slice[i] == "rrr" {
			s1, s2 = p.Rrr(s1, s2)
		} else if slice[i] == "" {
			continue
		} else {
			return []int{}, []int{}
		}
	}
	return s1, s2
}

func Sorted(s1 []int) bool {
	for i := 0; i < len(s1)-1; i++ {
		if s1[i] > s1[i+1] {
			return false
		}
	}
	return true
}

func Checker(s1, s2 []int) string {
	if len(s2) == 0 && Sorted(s1) {
		return "OK"
	}
	return "KO"
}
