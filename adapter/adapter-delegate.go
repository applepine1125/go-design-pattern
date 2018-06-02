package adapter

import (
	"fmt"
	"io"
)

type DelegateBanner struct {
	String string
}

type DelegatePrint interface {
	PrintWeak()
	PrintStrong()
}

type DelegatePrintBanner struct {
	delegateBanner *DelegateBanner
}

func (ib *DelegateBanner) showWithParen(w io.Writer) {
	fmt.Fprint(w, "("+ib.String+")")
}

func (ib *DelegateBanner) showWithAster(w io.Writer) {
	fmt.Fprint(w, "*"+ib.String+"*")
}

func (pb *DelegatePrintBanner) printWeak(w io.Writer) {
	pb.delegateBanner.showWithParen(w)
}

func (pb *DelegatePrintBanner) printStrong(w io.Writer) {
	pb.delegateBanner.showWithAster(w)
}
