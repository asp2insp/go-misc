package algorithms

import "strings"

func IsPermutation(s, t string) bool {
	if len(s) == 0 && len(t) == 0 {
		return true
	} else if len(s) == 0 || len(t) == 0 {
		return false
	} else if strings.Contains(t, string(s[0])) {
		return IsPermutation(string(s[1:]), strings.Replace(t, string(s[0]), "", 1))
	} else {
		return false
	}
}
