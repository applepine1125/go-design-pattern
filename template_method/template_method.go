package template_method

type AbstractDisplay struct {
}

type printer interface {
	open() string
	print() string
	close() string
}

func (ab *AbstractDisplay) display(printer printer) string {
	result := printer.open()
	for i := 0; i < 5; i++ {
		result += printer.print()
	}
	result += printer.close()
	return result
}

type CharDisplay struct {
	*AbstractDisplay
	char rune
}

func (cd *CharDisplay) open() string {
	return "<<"
}

func (cd *CharDisplay) print() string {
	return string(cd.char)
}

func (cd *CharDisplay) close() string {
	return ">>"
}

type StringDisplay struct {
	*AbstractDisplay
	str string
}

func (sd *StringDisplay) open() string {
	return sd.printLine()
}

func (sd *StringDisplay) print() string {
	return "| " + sd.str + " |\n"
}

func (sd *StringDisplay) close() string {
	return sd.printLine()
}

func (sd *StringDisplay) printLine() string {
	line := "+-"
	for _ = range sd.str {
		line += "-"
	}
	line += "-+\n"
	return line
}
