package treepair

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/loeksnokes/prefcode"
)

/*
TreePair An interface built specifically for the type treePair.

TreePair::

Has:
Two prefix codes with the same base alphabet and same cardinality.  One
called the range and one called the domain.  Either side can carry a permutation,
as the PrefixCode interface inherently assocites a bijection from the code to a set
{0, 1, ..., k-1} while the words in the prefix code have a natural order.

Can:
 1. Expand both trees at corresponding location in domain or range.
 2. Permute domain or range permutation
 3. ResetLabels sets the
    domain tree labels as 0 1 ... k-1 corresponding to its dictionary order and
    modifies range tree so the initial prefix map is unchanged.
 4. Minimise and Minimize (same function but for British English or American English.)
 5. Multiply tree pairs based off of same alphabet.
 6. Invert an element.
 7. Detect if the element is in F, T, or V.
 8. Initialise (trivial permutation elt) from a list of expansions in D/R
 9. Initialise from DFS notation representation string: e.g. "{11000,10100,1 2 0}"" (an elt of T)
 10. Initialise from Full representation string: e.g. "D: [00 0], [01 1], [1 2], R: [0 1], [10 2], [11 0]"
 11. Return domain/range permutations (natural permutation from prefix code in
    dictionary order to the numeric labels of leaves)
*/
type TreePair interface {
	Alphabet() []rune
	ApplyPermDomain(perm map[int]int) bool
	ApplyPermRange(perm map[int]int) bool
	CodeDomain() prefcode.PrefCode
	CodeRange() prefcode.PrefCode
	Equals(tp *TreePair) bool
	ExpandRangeAt(s string)
	ExpandDomainAt(s string)
	ExposedCarets() []string
	FullString() string
	InF() bool
	InT() bool
	InV() bool
	Invert()
	Minimise()
	Minimize()
	PermuteLabels(perm map[int]int) bool
	ResetLabels() bool
	ReduceDomainAt(s string) bool
	ReduceRangeAt(s string) bool
	Size() int
	SwapPermAtRangeKeys(a, b string) bool
	SwapPermAtDomainKeys(a, b string) bool
	// DFSString() string
}

type treePair struct {
	alphabet []rune
	dom      prefcode.PrefCode
	ran      prefcode.PrefCode
}

// NewTreePairAlpha returns a treepair as a TreePair and sets alphabet of runes by input string.
func NewTreePairAlpha(alphaStr string) (*treePair, error) {
	dpc, errd := prefcode.NewPrefCodeAlphaString(alphaStr)
	rpc, errr := prefcode.NewPrefCodeAlphaString(alphaStr)
	if nil != errd {
		outStr := "NewTreePairAlpha(): Failed to create domaintree from " + alphaStr
		fmt.Println(outStr)
		return nil, errd
	}
	if nil != errr {
		outStr := "NewTreePairAlpha(): Failed to create rangetree from " + alphaStr
		fmt.Println(outStr)
		return nil, errr
	}
	return &treePair{alphabet: prefcode.StringToRuneSlice(alphaStr),
		dom: dpc,
		ran: rpc}, nil
}

