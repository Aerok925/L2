package main

import (
	"fmt"
	"sort"
	"strings"
)

type mapped map[string][]string

type myRune struct {
	str []rune
}

func (r *myRune) Len() int {
	return len(r.str)
}

func NewMyRune() *myRune {
	return &myRune{}
}

func (r *myRune) Swap(i, j int) {
	r.str[i], r.str[j] = r.str[j], r.str[i]
}

func (r *myRune) Less(i, j int) bool {
	return r.str[i] < r.str[j]
}

func deleteSinglesAnagram(m mapped) {
	for key, value := range m {
		if len(value) == 0 {
			delete(m, key)
		}
	}
}

func isAnagram(str, str2 string) bool {
	if str == str2 {
		return false
	}
	first := NewMyRune()
	second := NewMyRune()
	first.str = []rune(str)
	second.str = []rune(str2)
	sort.Sort(first)
	sort.Sort(second)
	if strings.Compare(string(first.str), string(second.str)) == 0 {
		return true
	}
	return false
}

func AnagramInMap(strs []string) mapped {
	retMap := make(mapped)
	for _, str := range strs {
		str = strings.ToLower(str)
		if !func() bool {
			for key, _ := range retMap {
				if isAnagram(key, str) {
					for _, strn := range retMap[key] {
						if strn == str {
							return false
						}
					}
					retMap[key] = append(retMap[key], str)
					return true
				}
			}
			return false
		}() {
			if _, ok := retMap[str]; ok {
				continue
			}
			retMap[str] = make([]string, 0)
		}
	}
	deleteSinglesAnagram(retMap)
	for _, value := range retMap {
		sort.Strings(value)
	}
	return retMap
}

func main() {
	str := "икс кси оск сип пик пси сип сок иск сип"
	strs := strings.Split(str, " ")
	fmt.Println(AnagramInMap(strs))

}
