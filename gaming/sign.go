package gaming

const (
	NO_SIGN = iota
	X_SIGN  = iota
	O_SIGN  = iota
)

func GetSignChar(sign int) rune {
	switch sign {
	case NO_SIGN:
		return '.'
	case X_SIGN:
		return 'X'
	case O_SIGN:
		return 'O'
	}
	return '?'
}
