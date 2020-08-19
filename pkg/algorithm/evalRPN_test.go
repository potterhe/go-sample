package algorithm

import (
	"fmt"
	"strconv"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		tokens []string
		result int
	}{
		{
			tokens: []string{"2", "1", "+", "3", "*"},
			result: 9,
		},
		{
			tokens: []string{"4", "13", "5", "/", "+"},
			result: 6,
		},
		{
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			result: 22,
		},
	}

	for _, test := range tests {
		//tokens := []string{"2", "1", "+", "3", "*"}

		r1 := evalRPN(test.tokens)
		if test.result != r1 {
			t.Errorf("except: %d, %d given", test.result, r1)
		}
	}
	fmt.Println("---")
	for _, test := range tests {
		//tokens := []string{"2", "1", "+", "3", "*"}

		r1 := evalRPN2(test.tokens)
		if test.result != r1 {
			t.Errorf("except: %d, %d given", test.result, r1)
		}
	}
}

func evalRPN(tokens []string) int {
	number := []int{}
	for _, val := range tokens {
		l := len(number)
		switch val {
		case "+":
			number = append(number[:l-2], number[l-2]+number[l-1])
		case "-":
			number = append(number[:l-2], number[l-2]-number[l-1])
		case "*":
			number = append(number[:l-2], number[l-2]*number[l-1])
		case "/":
			number = append(number[:l-2], number[l-2]/number[l-1])
		default:
			num, _ := strconv.Atoi(val)
			number = append(number, num)
		}
		fmt.Println(number)
	}
	return number[0]
}

func evalRPN2(tokens []string) int {
	number := intStack{}
	var l, r int
	for _, val := range tokens {
		switch val {
		case "+":
			l = number.pop()
			r = number.pop()
			number.push(l + r)
		case "-":
			r = number.pop()
			l = number.pop()
			number.push(l - r)
		case "*":
			l = number.pop()
			r = number.pop()
			number.push(l * r)
		case "/":
			r = number.pop()
			l = number.pop()
			number.push(l / r)
		default:
			num, _ := strconv.Atoi(val)
			number.push(num)
		}
		fmt.Println(number)
	}
	return number.pop()
}

type intStack struct {
	s []int
}

func (s *intStack) push(n int) {
	s.s = append(s.s, n)
}

func (s *intStack) pop() int {
	v := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return v
}
