package main

import (
	"testing"
)

func aTestListen(t *testing.T) {
	in := make(chan rune)
	sync := make(chan prs)
	in <- 'b'
	listen(in, sync)

}

func TestConsume(t *testing.T) {
	prs := setup(100)
	want := prs[1].r
	prs.consume()
	got := prs[0].r
	if  got != want {
		t.Log(prs)
		t.Errorf("wanted '%c', but got '%c'", want, got)
	}
}

func TestProduce(t *testing.T) {
	wordSize = 3

	// convenience function to make our tests input more readable
	makePrs := func(s string) prs {
		prs := make(prs, len(s))
		for i, r := range s {
			prs[i] = &pr{r, 0}
		}
		return prs
	}

	// convenience function for producing and getting produced value
	produce := func(prs prs) *pr {
		prs.produce()
		return prs[len(prs)-1]
	}

	wantNonSpace := func(got *pr, t *testing.T) {
		if got.r == ' ' {
			t.Errorf("did not want a ' ', but got '%c'", got.r)
		}
	}

	prs := makePrs("ab abb")
	got := produce(prs)
	wantNonSpace(got, t)

	prs = makePrs("ab cc")
	got = produce(prs)
	wantNonSpace(got, t)

	prs = makePrs("ab cccc")
	got = produce(prs)
	if got.r != ' ' {
		t.Errorf("wanted ' ', but got '%c'", got.r)
	}

}

func TestAdvance(t *testing.T) {
	prs := prs{{'a', 10}, {'b', 11}, {'c', 12}}
	prs.advance()
	
	got := prs[0].offset
	if got != 0 {
		t.Errorf("wanted 0 but got %d", got)
	}
	
	got = prs[1].offset
	if got != 10 {
			t.Errorf("wanted 12 but got %d", got)
	}
	
		got = prs[2].offset
	if got != 11 {
			t.Errorf("wanted 13 but got %d", got)
	}
}
