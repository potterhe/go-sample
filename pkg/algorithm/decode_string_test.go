package algorithm

import (
	"fmt"
	"strconv"
	"testing"
)

func TestDecodeString(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "1[a]",
			output: "a",
		},
		{
			input:  "3[a]2[bc]",
			output: "aaabcbc",
		},
		{
			input:  "3[a2[c]]",
			output: "accaccacc",
		},
		{
			input:  "2[中国]",
			output: "中国中国",
		},
		{
			input:  "2[abc]3[cd]ef",
			output: "abcabccdcdcdef",
		},
		{
			input:  "abc3[cd]xyz",
			output: "abccdcdcdxyz",
		},
	}

	for _, test := range tests {
		fmt.Println(test)
		r := decodeString(test.input)
		if r != test.output {
			t.Errorf("except: %s, %s given", test.output, r)
		}
	}
}

func decodeString(s string) string {
	s1 := []rune(s)
	stack := runeStack{}
	var c rune
	//var r []rune
	st := 1 // 1:找[], 2:找系数
	var curr []rune
	//var coefficient int
	for i := len(s1) - 1; i >= 0; i-- {
		fmt.Printf("%q \n", s1[i])
		if s1[i] == '[' {
			// pop
			for {
				c = stack.pop()
				if c == ']' {
					break
				}
				curr = append(curr, c)
			}

			fmt.Printf("curr: %q\n", string(curr))
			// 找系数
			st = 2
			continue

		}
		if st == 2 {

			coefficient, _ := strconv.Atoi(string(s1[i]))
			for j := 0; j < coefficient; j++ {
				for m := len(curr) - 1; m >= 0; m-- {
					stack.push(curr[m])
				}
			}
			curr = curr[0:0]
			st = 1

		} else {
			stack.push(s1[i])
		}
	}

	var r []rune
	l := len(stack.Stack())
	for n := 0; n < l; n++ {
		r = append(r, stack.pop())
	}

	return string(r)
}

type runeStack struct {
	stack []rune
}

func (s *runeStack) push(r rune) {
	s.stack = append(s.stack, r)
}

func (s *runeStack) pop() rune {
	r := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return r
}

func (s *runeStack) Stack() []rune {
	return s.stack
}
