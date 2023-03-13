package popcount

import (
	"regexp"
	"strconv"
	"testing"
)

// TestPopCountLUTAll calls popcount.PopCountLUT with all bits 1 value,
// checking for a valid return value.
func TestPopCountLUTAll(t *testing.T) {
	var x uint64 = 0xffffffffffffffff
	want := regexp.MustCompile("64")
	cnt := PopCountLUT(x)
	if !want.MatchString(strconv.Itoa(cnt)) {
		t.Fatalf(`PopCountLUT(%#x) = %d, want match for %#q`, x, cnt, want)
	}
}

// TestPopCountLUTNone calls popcount.PopCountLUT with all bits 0 value,
// checking for a valid return value.
func TestPopCountLUTNone(t *testing.T) {
	var x uint64 = 0x0000000000000000
	want := regexp.MustCompile("0")
	cnt := PopCountLUT(x)
	if !want.MatchString(strconv.Itoa(cnt)) {
		t.Fatalf(`PopCountLUT(%#x) = %d, want match for %#q`, x, cnt, want)
	}
}

// TestPopCountRNZAll calls popcount.PopCountRNZ with all bits 1 value,
// checking for a valid return value.
func TestPopCountRNZAll(t *testing.T) {
	var x uint64 = 0xffffffffffffffff
	want := regexp.MustCompile("64")
	cnt := PopCountRNZ(x)
	if !want.MatchString(strconv.Itoa(cnt)) {
		t.Fatalf(`PopCountRNZ(%#x) = %d, want match for %#q`, x, cnt, want)
	}
}

// TestPopCountRNZNone calls popcount.PopCountRNZ will all bits 0 value,
// checking for a valid return value.
func TestPopCountRNZNone(t *testing.T) {
	var x uint64 = 0x0000000000000000
	want := regexp.MustCompile("0")
	cnt := PopCountRNZ(x)
	if !want.MatchString(strconv.Itoa(cnt)) {
		t.Fatalf(`PopCountRNZ(%#x) = %d, want match for %#q`, x, cnt, want)
	}
}

// TestPopCountAll calls all 3 popcount.PopCount_ functions with value from
// [0, 100000). Then compare them to each other.
func TestPopCountAll(t *testing.T) {
	var i uint64
	for i = 0; i < 100000; i++ {
		cntLUT := PopCountLUT(i)
		cntShift := PopCountShift(i)
		cntRNZ := PopCountRNZ(i)

		if cntLUT != cntShift || cntShift != cntRNZ {
			t.Fatalf(`PopCountLUT(%#x) = %d, PopCountShift(%#[1]x) = %d, PopCountRNZ(%#[1]x) = %d`,
				i, cntLUT, cntShift, cntRNZ)
		}
	}
}
