package main

import (
	"time"
)

var wordSize = 1

type view interface {
	Width() int
}

// type pr represents a positioned rune
type pr struct {
	r      rune
	offset int
}

type prs []*pr

func (prs prs) consume() {
	for i, _ := range prs {
		if i < len(prs)-1 {
			prs[i] = prs[i+1]
		}
	}
	if prs[0].r == ' ' {
		prs.consume()
	}
}

func (prs prs) produce() {
	lastWordSize := 0
	pr := &pr{}

	for i := len(prs) - 2; i > 0 && prs[i].r != ' '; i-- {
		lastWordSize++
	}

	if lastWordSize == wordSize {
		pr.r = ' '
	} else {
		pr.r = 'a'
	}

	pr.offset = prs[len(prs)-2].offset + 1

	prs[len(prs)-1] = pr
}

func (prs prs) advance() {
	for i, r := range prs {
		if i == 0 {
			r.offset = 0
		} else {
			r.offset--
		}
	}
}

func setup(width int) prs {
	prs := make(prs, 100)
	for i, _ := range prs {
		prs[i] = &pr{
			r:      rune('a' + i),
			offset: width + 1,
		}
	}
	return prs
}

func main() {

	//  width := view.Width()
	in := make(chan rune)
	sync := make(chan prs)
	prs := setup(100)
	sync <- prs

	go listen(in, sync)

	go tick(2000, sync)

}

func listen(in chan rune, sync chan prs) {
	i := <-in
	prs := <-sync

	if i != prs[0].r {
		prs.advance()
		return
	}

	prs.consume()
	prs.produce()
	prs.advance()
	sync <- prs
}

func tick(interval int, sync chan prs) {
	for {
		time.Sleep(time.Duration(interval) * time.Millisecond)
		prs := <-sync
		prs.advance()
		sync <- prs
	}
}
