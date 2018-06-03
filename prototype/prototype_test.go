package prototype

import (
	"testing"
)

func TestPrototype(t *testing.T) {

	expect1 := "*********\n* Hello *\n*********"
	expect2 := "---------\n- Hello -\n---------"
	expect3 := "\"Hello\"\n ~~~~~"

	manager := Manager{showcase: map[string]Product{}}

	asteriskbox := &MessageBox{decochar: "*"}
	manager.register("*", asteriskbox)
	clone1 := manager.create("*")

	if asteriskbox.Use("Hello") != expect1 {
		t.Errorf("Expect string is \n%s\n but get \n%s", expect1, asteriskbox.Use("Hello"))
	}

	if asteriskbox.Use("Hello") != clone1.Use("Hello") {
		t.Errorf("Expect instances are equal, but not.")
	}

	minusbox := &MessageBox{decochar: "-"}
	manager.register("-", minusbox)
	clone2 := manager.create("-")

	if minusbox.Use("Hello") != expect2 {
		t.Errorf("Expect string is \n%s\n but get \n%s", expect2, minusbox.Use("Hello"))
	}
	if minusbox.Use("Hello") != clone2.Use("Hello") {
		t.Errorf("Expect instances are equal, but not.")
	}

	underline := &Underline{ulchar: "~"}
	manager.register("~", underline)
	clone3 := manager.create("~")

	if underline.Use("Hello") != expect3 {
		t.Errorf("Expect string is \n%s\n but get \n%s", expect3, underline.Use("Hello"))
	}
	if underline.Use("Hello") != clone3.Use("Hello") {
		t.Errorf("Expect instances are equal, but not.")
	}

}
