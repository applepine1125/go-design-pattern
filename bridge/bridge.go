package bridge

type Display struct {
	impl DisplayImpl
}

func (d *Display) open() string {
	return d.impl.rawOpen()
}

func (d *Display) print() string {
	return d.impl.rawPrint()
}

func (d *Display) close() string {
	return d.impl.rawClose()
}

func (d *Display) display() string {
	return d.impl.rawOpen() + d.impl.rawPrint() + d.impl.rawClose()
}

type CountDisplay struct {
	display *Display
}

func (cd *CountDisplay) mulitDisplay(times int) string {
	str := cd.display.open()

	for i := 0; i < times; i++ {
		str += cd.display.print()
	}

	str += cd.display.close()

	return str
}

type DisplayImpl interface {
	rawOpen() string
	rawPrint() string
	rawClose() string
}

type StringDisplayImpl struct {
	str string
}

func (sd *StringDisplayImpl) rawOpen() string {
	return sd.printLine()
}

func (sd *StringDisplayImpl) rawPrint() string {
	return "|" + sd.str + "|\n"
}

func (sd *StringDisplayImpl) rawClose() string {
	return sd.printLine()
}

func (sd *StringDisplayImpl) printLine() string {
	str := "+"
	for i := 0; i < len(sd.str); i++ {
		str += "-"
	}
	str += "+\n"
	return str
}
