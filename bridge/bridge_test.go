package bridge

import (
	"testing"
)

func TestBridge(t *testing.T) {
	d := Display{&StringDisplayImpl{"Hello world."}}
	expect := "+------------+\n|Hello world.|\n+------------+\n"

	output := d.display()
	if output != expect {
		t.Errorf("Expect output to %s, but %s\n", expect, output)
	}

	cd := CountDisplay{&Display{&StringDisplayImpl{"Hello world."}}}
	expect = "+------------+\n|Hello world.|\n+------------+\n"

	output = cd.display.display()
	if output != expect {
		t.Errorf("Expect output to %s, but %s\n", expect, output)
	}

	expect = "+------------+\n|Hello world.|\n|Hello world.|\n+------------+\n"

	output = cd.mulitDisplay(2)
	if output != expect {
		t.Errorf("Expect output to %s, but %s\n", expect, output)
	}
}
