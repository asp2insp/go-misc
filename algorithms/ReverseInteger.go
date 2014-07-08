package algorithms

//import "fmt"

func ReverseInteger(val uint32) uint32 {
	var ret uint32
	var exp uint32
	for exp = 1; exp <= val; exp *= BASE {
		ret *= BASE
		ret += (val % (exp * BASE)) / exp
	}
	return ret
}
