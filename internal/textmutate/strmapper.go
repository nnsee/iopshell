/*
   Copyright (c) 2018 Rasmus Moorats (neonsea)

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

package textmutate

import (
	"strings"
)

type parser struct {
	Input  string
	SChars []rune
	Cursor int
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
		if p.specialChar() {
			p.Cursor++
			return word.String()
		}
		word.WriteByte(p.Input[p.Cursor])
		p.Cursor++
	}
	return word.String()
}

// Return what the next special char is
func (p *parser) peekChar() rune {
	for i, char := range p.Input[p.Cursor:] {
		if p.specialChar(i) {
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
