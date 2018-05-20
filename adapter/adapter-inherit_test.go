package adapter

import (
	"bytes"
	"testing"
)

func TestPrintWeak(t *testing.T) {
	p := InheritPrintBanner{&InheritBanner{"Hello"}}
	buf := &bytes.Buffer{}
	p.printWeak(buf)

	correctStr := "(Hello)"
	outputStr := buf.String()

	if outputStr != correctStr {
		t.Errorf("Expect string is %s, but get %s", correctStr, outputStr)
	}

}

func TestPrintStrong(t *testing.T) {
	p := InheritPrintBanner{&InheritBanner{"Hello"}}
	buf := &bytes.Buffer{}
	p.printStrong(buf)

	correctStr := "*Hello*"
	outputStr := buf.String()

	if outputStr != correctStr {
		t.Errorf("Expect string is %s, but get %s", correctStr, outputStr)
	}

}
