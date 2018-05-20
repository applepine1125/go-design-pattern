package adapter

import (
	"fmt"
	"io"
)

type InheritBanner struct {
	String string
}

type InheritPrint interface {
	PrintWeak()
	PrintStr0ng()
}

type InheritPrintBanner struct {
	*InheritBanner
}

func (ib *InheritBanner) showWithParen(w io.Writer) {
	fmt.Fprint(w, "("+ib.String+")")
}

func (ib *InheritBanner) showWithAster(w io.Writer) {
	fmt.Fprint(w, "*"+ib.String+"*")
}

func (pb *InheritPrintBanner) printWeak(w io.Writer) {
	pb.showWithParen(w)
}

func (pb *InheritPrintBanner) printStrong(w io.Writer) {
	pb.showWithAster(w)
}
