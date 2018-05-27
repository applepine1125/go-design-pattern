package adapter

import (
	"bytes"
	"testing"
)

func TestDelegatePrintWeak(t *testing.T) {
	p := DelegatePrintBanner{&DelegateBanner{"Hello"}}
	buf := &bytes.Buffer{}
	p.printWeak(buf)

	correctStr := "(Hello)"
	outputStr := buf.String()

	if outputStr != correctStr {
		t.Errorf("Expect string is %s, but get %s", correctStr, outputStr)
	}

}

func TestDelegatePrintStrong(t *testing.T) {
	p := DelegatePrintBanner{&DelegateBanner{"Hello"}}
	buf := &bytes.Buffer{}
	p.printStrong(buf)

	correctStr := "*Hello*"
	outputStr := buf.String()

	if outputStr != correctStr {
		t.Errorf("Expect string is %s, but get %s", correctStr, outputStr)
	}

}
