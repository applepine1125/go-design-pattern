package builder

import (
	"fmt"
	"io"
)

type Builder interface {
	makeTitle(title string)
	makeString(str string)
	makeItems(items []string)
	close()
}

type Director struct {
	builder Builder
}

func (d *Director) consrtuct() {
	d.builder.makeTitle("Greeting")
	d.builder.makeString("at morning to noon")
	d.builder.makeItems([]string{
		"Good morning",
		"Good afternoon",
	})
	d.builder.makeString("at evening")
	d.builder.makeItems([]string{
		"Good evening",
		"Good night",
		"Good bye",
	})

	d.builder.close()

}

type TextBuilder struct {
	w io.Writer
}

func (tb *TextBuilder) makeTitle(title string) {
	fmt.Fprint(tb.w, "====================\n")
	fmt.Fprint(tb.w, "["+title+"]\n\n")
}

func (tb *TextBuilder) makeString(str string) {
	fmt.Fprint(tb.w, "* "+str+"\n\n")
}

func (tb *TextBuilder) makeItems(items []string) {
	for i := 0; i < len(items); i++ {
		fmt.Fprint(tb.w, "  - "+items[i]+"\n")
	}
	fmt.Fprint(tb.w, "\n")
}

func (tb *TextBuilder) close() {
	fmt.Fprint(tb.w, "====================\n")
}

type HTMLBuilder struct {
	w io.Writer
}

func (hb *HTMLBuilder) makeTitle(title string) {
	fmt.Fprint(hb.w, "<html><head><title>"+title+"</title></head><body>")
	fmt.Fprint(hb.w, "<h1>"+title+"</h1>")
}

func (hb *HTMLBuilder) makeString(str string) {
	fmt.Fprint(hb.w, "<p>"+str+"</p>")
}

func (hb *HTMLBuilder) makeItems(items []string) {
	fmt.Fprint(hb.w, "<ul>")
	for i := 0; i < len(items); i++ {
		fmt.Fprint(hb.w, "<li>"+items[i]+"</li>")
	}
	fmt.Fprint(hb.w, "</ul>")
}

func (hb *HTMLBuilder) close() {
	fmt.Fprint(hb.w, "</body></html>")
}