// EncodeDFS returns a treepair from an alphabet string (like "01") and a DFS string like
// "{11000,10100,1 2 0}".  In this case, the Depth-First-Search encoding gives the tree with
// leaves {00,01,1} in domain and leaves {0,10,11} in range.  That is,
// 11000 encodes the shape of the domain tree and 10100 encodes the shape of the range tree,
// "1 2 0" is the permutation applied to labels of range to where to send the domain leaves in the prefix map:
// 00 -> 11
// 01 -> 0
// 1 -> 10
// in this example.  Code verifies that the DFS strings work for alphabet cardinality along the way.
func EncodeDFS(tp TreePair, DFS string) bool {

	//fmt.Println("Encode DFS: " + DFS)
	s := strings.Split(DFS, ",")
	//a do nothing tree pair since the DFS was poorly formatted.
	if len(s) != 3 {
		fmt.Println(DFS + " did not have three fields between commas.")
		return false
	}
	if !strings.HasPrefix(s[0], "{") || !strings.HasSuffix(s[2], "}") {
		fmt.Println(DFS + " did not have first field starting with `{`." +
			"or final field did not end with `}`.")
		return false
	}
	s[0] = strings.TrimPrefix(s[0], "{")
	s[2] = strings.TrimSuffix(s[2], "}")

	//fmt.Println("Encode DFS: s[0]: " + s[0])
	//fmt.Println("Encode DFS: s[1]: " + s[1])
	//fmt.Println("tp.FullString():" + tp.FullString())

	alphaSize := len(tp.Alphabet())
	if !prefcode.ValidDFSForPrefC(alphaSize, s[0]) ||
		!prefcode.ValidDFSForPrefC(alphaSize, s[1]) {
		return false
	}
	//fmt.Println("Encode DFS: Valid codes!")

	if !prefcode.DFSToPrefCode(tp.CodeDomain(), s[0]) {
		return false
	}
	//fmt.Println("Encoded Domain code")
	//fmt.Println("Resulting tp: " + tp.FullString())
	if !prefcode.DFSToPrefCode(tp.CodeRange(), s[1]) {
		return false
	}
	//fmt.Println("Encoded Range code")
	//fmt.Println("Resulting tp: " + tp.FullString())

	perm := make(map[int]int, (len(s[2])+1)/2)
	permNumStrings := strings.Split(s[2], " ")

	//fmt.Println("Encoding permutation at range: " + s[2])

	//apply permutation to range from DFSString
	for k, v := range permNumStrings {
		pv, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("NewTreePair DFS: bad perm conversion")
			return false
		}
		perm[k] = pv
	}
	tp.ApplyPermRange(perm)
	//fmt.Println("Resulting tp: " + tp.FullString())
	return true
}

// returns a ptr to a copy of the alphabet runes.
func (tp treePair) Alphabet() []rune {
	retVal := tp.dom.Alphabet()
	return retVal
}

// CodeDomain returns a ptr to the prefcode in domain
func (tp treePair) CodeDomain() prefcode.PrefCode {
	return tp.dom
}

// CodeRange  returns a ptr to the prefcode in range
func (tp treePair) CodeRange() prefcode.PrefCode {
	return tp.ran
}

func (tp treePair) FullString() (fullString string) {
	fullString = "{D: " + tp.dom.String() + " || R: " + tp.ran.String() + "}"
	return
}

// Equals compares a treepair to an input treepair as formal combinatorial objects.
// It is not a comparison of maps.  For that, one should minimise both tree pairs first.
func (tp treePair) Equals(tpp *TreePair) bool {
	return tp.FullString() == (*tpp).FullString()
}

// ApplyPermDomain acts by permutation on labels of domain tree
func (tp treePair) ApplyPermDomain(perm map[int]int) bool {
	return tp.dom.ApplyPerm(perm)
}

// ApplyPermRange acts by permutation on labels of range tree
func (tp treePair) ApplyPermRange(perm map[int]int) bool {
	return tp.ran.ApplyPerm(perm)
}

// PermuteLabels acts by same permutation on labels of domain and range tree
func (tp treePair) PermuteLabels(perm map[int]int) bool {
	domSuccess := tp.ApplyPermDomain(perm)
	ranSuccess := tp.ApplyPermRange(perm)
	return domSuccess && ranSuccess
}

// ResetLabels applies the same permutation to the labels of domain and range tree
// so that the resulting permutation on domain tree corresponds to the natural
// dictionary order on that prefix code.
func (tp treePair) ResetLabels() bool {
	currentPerm := tp.dom.Permutation()
	permSize := len(currentPerm)
	inversePerm := make(map[int]int, permSize)

	for k, v := range currentPerm {
		inversePerm[v] = k
	}

	return tp.PermuteLabels(inversePerm)
}

