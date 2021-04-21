/*
   Copyright (c) 2021 Rasmus Moorats (neonsea)

   This file is part of iopshell.

   iopshell is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   iopshell is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with iopshell. If not, see <https://www.gnu.org/licenses/>.
*/

// If you do not wish to inflict pain upon yourself, turn back now.

package textmutate

import (
	"strings"
)

type parser struct {
	Input   string
	SChars  []rune
	Cursor  int
	escaped []int
}

// Is the character specified at `index` escaped?
func (p *parser) isEscaped(index int) bool {
	for _, i := range p.escaped {
		if i == index {
			return true
		}
	}

	return false
}

// Preproccess escaped characters and populate the
// `escaped` slice
func (p *parser) processEscapes() {
	cursor := 0
	for cursor < len(p.Input) {
		if p.Input[cursor] == '\\' {
			p.escaped = append(p.escaped, cursor+1)
			cursor++
		}
		cursor++
	}
}

// Simply returns whether the current char (or char + offset) is
// a special char (one specified in SChars)
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

// Consume a word until a special char and return it
func (p *parser) word() string {
	var word strings.Builder

	for p.Cursor < len(p.Input) {
		if p.Input[p.Cursor] == '\\' && !p.isEscaped(p.Cursor) {
			p.Cursor++
			continue
		}

		if p.isEscaped(p.Cursor) {
			switch c := (p.Input[p.Cursor]); c {
			case 'a':
				word.WriteByte('\a')
			case 'b':
				word.WriteByte('\b')
			case 'f':
				word.WriteByte('\f')
			case 'n':
				word.WriteByte('\n')
			case 'r':
				word.WriteByte('\r')
			case 't':
				word.WriteByte('\t')
			case 'v':
				word.WriteByte('\v')
			case ':', ';', ',', '\\':
				word.WriteByte(c)
			}
		} else if p.specialChar() {
			p.Cursor++
			break
		} else {
			word.WriteByte(p.Input[p.Cursor])
		}
		p.Cursor++
	}
	return word.String()
}

// Return what the next special char is
func (p *parser) peekChar() rune {
	for i, char := range p.Input[p.Cursor:] {
		if p.specialChar(i) && !p.isEscaped(i+p.Cursor) {
			return char
		}
	}
	return 0
}

// StrToMap takes an input string and creates a map out of it, for use
// in json requests. For example,
// "key1:key1-1:value1-1,key1-2:value1-2;key2:value2" returns
// {"key1": {"key1-1": "value1-1", "key1-2": "value1-2"}, "key2": "value2"}
func StrToMap(input string) (map[string]interface{}, int) {
	out := make(map[string]interface{})
	p := parser{Input: input,
		SChars: []rune{':', ';', ','}}

	p.processEscapes()

	for key := p.word(); key != ""; key = p.word() {
		char := p.peekChar()
		switch char {
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
		}
	}
	return out, p.Cursor
}
