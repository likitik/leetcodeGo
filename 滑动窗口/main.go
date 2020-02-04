package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	var left, right, match int
	var win [26]int
	var need [26]int
	var needs int

	for _, i := range s1 {
		if need[i-'a'] == 0 {
			needs++
		}
		need[i-'a']++
	}
	for right < len(s2) {
		c1 := s2[right] - 'a'
		if need[c1] > 0 {
			win[c1]++
			if win[c1] == need[c1] {
				match++
			}
		}
		right++
		for match == needs {
			fmt.Println(right, left)
			if right-left+1 == len(s1) {
				return true
			}
			c2 := s2[left] - 'a'
			if need[c2] > 0 {
				win[c2]--
				if win[c2] < need[c2] {
					match--
				}
			}
			left++
		}
	}
	return false
}

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
}