// Invert returns the inverse tree-pair element.  Labels are not reset.
func (tp *treePair) Invert() {
	tp.dom, tp.ran = tp.ran, tp.dom
}

// InF assesses if elmt is in R. Thompson's group F
// does not relabel the element
func (tp *treePair) InF() bool {
	domainPerm := (*tp).CodeDomain().Permutation()
	rangePerm := (*tp).CodeRange().Permutation()
	lrp := len(rangePerm)

	for k := 0; k < lrp; k++ {
		if rangePerm[k] != domainPerm[k] {
			return false
		}
	}
	return true
}

// InF assesses if elmt is in R. Thompson's group F
// does not relabel the element
func (tp *treePair) InT() bool {
	domainPerm := (*tp).CodeDomain().Permutation()
	rangePerm := (*tp).CodeRange().Permutation()
	lrp := len(rangePerm)

	//makes a double copy of rangePerm
	doubleRange := make(map[int]int, 2*lrp)
	for k, v := range rangePerm {
		doubleRange[k] = v
		doubleRange[k+lrp] = v
	}

	//fmt.Println("InT(): tp: " + (*tp).FullString())
	//fmt.Println("Int(): doubleRange == " + strconv.Itoa(len(doubleRange)))

	// checks if there is contiguous subslice in doubleRange
	// that looks like domainPerm: same thing as rangePerm being
	// a rotation of domainperm
	lenDR := len(doubleRange)
	firstVal := domainPerm[0]
	startFound := false
	startSpot := 2 * lrp // This value will be out-of-bounds but not checked.  Crash == bug in code!
	for k := 0; k < lenDR; k++ {
		//	fmt.Println("InT(): k==" + strconv.Itoa(k) +
		//		", doubleRange[k]==" + strconv.Itoa(doubleRange[k]) +
		//		", firstVal==" + strconv.Itoa(firstVal))

		if !startFound && doubleRange[k] == firstVal { //found start of possible domain sequence
			startSpot = k
			startFound = true
			//		fmt.Println("InT(): found start: k==" + strconv.Itoa(k) + " firstVal==" + strconv.Itoa(k) + "doubleRange[k]==" + strconv.Itoa(doubleRange[k]))
		}
		if startFound && k < (startSpot+lrp) {
			//		fmt.Println("InT(): k==" + strconv.Itoa(k) +
			//		", domainPerm[k-startSpot]==" + strconv.Itoa(domainPerm[k-startSpot]) +
			//		"doubleRange[k]==" + strconv.Itoa(doubleRange[k]))

			if domainPerm[k-startSpot] != doubleRange[k] {
				// doubleRange stopped copying domainPerm prematurely
				//fmt.Println("InT(): k==" + strconv.Itoa(k) +
				//", domainPerm[k-startSpot]==" +
				//strconv.Itoa(domainPerm[k-startSpot]) +
				//"doubleRange[k]==" + strconv.Itoa(doubleRange[k]) +
				//" so not in T.")
				return false
			}
		}
		//if we got here without returning, the elt is in T.
		if k == (startSpot + lrp - 1) {
			//fmt.Println("InT(): It all checked out!  InT true.")
			return true
		}
	}
	//fmt.Println("InT: we should never print this.")
	return false
}

// InV always is true.
func (tp *treePair) InV() bool { return true }

