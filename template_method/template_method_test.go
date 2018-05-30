package template_method

import (
	"testing"
)

func TestCharDisplay(t *testing.T) {
	charDisplay := &CharDisplay{char: 'H'}

	assert := "<<HHHHH>>"
	result := charDisplay.display(charDisplay)

	if assert != result {
		t.Errorf("Expect book.name is %s, but get %s", assert, result)
	}
}

func TestStringDisplay(t *testing.T) {
	strDisplay := &StringDisplay{str: "ABCDE"}

	assert := "+-------+\n| ABCDE |\n| ABCDE |\n| ABCDE |\n| ABCDE |\n| ABCDE |\n+-------+\n"
	result := strDisplay.display(strDisplay)

	if assert != result {
		t.Errorf("Expect book.name is %s, but get %s", assert, result)
	}
}
