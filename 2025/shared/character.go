package shared

type Character rune

const (
	Newline        rune = '\n'
	CarriageReturn rune = '\r'
	Space          rune = ' '
)

func (c Character) String() string {
	return string(c)
}

func (c Character) Rune() rune {
	return rune(c)
}
