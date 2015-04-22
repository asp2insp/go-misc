package algorithms

import "testing"
import "github.com/asp2insp/go-misc/testutils"

func TestIsPermutation(t *testing.T) {
	testutils.ExpectTrue(IsPermutation("abcd", "dacb"), "abcd is a permutation of dacb", t)
	testutils.ExpectFalse(IsPermutation("bad", "acb"), "bad is not a permutation of acb", t)
	testutils.ExpectTrue(IsPermutation("aaa", "aaa"), "aaa is a permutation of aaa", t)
	testutils.ExpectTrue(IsPermutation("a", "a"), "a is a permutation of itself", t)
}
