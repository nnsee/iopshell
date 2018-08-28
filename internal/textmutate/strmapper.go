package textmutate

import (
	"strings"
)

type parser struct {
	Input  string
	SChars []rune
	Cursor int
}

func (p *parser) specialChar(offset ...int) bool {
	cursor := p.Cursor
	if len(offset) > 0 {
		cursor += offset[0]
	}
	for _, char := range p.SChars {
		if rune(p.Input[cursor]) == char {
			return true
		}
	}
	return false
}

func (p *parser) word() string {
	var word strings.Builder

	for p.Cursor < len(p.Input) {
		if p.specialChar() {
			p.Cursor++
			return word.String()
		}
		word.WriteByte(p.Input[p.Cursor])
		p.Cursor++
	}
	return word.String()
}

func (p *parser) peekChar() rune {
	for i, char := range p.Input[p.Cursor:] {
		if p.specialChar(i) {
			return char
		}
	}
	return 0
}

func StrToMap(input string) (map[string]interface{}, int) {
	out := make(map[string]interface{})
	p := parser{Input: input,
		SChars: []rune{':', ';', ','}}
	for key := p.word(); key != ""; key = p.word() {
		switch char := p.peekChar(); char {
		case ':':
			newGen, cur := StrToMap(p.Input[p.Cursor:])
			out[key] = newGen
			p.Cursor += cur
		case ',':
			out[key] = p.word()
		case ';':
			out[key] = p.word()
			return out, p.Cursor
		default:
			out[key] = p.word()
			break
		}
	}
	return out, len(input)
}