// ReduceDomainAt takes as input the root of a claimed exposed caret and reduces domain and
// range of both trees at corresponding carets IF the permutation labels are identical
// across the leaves of those carets.
// **In all cases has a side effect of resetting labels (even if no reduction is possible).**
// true if reduction occurred, false if it was not possible.
func (tp treePair) ReduceDomainAt(s string) bool {
	tp.ResetLabels()

	reductionSpots := tp.dom.ExposedCarets()

	sRootOfExposedCaret := false
	for _, v := range reductionSpots {
		if v == s {
			sRootOfExposedCaret = true
			break
		}
	}
	if !sRootOfExposedCaret {
		return false
	}

	// alphabet size is a, and labels are in order.  Check if the corresponding leaves in
	// range are the leaves of an exposed caret.
	firstLeaf := s + string(tp.alphabet[0])
	leftLeafLabelDomain := tp.dom.LabelAtLeaf(firstLeaf)
	firstImageLeaf := tp.ran.LeafAtLabel(leftLeafLabelDomain)

	if "" == firstImageLeaf {
		return false
	}

	rangeRoot := firstImageLeaf[:len(firstImageLeaf)-1]
	for k, v := range tp.alphabet {
		if (leftLeafLabelDomain + k) != tp.ran.LabelAtLeaf(rangeRoot+string(v)) {
			return false
		}
	}

	//if we got here, the same caret by labels is exposed in domain and range and we
	//can reduce both.  Further, the permutation labelling this caret is of the form
	// a,a+1,...,a+(alphaSize-1)

	//Payload!  Reduce on both sides!!
	tp.dom.ReduceAt(s)
	tp.ran.ReduceAt(rangeRoot)

	//reindex from domain tree (this should actually do nothing!)
	tp.ResetLabels()
	return true
}

// ReduceRangeAt reduces treepair tp at exposed caret in range if the preimage set is
// also an exposed caret of leaves listed with same corresponding labels.
func (tp treePair) ReduceRangeAt(s string) bool {
	tp.Invert()
	wasReduced := tp.ReduceDomainAt(s)
	tp.Invert()
	tp.ResetLabels()
	return wasReduced
}

// ExpandDomainAt at string s:  if s is deeper than domain prefix code, the domain prefix
// code is expanded minimally so that s becomes a root of an exposed caret.  The range tree and
// permutations are expanded correspondingly.  If s is shallower than leaves of Domain tree
// then nothing happens.
func (tp treePair) ExpandDomainAt(s string) {
	prefixLeaf := tp.dom.GetPrefixOf(s)
	lenPref := len(prefixLeaf)

	// s was too shallow
	if "" == prefixLeaf && prefcode.EmptyString != tp.dom.LeafAtLabel(0) {
		return
	}

	suffix := s[lenPref:]

	permValue := tp.dom.LabelAtLeaf(prefixLeaf)
	newPrefix := tp.ran.LeafAtLabel(permValue)

	ranExpandPt := newPrefix + suffix

	tp.dom.ExpandAt(s)
	tp.ran.ExpandAt(ranExpandPt)

	return
}

// ExpandRanAt expands the treepair if s is  a leaf or is deeper that the range tree code.
func (tp treePair) ExpandRangeAt(s string) {
	tp.Invert()
	tp.ExpandDomainAt(s)
	tp.Invert()
	return
}

