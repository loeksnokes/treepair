package treepair

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

/* PrefixCode can (23-05-2020), all tested:
    1) Describe its cardinality.
    2) expand at a location.
    3) reduce at a location.
    4) list exposed carets.
	5) print itself.

	TODO (23-05-2020):
	1) Meet and Join.
	2) say which part of the code is a prefix of a long enough entry.
	3) list integer values across an exposed caret in alphabet order of children.

	Currently, a prefix code must have at least one child for each letter of
	alphabet: the empty string is not a prefix code, and this should probably
	be changed.
*/

func Test(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	// MakeTreePairByExpansions expands out a tree pair
	// from domain and range lists of expansion points.
	t.Run("MakeByExpansion test", func(t *testing.T) {
		// makes permutation 0 1 2 3 ... 7
		tp, err := NewTreePairAlpha("01")
		if nil != err {
			assertCorrectMessage(t, "Failed to NewTreePairAlpha('01')", " in MakeByExpansion test.")
		}
		tp.ExpandDomainAt("01")
		tp.ExpandRangeAt("10")
		got := tp.FullString()
		want := "{D: [00 0], [010 1], [011 2], [100 3], [101 4], [11 5] || R: [00 0], [010 1], [011 2], [100 3], [101 4], [11 5]}"
		//fmt.Println("First Test: ")
		//fmt.Println("Got:  " + got)
		//fmt.Println("Want: " + want)
		assertCorrectMessage(t, got, want)
	})

	// MakeTreePairByDFSCode creates tree pair from
	// Depth First Search description of trees.
	t.Run("MakeByDFSString test", func(t *testing.T) {
		//makes permutation 1 2 3 ... 8
		tp, err := NewTreePairAlpha("01")
		if nil != err {
			assertCorrectMessage(t, "Failed to NewTreePairAlpha('01')", " in MakeByDFSString test.")
		}

		EncodeDFS(tp, "{1111000011000,1110100010100,0 1 2 3 4 5 6}")
		got := tp.FullString()
		want := "{D: [0000 0], [0001 1], [001 2], [01 3], [100 4], [101 5], [11 6] || R: [000 0], [0010 1], [0011 2], [01 3], [10 4], [110 5], [111 6]}"
		//fmt.Println("Second Test: ")
		//fmt.Println("Got:  " + got)
		//fmt.Println("Want: " + want)
		assertCorrectMessage(t, got, want)
	})

	// Minimise reduces tree pair
	t.Run("Minimise test", func(t *testing.T) {
		//reduces element to minimal tree pair.
		// makes permutation 1 2 3 ... 8
		//fmt.Println("Entered Third Test (minimise)")
		tp, err := NewTreePairAlpha("01")
		if nil != err {
			assertCorrectMessage(t, "Failed to NewTreePairAlpha('01')", " in Minimise test.")
		}

		got := ""
		want := ""
		if !EncodeDFS(tp, "{111000100,111100000,0 1 2 3 4}") {
			got = "tp was poorly formed."
			want = "{D: [0 0], [10 1], [11 2] || R: [00 0], [01 1], [1 2]}"
			assertCorrectMessage(t, got, want)
		}
		//fmt.Println("Unminimised: " + tp.FullString())
		tp.Minimise()
		//fmt.Println("Minimised: " + tp.FullString())
		got = tp.FullString()
		want = "{D: [0 0], [10 1], [11 2] || R: [00 0], [01 1], [1 2]}"
		assertCorrectMessage(t, got, want)
	})

	// Minimize reduces tree pairs for Americans
	t.Run("Minimize test", func(t *testing.T) {
		//reduces element to minimal tree pair.
		// makes permutation 1 2 3 ... 8
		//fmt.Println("Entered Third Test (minimise)")
		tp, err := NewTreePairAlpha("01")
		if nil != err {
			assertCorrectMessage(t, "Failed to NewTreePairAlpha('01')", " in Minimize test.")
		}

		got := ""
		want := ""
		if !EncodeDFS(tp, "{111000100,111100000,0 1 2 3 4}") {
			got = "tp was poorly formed."
			want = "{D: [0 0], [10 1], [11 2] || R: [00 0], [01 1], [1 2]}"
			assertCorrectMessage(t, got, want)
		}
		//fmt.Println("Unminimised: " + tp.FullString())
		tp.Minimize()
		//fmt.Println("Minimised: " + tp.FullString())
		got = tp.FullString()
		want = "{D: [0 0], [10 1], [11 2] || R: [00 0], [01 1], [1 2]}"
		assertCorrectMessage(t, got, want)
	})

	// Permutation applies permutation at range.
	t.Run("Permute range test", func(t *testing.T) {
		tp, err := NewTreePairAlpha("01")
		if nil != err {
			assertCorrectMessage(t, "Failed to NewTreePairAlpha('01')", " in PermuteRange test.")
		}

		EncodeDFS(tp, "{110011000,101010100,0 1 2 3 4}")
		permutation := map[int]int{
			0: 1,
			1: 4,
			2: 2,
			3: 0,
			4: 3,
		}
		tp.ApplyPermRange(permutation)
		got := tp.FullString()
		want := "{D: [00 0], [01 1], [100 2], [101 3], [11 4] || R: [0 1], [10 4], [110 2], [1110 0], [1111 3]}"
		assertCorrectMessage(t, got, want)
	})

	// PermuteRange applies permutation at domain.
	t.Run("Permute domain test", func(t *testing.T) {
		tp, err := NewTreePairAlpha("01")
		if nil != err {
			assertCorrectMessage(t, "Failed to NewTreePairAlpha('01')", " in PermuteDomain test.")
		}

		EncodeDFS(tp, "{110011000,101010100,0 1 2 3 4}")
		permutation := map[int]int{
			0: 1,
			1: 4,
			2: 2,
			3: 0,
			4: 3,
		}
		tp.ApplyPermDomain(permutation)
		got := tp.FullString()
		want := "{D: [00 1], [01 4], [100 2], [101 0], [11 3] || R: [0 0], [10 1], [110 2], [1110 3], [1111 4]}"
		assertCorrectMessage(t, got, want)
	})

	// PermuteLabels applies permutation at domain and range.
	t.Run("Permute Labels test", func(t *testing.T) {
		tp, err := NewTreePairAlpha("01")
		if nil != err {
			assertCorrectMessage(t, "Failed to NewTreePairAlpha('01')", " in Permute Labels test.")
		}

		EncodeDFS(tp, "{110011000,101010100,0 1 2 3 4}")
		permutation := map[int]int{
			0: 1,
			1: 4,
			2: 2,
			3: 0,
			4: 3,
		}
		tp.PermuteLabels(permutation)
		got := tp.FullString()
		want := "{D: [00 1], [01 4], [100 2], [101 0], [11 3] || R: [0 1], [10 4], [110 2], [1110 0], [1111 3]}"
		assertCorrectMessage(t, got, want)
	})

	// ResetLabels forces domain to be labelled in natural order and
	// relabels range to maintain the actual element.
	t.Run("ResetLabels test", func(t *testing.T) {

		tp, err := NewTreePairAlpha("01")
		if nil != err {
			assertCorrectMessage(t, "Failed to NewTreePairAlpha('01')", " in ResetLables test.")
		}

		EncodeDFS(tp, "{110011000,101010100,0 1 2 3 4}")
		permutation := map[int]int{
			0: 1,
			1: 4,
			2: 2,
			3: 0,
			4: 3,
		}
		tp.PermuteLabels(permutation)
		tp.ResetLabels()
		got := tp.FullString()
		want := "{D: [00 0], [01 1], [100 2], [101 3], [11 4] || R: [0 0], [10 1], [110 2], [1110 3], [1111 4]}"
		assertCorrectMessage(t, got, want)
	})

	// Inverse swaps the tree pairs.  It does not reset labels afterward.
	t.Run("Inverse test", func(t *testing.T) {
		tp, err := NewTreePairAlpha("01")
		if nil != err {
			assertCorrectMessage(t, "Failed to NewTreePairAlpha('01')", " in Inverse test.")
		}
		EncodeDFS(tp, "{110011000,101010100,0 1 2 3 4}")
		permutation := map[int]int{
			0: 1,
			1: 4,
			2: 2,
			3: 0,
			4: 3,
		}
		tp.ApplyPermRange(permutation)
		tp.ResetLabels()
		tp.Invert()
		tp.ResetLabels()
		got := tp.FullString()
		want := "{D: [0 0], [10 1], [110 2], [1110 3], [1111 4] || R: [00 3], [01 0], [100 2], [101 4], [11 1]}"
		assertCorrectMessage(t, got, want)
	})

	//TODO: fix this test.
	t.Run("Multiply test", func(t *testing.T) {
		//reduces element to minimal tree pair.
		// makes permutation 0 1 2 3 ... 7
		dTP, err := NewTreePairAlpha("01")
		if nil != err {
			assertCorrectMessage(t, "Failed to NewTreePairAlpha('01')", " in domaintp constructor in Multiply test.")
		}

		rTP, err := NewTreePairAlpha("01")
		if nil != err {
			assertCorrectMessage(t, "Failed to NewTreePairAlpha('01')", " in rangetp constructor in Multiply test.")
		}

		EncodeDFS(dTP, "{11110000111010000,11101000110100100,0 1 2 5 4 3 6 8 7}")
		EncodeDFS(rTP, "{11001101000,11101000100,5 1 2 4 0 3}")
		resultTP := Multiply(dTP, rTP)
		got := resultTP.FullString()
		want := "{D: [0000 0], [0001 1], [001 2], [01 3], [1000 4], [10010 5], [10011 6], [101 7], [11 8]" +
			" || R: " +
			"[0000 8], [0001 7], [0010 5], [0011 4], [01 6], [100 0], [1010 1], [1011 2], [11 3]}"
		assertCorrectMessage(t, got, want)
	})

	// InF false tests we can recognise the element is not in R. Thompson's group F
	t.Run("InF false", func(t *testing.T) {
		//reduces element to minimal tree pair.
		// makes permutation 1 2 3 ... 8
		tp, err := NewTreePairAlpha("01")
		if nil != err {
			assertCorrectMessage(t, "Failed to NewTreePairAlpha('01')", " in InF false test.")
		}
		EncodeDFS(tp, "{1110000,1010100,0 2 1 3}")
		got := strconv.FormatBool(tp.InF())
		want := strconv.FormatBool(false)
		assertCorrectMessage(t, got, want)
	})

	// InF true tests we can recognise the element is in R. Thompson's group F
	// even if a permtutation has been applied consistently on domain and range.
	t.Run("InF true", func(t *testing.T) {
		//reduces element to minimal tree pair.
		// makes permutation 1 2 3 ... 8
		tp, err := NewTreePairAlpha("01")
		if nil != err {
			assertCorrectMessage(t, "Failed to NewTreePairAlpha('01')", " in InF true test.")
		}

		EncodeDFS(tp, "{1110000,1010100,0 1 2 3}")
		tp.PermuteLabels(map[int]int{0: 1, 1: 2, 2: 3, 3: 0})
		got := strconv.FormatBool(tp.InF())
		want := strconv.FormatBool(true)
		assertCorrectMessage(t, got, want)
	})

	// InT false checks we can verify that an elt is not in T.
	t.Run("InT false", func(t *testing.T) {
		//reduces element to minimal tree pair.
		// makes permutation 1 2 3 ... 8
		tp, err := NewTreePairAlpha("01")
		if nil != err {
			assertCorrectMessage(t, "Failed to NewTreePairAlpha('01')", " in InT false test.")
		}

		EncodeDFS(tp, "{1110000,1010100,0 1 2 3}")
		tp.ApplyPermRange(map[int]int{0: 1, 1: 3, 2: 2, 3: 0})
		got := strconv.FormatBool(tp.InT())
		want := strconv.FormatBool(false)
		assertCorrectMessage(t, got, want)
	})

	// InT true checks we can verify that an elt is in T.
	t.Run("InT true", func(t *testing.T) {
		//reduces element to minimal tree pair.
		// makes permutation 1 2 3 ... 8
		tp, err := NewTreePairAlpha("01")
		if nil != err {
			assertCorrectMessage(t, "Failed to NewTreePairAlpha('01')", " in InT true test.")
		}

		EncodeDFS(tp, "{1110000,1010100,0 1 2 3}")
		tp.ApplyPermRange(map[int]int{0: 1, 1: 2, 2: 3, 3: 0})
		got := strconv.FormatBool(tp.InT())
		want := strconv.FormatBool(true)
		assertCorrectMessage(t, got, want)
	})

	// InT true checks we can verify that an elt is in T.
	// even if for this elt a weird permutation has been applied
	// consistently in domain and range.
	t.Run("InT true tough", func(t *testing.T) {
		//reduces element to minimal tree pair.
		// makes permutation 1 2 3 ... 8
		tp, err := NewTreePairAlpha("01")
		if nil != err {
			assertCorrectMessage(t, "Failed to NewTreePairAlpha('01')", " in InT true tough test.")
		}

		EncodeDFS(tp, "{1110000,1010100,0 1 2 3}")
		tp.ApplyPermRange(map[int]int{0: 1, 1: 2, 2: 3, 3: 0})
		tp.PermuteLabels(map[int]int{0: 1, 1: 3, 2: 2, 3: 0})
		got := strconv.FormatBool(tp.InT())
		want := strconv.FormatBool(true)
		assertCorrectMessage(t, got, want)
	})

	// InV true checks we can verify that an elt is in V.
	// since currently everything is, there is not much to do.
	t.Run("InV true", func(t *testing.T) {
		//reduces element to minimal tree pair.
		// makes permutation 1 2 3 ... 8
		tp, err := NewTreePairAlpha("01")
		if nil != err {
			assertCorrectMessage(t, "Failed to NewTreePairAlpha('01')", " in InV test.")
		}
		EncodeDFS(tp, "{1110000,1010100,0 1 2 3}")
		tp.ApplyPermRange(map[int]int{0: 1, 1: 3, 2: 2, 3: 0})
		got := strconv.FormatBool(tp.InV())
		want := strconv.FormatBool(true)
		assertCorrectMessage(t, got, want)
	})

	t.Run("LessEqual test", func(t *testing.T) {
		dTP, err := NewTreePairAlpha("01")
		if nil != err {
			assertCorrectMessage(t, "Failed to NewTreePairAlpha('01')", " in domaintp constructor in Multiply test.")
		}

		rTP, err := NewTreePairAlpha("01")
		if nil != err {
			assertCorrectMessage(t, "Failed to NewTreePairAlpha('01')", " in rangetp constructor in Multiply test.")
		}

		EncodeDFS(dTP, "{11001101000,11101000100,5 1 2 4 0 3}")
		EncodeDFS(rTP, "{11110000111010000,11101000110100100,0 1 2 5 4 3 6 8 7}")

		assert.True(t, LessEqual(*dTP, *rTP), "dTp <= rTP failed to be true.")

		assert.True(t, LessEqual(*dTP, *dTP), "dTP <= dTP failed to be true.")

		assert.False(t, LessEqual(*rTP, *dTP), "rTP was not greater than dTP")
	})
}
