package parse

import (
	"bufio"
	"io"
	"strings"
)

func removeComments(s string) string {
	r := bufio.NewReader(strings.NewReader(s))
	var result []rune
	for {
		ru, _, err := r.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		switch ru {
		case '\'', '"':
			result = append(result, ru)
			result = append(result, readString(r, ru == '\'')...)
		case '-':
			peeked, err := r.Peek(2)
			if err == nil &&
				len(peeked) == 2 &&
				rune(peeked[0]) == '-' &&
				rune(peeked[1]) == ' ' {
				discardUntilEOL(r)
			} else {
				result = append(result, ru)
			}
		case '/':
			peeked, err := r.Peek(1)
			if err == nil &&
				len(peeked) == 1 &&
				rune(peeked[0]) == '*' {
				// read the char we peeked
				_, _, _ = r.ReadRune()
				discardMultilineComment(r)
			} else {
				result = append(result, ru)
			}
		default:
			result = append(result, ru)
		}
	}
	return string(result)
}

func discardUntilEOL(r *bufio.Reader) {
	for {
		ru, _, err := r.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		if ru == '\n' {
			break
		}
	}
}
func discardMultilineComment(r *bufio.Reader) {
	for {
		ru, _, err := r.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		if ru == '*' {
			peeked, err := r.Peek(1)
			if err == nil && len(peeked) == 1 && rune(peeked[0]) == '/' {
				// read the rune we just peeked
				_, _, _ = r.ReadRune()
				break
			}
		}
	}
}

func readString(r *bufio.Reader, single bool) []rune {
	var result []rune
	var escaped bool
	for {
		ru, _, err := r.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		result = append(result, ru)
		if (!single && ru == '"' && !escaped) ||
			(single && ru == '\'' && !escaped) {
			break
		}
		escaped = false
		if ru == '\\' {
			escaped = true
		}
	}
	return result
}