// Multiply returns a new TreePair that is the product of the two that are fed in.
func Multiply(nonLocalFirst, nonLocalSecond TreePair) *treePair {

	first := treePair{alphabet: nonLocalFirst.Alphabet(), dom: nonLocalFirst.CodeDomain(), ran: nonLocalFirst.CodeRange()}
	second := treePair{alphabet: nonLocalSecond.Alphabet(), dom: nonLocalSecond.CodeDomain(), ran: nonLocalSecond.CodeRange()}

	fmt.Println("first: " + first.FullString())
	fmt.Println("second: " + second.FullString())
	fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")
	fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")
	//build steady labelling
	first.ResetLabels()
	second.ResetLabels()

	// Make a prefix code that is join of range of first element and domain of second element
	fmt.Println("First Range: " + first.CodeRange().String())
	fmt.Println("Second Domain: " + second.CodeDomain().String())
	fullCode, err := first.CodeRange().Join(second.CodeDomain())
	if nil != err {
		panic("Multiply(): err return for join")
	}
	fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")
	fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")
	fmt.Println("Join D-R code: " + fullCode.String())

	//for each leaf of the join tree, force it to be a leaf in range first/domain second
	for key, _ := range fullCode.Code() {
		first.ExpandRangeAt(key)
		second.ExpandDomainAt(key)
	}

	/*
		// This commented out code tries to be efficient by only adding the exposed carets but it
		// does not seem to wrok properly.
		// Get the exposed carets we can use to expand our two elements.


			exposed := fullCode.ExposedCarets()

			// Expand first and second so range of first = domain of second = join we found.
			for _, v := range exposed {
				first.ExpandRangeAt(v)
				second.ExpandDomainAt(v)
			}*/

	fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")
	fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")
	fmt.Println("Expanded treepairs")
	fmt.Println("first: " + first.FullString())
	fmt.Println("second: " + second.FullString())

	// align the permutation of domain of second element to the permutation on range of first element.
	second.PermuteLabels(first.CodeRange().Permutation())

	// return a new treepair with the correct domain, range, and permutation.
	return &treePair{alphabet: first.Alphabet(), dom: first.CodeDomain(), ran: second.CodeRange()}
}

func Power(first TreePair, pow int) *treePair {
	if pow == 0 {
		// return the identity in a way that multiplies easily with previous
		return &treePair{alphabet: first.Alphabet(), dom: first.CodeRange(), ran: first.CodeRange()}
	}
	if pow < 0 {
		first.Invert()
		pow *= -1
	}
	first.Minimise()
	return Multiply(first, Power(first, pow-1))
}

// Minimise reduces a tree-pair.  Even if no reductions
// are possible, the labels will be reset (domain tree labels
// will appear in natural order)
func (tp treePair) Minimise() {
	domExposed := tp.dom.ExposedCarets()

	madeReduction := false
	for _, v := range domExposed {
		if tp.ReduceDomainAt(v) {
			madeReduction = true
		}
	}
	if madeReduction { // if reductions occurred, new reductions can become possible.
		tp.Minimise()
	}
	return
}

// Minimize This does Minimise, but For American English spellers
func (tp treePair) Minimize() {
	tp.Minimise()
	return
}

func (tp treePair) SwapPermAtRangeKeys(a, b string) bool  { return true }
func (tp treePair) SwapPermAtDomainKeys(a, b string) bool { return true }

// NewTreePairDFS(s string)
func (tp treePair) ExposedCarets() []string { return tp.dom.ExposedCarets() }
func (tp treePair) Size() int               { return tp.dom.Size() }
func (tp treePair) DFSString() string       { return "Stuff" }

func badSpeed(DFS string, cap int) (fast bool) {
	fast = false
	strLen := len(DFS)
	//Empty string DFS is not allowed: returned as too Fast.
	if strLen < 1 {
		fmt.Println("badSpeed(): Tree description by DFS cannot be empty.")
		return true
	}
	stackHeight := 1
	limit := cap - 1
	for ii, v := range DFS {
		if `1` == string(v) {
			stackHeight = stackHeight + cap - 1
			continue
		}
		if `0` == string(v) { //certainly the case
			stackHeight = stackHeight - 1
			if 0 == stackHeight && ii < limit {
				fmt.Println("badSpeed(): Tree description by DFS cannot be empty.")
				return true
			}
		}
	}
	if 0 == stackHeight {
		return false
	}
	fmt.Println("badSpeed(): Tree description by DFS cannot have too many `1`'s.")
	return true

}

// Checks if A is less than or equal to B as tree-pairs.
// Dictionary order on pair (size of domain tree, Full Description String)
func LessEqual(tpA treePair, tpB treePair) bool {
	if tpA.dom.Size() < tpB.dom.Size() {
		return true
	}
	AString := tpA.FullString()
	BString := tpB.FullString()
	return AString <= BString
}
