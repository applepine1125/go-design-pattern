package prototype

import (
	"bytes"
	"fmt"
)

type Product interface {
	Use(s string) string
	createClone() Product
}

type Manager struct {
	showcase map[string]Product
}

func (m *Manager) register(name string, proto Product) {
	m.showcase[name] = proto
}

func (m *Manager) create(name string) Product {
	return m.showcase[name].createClone()
}

type MessageBox struct {
	decochar string
}

func (mb *MessageBox) Use(s string) string {
	length := len(s)
	w := &bytes.Buffer{}

	for i := 0; i < length+4; i++ {
		fmt.Fprint(w, mb.decochar)
	}
	fmt.Fprint(w, "\n")
	fmt.Fprint(w, mb.decochar+" "+s+" "+mb.decochar+"\n")
	for i := 0; i < length+4; i++ {
		fmt.Fprint(w, mb.decochar)
	}
	return w.String()
}

func (mb *MessageBox) createClone() Product {
	return &MessageBox{decochar: mb.decochar}
}

type Underline struct {
	ulchar string
}

func (u *Underline) Use(s string) string {
	length := len(s)
	w := &bytes.Buffer{}

	fmt.Fprint(w, "\""+s+"\"\n ")
	for i := 0; i < length; i++ {
		fmt.Fprint(w, u.ulchar)
	}
	return w.String()
}

func (u *Underline) createClone() Product {
	return &Underline{ulchar: u.ulchar}
}
