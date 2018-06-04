package builder

import (
	"bytes"
	"testing"
)

func TestTextBuilder(t *testing.T) {
	buf := &bytes.Buffer{}
	director := Director{&TextBuilder{w: buf}}
	director.consrtuct()

	expected := `====================
[Greeting]

* at morning to noon

  - Good morning
  - Good afternoon

* at evening

  - Good evening
  - Good night
  - Good bye

====================
`
	result := buf.String()

	if expected != result {
		t.Errorf("expected value is \n%s\n, but got \n%s\n", expected, result)
	}
}

func TestHTMLBuilder(t *testing.T) {
	buf := &bytes.Buffer{}
	director := Director{&HTMLBuilder{w: buf}}
	director.consrtuct()

	expected := `<html><head><title>Greeting</title></head><body><h1>Greeting</h1><p>at morning to noon</p><ul><li>Good morning</li><li>Good afternoon</li></ul><p>at evening</p><ul><li>Good evening</li><li>Good night</li><li>Good bye</li></ul></body></html>`
	result := buf.String()

	if expected != result {
		t.Errorf("expected value is \n%s\n, but got \n%s\n", expected, result)
	}
}
