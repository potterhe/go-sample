/*
1000瓶水里面有一瓶毒水，通过用兔子去喝水的方式检验，只能让兔子喝一次水，那么最少用多少只兔子检验出来是哪一瓶？
*/
package algorithm

import (
	"fmt"
	"testing"
)

func TestRabbitToxicant(t *testing.T) {

	tests := []struct {
		n        int
		r        int
		die      []int
		toxicant int
	}{
		//{1, 1},
		//{2, 1},
		//{3, 2},
		//{4, 3},
		//{8, 4},
		{10, 4, []int{1}, 1},
		{10, 4, []int{1, 2}, 3},
		{10, 4, []int{1, 2, 3}, 7},
		//{1000, 10},
	}

	for _, test := range tests {
		mapRabbit := rabbitToxicant(test.n)
		numRabbit := len(mapRabbit)
		if numRabbit != test.r {
			t.Errorf("expect %d, %d given", test.r, numRabbit)
		}
		sfmt := fmt.Sprintf("%%%db %%s %%d\n", numRabbit)
		fmt.Println(sfmt)
		for k, v := range mapRabbit {
			fmt.Printf(sfmt, k, "rabbit", k)
			for _, v1 := range v {
				fmt.Printf(sfmt, v1, "bottle", v1)
			}
		}

		toxicant := findToxicant(mapRabbit, test.die)
		if toxicant != test.toxicant {
			t.Errorf("find toxicant: expect %d, %d given\n", test.toxicant, toxicant)
		}
	}

	a := 0x1
	b := 0x1
	c := 0x4
	d := 0x8

	e := a | b | c | d

	fmt.Printf("%b, %b\n", e, e&^c)

	v1 := 1
	v2 := 3

	fmt.Printf("%b, %b, %b\n", v1, v2, v1^v2)
}

func rabbitToxicant(n int) map[uint64][]int {
	if n <= 0 {
		return nil
	}

	buf := fmt.Sprintf("%b", n)
	numRabbit := len([]byte(buf))

	rabbits := make([]uint64, numRabbit)
	for i := 0; i < numRabbit; i++ {
		rabbits[i] = 1 << i
		fmt.Printf("%b\n", rabbits[i])
	}

	pending := make([]uint64, n)
	mapRabbit := make(map[uint64][]int, numRabbit)
	for i := 1; i <= n; i++ {
		pending[i-1] = uint64(i)

		for _, rabbit := range rabbits {
			if uint64(i)&rabbit > 0 {
				mapRabbit[rabbit] = append(mapRabbit[rabbit], i)
			}
		}
	}

	return mapRabbit
}

func findToxicant(mapRabbit map[uint64][]int, die []int) int {
	var mask uint64 = 0
	var rabbit uint64
	for _, v := range die {
		rabbit = 1 << (v - 1)
		mask = mask | rabbit
	}
	fmt.Printf("mask: %b \n", mask)

	for k, v := range mapRabbit {
		fmt.Printf("rabbit: %b  k&mask: %b\n", k, k&mask)
		if k&mask > 0 {
			for _, p := range v {
				fmt.Printf("p: %b  p^mask: %b\n", p, uint64(p)^mask)
				if uint64(p)^mask == 0 {
					return p
				}
			}
		}
	}

	return 0
}
