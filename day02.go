package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"sort"
	"strconv"
	"strings"
)

func min(v ...int) int {
	vmin := v[0]
	for _, e := range v {
		if e < vmin {
			vmin = e
		}
	}
	return vmin
}

func srt(v ...int) []int {
	sort.Ints(v)
	return v
}

func Day02(input io.Reader) {
	sc := bufio.NewScanner(input)
	ln := 0
	area, rlen := 0, 0
	for sc.Scan() {
		ln++
		line := sc.Text()
		dims := strings.Split(line, "x")
		if len(dims) != 3 {
			log.Fatalf("Cannot parse line %d: %s", ln, line)
		}
		l, err := strconv.Atoi(dims[0])
		if err != nil {
			log.Fatalf("%d: l: %s", err.Error())
		}
		w, err := strconv.Atoi(dims[1])
		if err != nil {
			log.Fatalf("%d: w: %s", err.Error())
		}
		h, err := strconv.Atoi(dims[2])
		if err != nil {
			log.Fatalf("%d: h: %s", err.Error())
		}
		a1, a2, a3 := l*w, w*h, h*l
		extra := min(a1, a2, a3)
		area += 2*(a1+a2+a3) + extra

		sdim := srt(l, w, h)
		rlen += 2*sdim[0] + 2*sdim[1]
		rlen += w * l * h
	}
	if sc.Err() != nil {
		log.Fatal(sc.Err().Error())
	}
	fmt.Println(area, rlen)
}
