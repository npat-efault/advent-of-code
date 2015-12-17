package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

func Day01(input io.Reader) {
	bin := bufio.NewReader(input)
	floor, bpos := 0, -1
	for i := 0; ; i++ {
		c, _, err := bin.ReadRune()
		if err != nil {
			if err != io.EOF {
				log.Fatal(err.Error())
			}
			break
		}
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}
		if bpos == -1 && floor == -1 {
			bpos = i + 1
		}
	}
	fmt.Println(floor, bpos)
}
