package lexorank

var (
	MinChar = byte('0')
	MaxChar = byte('z')
)

func Rank(prev, next string) (string, bool) {
	if prev == "" {
		prev = string(MinChar)
	}
	if next == "" {
		next = string(MaxChar)
	}

	rank := ""
	i := 0

	for {
		prevChar := getChar(prev, i, MinChar)
		nextChar := getChar(next, i, MaxChar)

		if prevChar == nextChar {
			rank += string(prevChar)
			i++
			continue
		}

		midChar := mid(prevChar, nextChar)
		if midChar == prevChar || midChar == nextChar {
			rank += string(prevChar)
			i++
			continue
		}

		rank += string(midChar)
		break
	}

	if rank >= next {
		return prev, false
	}

	return rank, true
}

func mid(prev, next byte) byte {
	return (prev + next) / 2
}

func getChar(s string, i int, defaultChar byte) byte {
	if i >= len(s) {
		return defaultChar
	}
	return s[i]
}
